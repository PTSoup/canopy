package p2p

import (
	"encoding/binary"
	"io"
	"net"
	"runtime/debug"
	"sync"
	"time"

	"github.com/canopy-network/canopy/lib/crypto"

	"github.com/alecthomas/units"
	"github.com/canopy-network/canopy/lib"
	limiter "github.com/mxk/go-flowrate/flowrate"
	"google.golang.org/protobuf/proto"
)

const (
	maxDataChunkSize       = 1024 - packetHeaderSize // maximum size of the chunk of bytes in a packet
	maxPacketSize          = 1024                    // maximum size of the full packet
	packetHeaderSize       = 47                      // the overhead of the protobuf packet header
	queueSendTimeout       = 10 * time.Second        // how long a message waits to be queued before throwing an error
	dataFlowRatePerS       = 500 * units.KB          // the maximum number of bytes that may be sent or received per second per MultiConn
	maxMessageSize         = 10 * units.Megabyte     // the maximum total size of a message once all the packets are added up
	maxChanSize            = 1                       // maximum number of items in a channel before blocking
	maxInboxQueueSize      = 15                      // maxinum number of items in inbox queue before blocking
	maxStreamSendQueueSize = 15                      // maximum number of items in a stream send queue before blocking

	// "Peer Reputation Points" are actively maintained for each peer the node is connected to
	// These points allow a node to track peer behavior over its lifetime, allowing it to disconnect from faulty peers
	PollMaxHeightTimeoutS   = 1   // wait time for polling the maximum height of the peers
	SyncTimeoutS            = 5   // wait time to receive an individual block (certificate) from a peer during syncing
	MaxBlockReqPerWindow    = 20  // maximum block (certificate) requests per window per requester
	BlockReqWindowS         = 2   // the 'window of time' before resetting limits for block (certificate) requests
	GoodPeerBookRespRep     = 3   // reputation points for a good peer book response
	GoodBlockRep            = 3   // rep boost for sending us a valid block (certificate)
	GoodTxRep               = 3   // rep boost for sending us a valid transaction (certificate)
	BadPacketSlash          = -1  // bad packet is received
	NoPongSlash             = -1  // no pong received
	TimeoutRep              = -1  // rep slash for not responding in time
	UnexpectedBlockRep      = -1  // rep slash for sending us a block we weren't expecting
	PeerBookReqTimeoutRep   = -1  // slash for a non-response for a peer book request
	UnexpectedMsgRep        = -1  // slash for an unexpected message
	InvalidMsgRep           = -3  // slash for an invalid message
	ExceedMaxPBReqRep       = -3  // slash for exceeding the max peer book requests
	ExceedMaxPBLenRep       = -3  // slash for exceeding the size of the peer book message
	UnknownMessageSlash     = -3  // unknown message type is received
	BadStreamSlash          = -3  // unknown stream id is received
	InvalidTxRep            = -3  // rep slash for sending us an invalid transaction
	NotValRep               = -3  // rep slash for sending us a validator only message but not being a validator
	InvalidBlockRep         = -3  // rep slash for sending an invalid block (certificate) message
	InvalidJustifyRep       = -3  // rep slash for sending an invalid certificate justification
	BlockReqExceededRep     = -3  // rep slash for over-requesting blocks (certificates)
	MaxMessageExceededSlash = -10 // slash for sending a 'Message (sum of Packets)' above the allowed maximum size
)

var (
	ReadWriteTimeout = 40 * time.Second // this is just the default; it gets set by config upon initialization
)

// MultiConn: A rate-limited, multiplexed connection that utilizes a series streams with varying priority for sending and receiving
type MultiConn struct {
	conn          net.Conn                    // underlying connection
	Address       *lib.PeerAddress            // authenticated peer information
	streams       map[lib.Topic]*Stream       // multiple independent bi-directional communication channels
	quitSending   chan struct{}               // signal to quit
	quitReceiving chan struct{}               // signal to quit
	onError       func(error, []byte, string) // callback to call if peer errors
	error         sync.Once                   // thread safety to ensure MultiConn.onError is only called once
	p2p           *P2P                        // a pointer reference to the P2P module
	log           lib.LoggerI                 // logging
}

// NewConnection() creates and starts a new instance of a MultiConn
func (p *P2P) NewConnection(conn net.Conn) (*MultiConn, lib.ErrorI) {
	// establish an encrypted connection using the handshake
	eConn, err := NewHandshake(conn, p.meta, p.privateKey)
	if err != nil {
		return nil, err
	}
	c := &MultiConn{
		conn:          eConn,
		Address:       eConn.Address,
		streams:       p.NewStreams(),
		quitSending:   make(chan struct{}, maxChanSize),
		quitReceiving: make(chan struct{}, maxChanSize),
		onError:       p.OnPeerError,
		error:         sync.Once{},
		p2p:           p,
		log:           p.log,
	}
	_ = c.conn.SetReadDeadline(time.Time{})
	_ = c.conn.SetWriteDeadline(time.Time{})
	// start the connection service
	c.Start()
	return c, err
}

// Start() begins send and receive services for a MultiConn
func (c *MultiConn) Start() {
	go c.startSendService()
	go c.startReceiveService()
}

// Stop() sends exit signals for send and receive loops and closes the connection
func (c *MultiConn) Stop() {
	c.p2p.log.Warnf("Stopping peer %s", lib.BytesToString(c.Address.PublicKey))
	c.quitReceiving <- struct{}{}
	c.quitSending <- struct{}{}
	close(c.quitSending)
	close(c.quitReceiving)
	_ = c.conn.Close()
}

// Send() queues the sending of a message to a specific Stream
func (c *MultiConn) Send(topic lib.Topic, msg *Envelope) (ok bool) {
	stream, ok := c.streams[topic]
	if !ok {
		c.log.Errorf("Stream %s does not exist", topic)
		return
	}
	bz, err := lib.Marshal(msg)
	if err != nil {
		c.log.Errorf("Marshaling failed of payload: %s", msg.String())
		return false
	}
	chunks := split(bz, maxDataChunkSize)
	var packets []*Packet
	for i, chunk := range chunks {
		packets = append(packets, &Packet{
			StreamId: topic,
			Eof:      i == len(chunks)-1,
			Bytes:    chunk,
		})
	}
	ok = stream.queueSends(packets)
	return
}

// startSendService() starts the main send service
// - converges and writes the send queue from all streams into the underlying tcp connection.
// - manages the keep alive protocol by sending pings and monitoring the receipt of the corresponding pong
func (c *MultiConn) startSendService() {
	defer func() {
		if r := recover(); r != nil {
			c.log.Errorf("panic recovered, err: %s, stack: %s", r, string(debug.Stack()))
		}
	}()
	m := limiter.New(0, 0)
	var packet *Packet
	defer func() { m.Done() }()
	for {
		// select statement ensures the sequential coordination of the concurrent processes
		select {
		case packet = <-c.streams[lib.Topic_CONSENSUS].sendQueue:
		case packet = <-c.streams[lib.Topic_BLOCK].sendQueue:
		case packet = <-c.streams[lib.Topic_BLOCK_REQUEST].sendQueue:
		case packet = <-c.streams[lib.Topic_TX].sendQueue:
		case packet = <-c.streams[lib.Topic_PEERS_RESPONSE].sendQueue:
		case packet = <-c.streams[lib.Topic_PEERS_REQUEST].sendQueue:
		case <-c.quitSending: // fires when Stop() is called
			return
		}
		c.sendPacket(packet, m)
	}
}

// startReceiveService() starts the main receive service
// - reads from the underlying tcp connection and 'routes' the messages to the appropriate streams
// - manages keep alive protocol by notifying the 'send service' of pings and pongs
func (c *MultiConn) startReceiveService() {
	defer func() {
		if r := recover(); r != nil {
			c.log.Errorf("panic recovered, err: %s, stack: %s", r, string(debug.Stack()))
		}
	}()
	m := limiter.New(0, 0)
	defer func() { m.Done() }()
	for {
		select {
		default: // fires unless quit was signaled
			// waits until bytes are received from the conn
			msg, err := c.waitForAndHandleWireBytes(m)
			if err != nil {
				c.Error(err)
				return
			}
			// handle different message types
			switch x := msg.(type) {
			case *Packet: // receive packet is a partial or full 'Message' with a Stream Topic designation and an EOF signal
				// load the proper stream
				stream, found := c.streams[x.StreamId]
				if !found {
					c.Error(ErrBadStream(), BadStreamSlash)
					return
				}
				// get the peer info from the peer set
				info, e := c.p2p.GetPeerInfo(c.Address.PublicKey)
				if e != nil {
					c.Error(e)
					return
				}
				// handle the packet within the stream
				if slash, er := stream.handlePacket(info, x); er != nil {
					c.log.Warnf(er.Error())
					c.Error(er, slash)
					return
				}
			default: // unknown type results in slash and exiting the service
				c.Error(ErrUnknownP2PMsg(x), UnknownMessageSlash)
				return
			}
		case <-c.quitReceiving: // fires when quit is signaled
			return
		}
	}
}

// Error() when an error occurs on the MultiConn execute a callback. Optionally pass a reputation delta to slash the peer
func (c *MultiConn) Error(err error, reputationDelta ...int32) {
	if len(reputationDelta) == 1 {
		c.p2p.ChangeReputation(c.Address.PublicKey, reputationDelta[0])
	}
	// call onError() for the peer
	c.error.Do(func() { c.onError(err, c.Address.PublicKey, c.conn.RemoteAddr().String()) })
}

// waitForAndHandleWireBytes() a rate limited handler of inbound bytes from the wire.
// Blocks until bytes are received converts bytes into a proto.Message using an Envelope
func (c *MultiConn) waitForAndHandleWireBytes(m *limiter.Monitor) (proto.Message, lib.ErrorI) {
	// initialize the wrapper object
	msg := new(Envelope)
	// restrict the instantaneous data flow to rate bytes per second
	// Limit() request maxPacketSize bytes from the limiter and the limiter
	// will block the execution until at or below the desired rate of flow
	//m.Limit(maxPacketSize, int64(dataFlowRatePerS), true)
	// read the proto message from the wire
	if err := receiveProtoMsg(c.conn, msg); err != nil {
		return nil, err
	}
	// unmarshal the payload from proto.any
	return lib.FromAny(msg.Payload)
}

// sendPacket() a rate limited writer of outbound bytes to the wire
// wraps a proto.Message into a universal Envelope, then converts to bytes and
// sends them across the wire without violating the data flow rate limits
// message may be a Packet, a Ping or a Pong
func (c *MultiConn) sendPacket(packet *Packet, m *limiter.Monitor) (err lib.ErrorI) {
	c.log.Debugf("Send Packet to %s (ID:%s, L:%d, E:%t), hash: %s",
		lib.BytesToTruncatedString(c.Address.PublicKey),
		lib.Topic_name[int32(packet.StreamId)],
		len(packet.Bytes),
		packet.Eof,
		crypto.ShortHashString(packet.Bytes),
	)
	// convert the proto.Message into a proto.Any
	a, err := lib.NewAny(packet)
	if err != nil {
		return err
	}
	// restrict the instantaneous data flow to rate bytes per second
	// Limit() request maxPacketSize bytes from the limiter and the limiter
	// will block the execution until at or below the desired rate of flow
	//m.Limit(maxPacketSize, int64(dataFlowRatePerS), true)
	startTime := time.Now()
	// send the proto message wrapped in an Envelope over the wire
	if err = sendProtoMsg(c.conn, &Envelope{Payload: a}); err != nil {
		c.Error(err)
	}
	// debug log to remove
	lib.TimeTrack("MultiConn.sendPacket's c.conn.Write", startTime)
	// update the rate limiter with how many bytes were written
	//m.Update(n)
	return
}

// Stream: an independent, bidirectional communication channel that is scoped to a single topic.
// In a multiplexed connection there is typically more than one stream per connection
type Stream struct {
	topic        lib.Topic                    // the subject and priority of the stream
	sendQueue    chan *Packet                 // a queue of unsent messages
	msgAssembler []byte                       // collects and adds incoming packets until the entire message is received (EOF signal)
	inbox        chan *lib.MessageAndMetadata // the channel where fully received messages are held for other parts of the app to read
	mu           sync.Mutex                   // mutex to prevent race conditions when sending packets (all packets of the same message should be one right after the other)
	logger       lib.LoggerI
}

// queueSends() schedules the packets to be sent ensuring coordination with the mutex
func (s *Stream) queueSends(packets []*Packet) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, packet := range packets {
		ok := s.queueSend(packet)
		if !ok {
			s.logger.Errorf("Packet(ID:%s, L:%d, E:%t) packet failed in queue", lib.Topic_name[int32(packet.StreamId)], len(packet.Bytes), packet.Eof)
			return false
		}
		s.logger.Debugf("Packet(ID:%s, L:%d, E:%t) packet queued", lib.Topic_name[int32(packet.StreamId)], len(packet.Bytes), packet.Eof)
	}
	return true
}

// queueSend() schedules the packet to be sent
func (s *Stream) queueSend(p *Packet) bool {
	s.logger.Debugf("Queuing packet (ID:%s, Q:%d, C:%d)",
		lib.Topic_name[int32(p.StreamId)], len(s.sendQueue), cap(s.sendQueue))
	select {
	case s.sendQueue <- p: // enqueue to the back of the line
		return true
	case <-time.After(queueSendTimeout): // may timeout if queue remains full
		return false
	}
}

// handlePacket() merge the new packet with the previously received ones until the entire message is complete (EOF signal)
func (s *Stream) handlePacket(peerInfo *lib.PeerInfo, packet *Packet) (int32, lib.ErrorI) {
	msgAssemblerLen, packetLen := len(s.msgAssembler), len(packet.Bytes)
	s.logger.Debugf("Received Packet from %s (ID:%s, L:%d, E:%t), hash: %s",
		lib.BytesToTruncatedString(peerInfo.Address.PublicKey),
		lib.Topic_name[int32(packet.StreamId)],
		len(packet.Bytes),
		packet.Eof,
		crypto.ShortHashString(packet.Bytes),
	)
	// if the addition of this new packet pushes the total message size above max
	if int(maxMessageSize) < msgAssemblerLen+packetLen {
		s.msgAssembler = s.msgAssembler[:0]
		return MaxMessageExceededSlash, ErrMaxMessageSize()
	}
	// combine this packet with the previously received ones
	s.msgAssembler = append(s.msgAssembler, packet.Bytes...)
	// if the packet is signalling message end
	if packet.Eof {
		// unmarshall all the bytes into the universal wrapper
		var msg Envelope
		if err := lib.Unmarshal(s.msgAssembler, &msg); err != nil {
			return BadPacketSlash, err
		}
		// read the payload into a proto.Message
		payload, err := lib.FromAny(msg.Payload)
		if err != nil {
			return BadPacketSlash, err
		}
		// wrap with metadata
		m := (&lib.MessageAndMetadata{
			Message: payload,
			Sender:  peerInfo,
		}).WithHash()
		// add to inbox for other parts of the app to read
		s.logger.Debugf("inbox %s queue: %d", lib.Topic_name[int32(packet.StreamId)], len(s.inbox))
		s.inbox <- m
		// reset receiving buffer
		s.msgAssembler = s.msgAssembler[:0]
	}
	return 0, nil
}

// HELPERS BELOW

// sendProtoMsg() encodes and sends a length-prefixed proto message to a net.Conn
func sendProtoMsg(conn net.Conn, ptr proto.Message) lib.ErrorI {
	// marshal into proto bytes
	bz, err := lib.Marshal(ptr)
	if err != nil {
		return err
	}
	// send the bytes prefixed by length
	return sendLengthPrefixed(conn, bz)
}

// receiveProtoMsg() receives and decodes a length-prefixed proto message from a net.Conn
func receiveProtoMsg(conn net.Conn, ptr proto.Message) lib.ErrorI {
	// read the message from the wire
	msg, err := receiveLengthPrefixed(conn)
	if err != nil {
		return err
	}
	// unmarshal into proto
	if err = lib.Unmarshal(msg, ptr); err != nil {
		return err
	}
	return nil
}

// sendLengthPrefixed() sends a message that is prefix by length through a tcp connection
func sendLengthPrefixed(conn net.Conn, bz []byte) lib.ErrorI {
	// create the length prefix (2 bytes, big endian)
	lengthPrefix := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthPrefix, uint16(len(bz)))
	//// set the write deadline to 20 second
	if e := conn.SetWriteDeadline(time.Now().Add(ReadWriteTimeout)); e != nil {
		return ErrFailedWrite(e)
	}
	// write the message (length prefixed)
	if _, er := conn.Write(append(lengthPrefix, bz...)); er != nil {
		return ErrFailedWrite(er)
	}
	// disable deadline
	_ = conn.SetWriteDeadline(time.Time{})
	return nil
}

// receiveLengthPrefixed() reads a length prefixed message from a tcp connection
func receiveLengthPrefixed(conn net.Conn) ([]byte, lib.ErrorI) {
	// set the read conn deadline
	if err := conn.SetReadDeadline(time.Now().Add(ReadWriteTimeout)); err != nil {
		return nil, ErrFailedRead(err)
	}
	// read the 2-byte length prefix
	lengthBuffer := make([]byte, 2)
	if _, err := io.ReadFull(conn, lengthBuffer); err != nil {
		return nil, ErrFailedRead(err)
	}
	// determine the length of the message
	messageLength := binary.BigEndian.Uint16(lengthBuffer)
	// ensure the message size isn't larger than the allowed max packet size
	if messageLength > maxPacketSize {
		return nil, ErrMaxMessageSize()
	}
	// read the actual message bytes
	msg := make([]byte, messageLength)
	if _, err := io.ReadFull(conn, msg); err != nil {
		return nil, ErrFailedRead(err)
	}
	// disable deadline
	_ = conn.SetReadDeadline(time.Time{})
	// exit with no error
	return msg, nil
}

// split returns bytes split to size up to the lim param
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
