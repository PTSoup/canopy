package p2p

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/ginchuco/ginchu/lib"
	"github.com/ginchuco/ginchu/lib/crypto"
	"github.com/phuslu/iploc"
	"golang.org/x/net/netutil"
	"google.golang.org/protobuf/proto"
	"math"
	"net"
	"runtime/debug"
	"slices"
	"sync"
	"time"
)

/*
	P2P
	- TCP/IP transport [x]
	- Multiplexing [x]
	- Encrypted connection [x]
	- UnPn & nat-pimp auto config [-]
	- DOS mitigation [x]
	- Peer configs: unconditional, num in/out, timeouts [x]
	- Peer list: discover[x], churn[x], share[x]
	- Message dissemination: gossip [x]
*/

const transport, dialTimeout = "tcp", time.Second

type P2P struct {
	privateKey             crypto.PrivateKeyI
	listener               net.Listener
	channels               lib.Channels
	meta                   *lib.PeerMeta
	PeerSet                          // active set
	book                   *PeerBook // not active set
	mustConnectReceiver    chan []*lib.PeerAddress
	maxMembersPerCommittee int
	bannedIPs              []net.IPAddr // banned IPs (non-string)
	config                 lib.Config
	log                    lib.LoggerI
}

// New() creates an initialized pointer instance of a P2P object
func New(p crypto.PrivateKeyI, maxMembersPerCommittee uint64, c lib.Config, l lib.LoggerI) *P2P {
	// Initialize the peer book
	peerBook := NewPeerBook(p.PublicKey().Bytes(), c, l)
	// Make inbound multiplexed channels
	channels := make(lib.Channels)
	for i := lib.Topic(0); i < lib.Topic_INVALID; i++ {
		channels[i] = make(chan *lib.MessageAndMetadata, maxChanSize)
	}
	// Load banned IPs
	var bannedIPs []net.IPAddr
	for _, ip := range c.BannedIPs {
		i, err := net.ResolveIPAddr("", ip)
		if err != nil {
			l.Fatalf(err.Error())
		}
		bannedIPs = append(bannedIPs, *i)
	}
	// Load supported chains
	var chains []uint64
	for _, plug := range c.Plugins {
		chains = append(chains, plug.ID)
	}
	meta := &lib.PeerMeta{Chains: chains}
	return &P2P{
		privateKey:             p,
		channels:               channels,
		config:                 c,
		meta:                   meta.Sign(p),
		PeerSet:                NewPeerSet(c),
		book:                   peerBook,
		mustConnectReceiver:    make(chan []*lib.PeerAddress, maxChanSize),
		maxMembersPerCommittee: int(maxMembersPerCommittee),
		bannedIPs:              bannedIPs,
		log:                    l,
	}
}

// Start() begins the P2P service
func (p *P2P) Start() {
	p.log.Info("Starting P2P 🤝 ")
	// Listens for 'must connect peer ids' from the main internal controller
	go p.ListenForMustConnects(p.mustConnectReceiver)
	// Starts the peer address book exchange service
	go p.StartPeerBookService()
	// Listens for external inbound peers
	go p.ListenForInboundPeers(&lib.PeerAddress{NetAddress: p.config.ListenAddress})
	// Dials external outbound peers
	go p.DialForOutboundPeers()
}

// Stop() stops the P2P service
func (p *P2P) Stop() {
	// it's possible the listener has not yet been initialized before stopping
	if p.listener != nil {
		if err := p.listener.Close(); err != nil {
			p.log.Error(err.Error())
		}
	}
	// gracefully closes all the existing connections
	p.PeerSet.Stop()
}

// ListenForInboundPeers() starts a rate-limited tcp listener service to accept inbound peers
func (p *P2P) ListenForInboundPeers(listenAddress *lib.PeerAddress) {
	ln, er := net.Listen(transport, listenAddress.NetAddress)
	if er != nil {
		p.log.Fatal(ErrFailedListen(er).Error())
	}
	p.log.Debugf("Starting net.Listener on tcp://%s", listenAddress.NetAddress)
	p.listener = netutil.LimitListener(ln, p.MaxPossibleInbound())
	// continuous service until program exit
	for {
		// wait for and then accept inbound tcp connection
		c, err := p.listener.Accept()
		if err != nil {
			return
		}
		// create a thread to prevent front-of-the-line blocking
		go func(c net.Conn) {
			// ephemeral connections are basic, inbound tcp connections
			defer p.catchPanic()
			p.log.Debugf("Received ephemeral connection %s", c.RemoteAddr().String())
			// begin to create a peer address using the inbound tcp conn while filtering any bad ips
			netAddress, e := p.filterBadIPs(c)
			if e != nil {
				p.log.Debugf("Closing ephemeral connection %s", c.RemoteAddr().String())
				_ = c.Close()
				return
			}
			// tries to create a full peer from the ephemeral connection and just the net address
			if err = p.AddPeer(c, &lib.PeerInfo{Address: &lib.PeerAddress{NetAddress: netAddress}}, false); err != nil {
				_ = c.Close()
				return
			}
		}(c)
	}
}

// DialForOutboundPeers() uses the config and peer book to try to max out the outbound peer connections
func (p *P2P) DialForOutboundPeers() {
	dialing := &Dialing{
		m: make(map[uint64]int),
		l: sync.RWMutex{},
	}
	// Try to connect to the DialPeers in the config
	for _, peerString := range p.config.DialPeers {
		peer := new(lib.PeerAddress)
		if err := peer.FromString(peerString); err != nil {
			p.log.Errorf("invalid dial peer %s: %s", peerString, err.Error())
			continue
		}
		go func() {
			// preload dial with backoff
			dialing.AddDialing(math.MaxUint16) // unknown chains
			// once dial with backoff finishes
			defer dialing.RemoveDialing(math.MaxUint16) // unknown chains
			// dial the peer with exponential backoff
			p.DialWithBackoff(peer)
		}()
	}
	// Continuous service until program exit, dial timeout loop frequency for resource break
	for {
		time.Sleep(dialTimeout)
		// for each supported plugin, try to max out peer config by dialing
		for _, chain := range p.meta.Chains {
			func() {
				// exit if maxed out config or none left to dial
				if (p.PeerSet.outbound[chain]+dialing.NumDialing(chain) >= p.config.MaxOutbound) || p.book.GetBookSizeForChain(chain) == 0 {
					return
				}
				// get random peer for chain
				rand := p.book.GetRandom(chain)
				if rand == nil || p.IsSelf(rand.Address) || p.Has(rand.Address.PublicKey) {
					return
				}
				// sequential operation means we'll never be dialing more than 1 peer at a time
				// the peer should be added before the next execution of the loop
				dialing.AddDialing(chain)
				defer dialing.RemoveDialing(chain)
				if err := p.Dial(rand.Address, false); err != nil {
					p.log.Warn(err.Error())
					return
				}
			}()
		}
	}
}

// Dial() tries to establish an outbound connection with a peer candidate
func (p *P2P) Dial(address *lib.PeerAddress, disconnect bool) lib.ErrorI {
	if p.IsSelf(address) || p.PeerSet.Has(address.PublicKey) {
		return nil
	}
	p.log.Debugf("Dialing %s@%s", lib.BytesToString(address.PublicKey), address.NetAddress)
	// try to establish the basic tcp connection
	conn, er := net.DialTimeout(transport, address.NetAddress, dialTimeout)
	if er != nil {
		return ErrFailedDial(er)
	}
	// try to use the basic tcp connection to establish a peer
	return p.AddPeer(conn, &lib.PeerInfo{Address: address, IsOutbound: true}, disconnect)
}

// AddPeer() takes an ephemeral tcp connection and an incomplete peerInfo and attempts to
// create a E2E encrypted channel with a fully authenticated peer and save it to
// the peer set and the peer book
func (p *P2P) AddPeer(conn net.Conn, info *lib.PeerInfo, disconnect bool) (err lib.ErrorI) {
	// create the e2e encrypted connection while establishing a full peer info object
	connection, err := p.NewConnection(conn)
	if err != nil {
		return err
	}
	p.log.Debugf("Try Add peer: %s@%s", lib.BytesToString(connection.Address.PublicKey), info.Address.NetAddress)
	// if peer is outbound, ensure the public key matches who we expected to dial
	if info.IsOutbound {
		if !bytes.Equal(connection.Address.PublicKey, info.Address.PublicKey) {
			return ErrMismatchPeerPublicKey(info.Address.PublicKey, connection.Address.PublicKey)
		}
	}
	// overwrite the incomplete peer info with the complete and authenticated info
	info.Address = connection.Address
	// disconnect immediately if prompted by params
	if disconnect {
		connection.Stop()
		return nil
	}
	// check if is must connect
	p.RLock()
	for _, item := range p.mustConnect {
		if bytes.Equal(item.PublicKey, info.Address.PublicKey) {
			info.IsMustConnect = true
			break
		}
	}
	p.RUnlock()
	// check if is trusted
	for _, item := range p.config.TrustedPeerIDs {
		if item == lib.BytesToString(info.Address.PublicKey) {
			info.IsTrusted = true
			break
		}
	}
	// check if is banned
	for _, item := range p.config.BannedPeerIDs {
		pubKeyString := lib.BytesToString(info.Address.PublicKey)
		if pubKeyString == item {
			return ErrBannedID(pubKeyString)
		}
	}
	// add peer to peer set and peer book
	p.log.Infof("Adding peer: %s@%s", lib.BytesToString(info.Address.PublicKey), info.Address.NetAddress)
	p.book.Add(&BookPeer{Address: info.Address})
	return p.PeerSet.Add(&Peer{
		conn:     connection,
		PeerInfo: info,
		stop:     sync.Once{},
	})
}

// DialWithBackoff() dials the peer with exponential backoff retry
func (p *P2P) DialWithBackoff(peerInfo *lib.PeerAddress) {
	dialAndLog := func() (err error) {
		if err = p.Dial(peerInfo, false); err != nil {
			p.log.Errorf("Dial %s@%s failed: %s", peerInfo.PublicKey, peerInfo.NetAddress, err.Error())
		}
		return
	}
	_ = backoff.Retry(dialAndLog, backoff.NewExponentialBackOff())
}

// DialAndDisconnect() dials the peer but disconnects once a fully authenticated connection is established
func (p *P2P) DialAndDisconnect(a *lib.PeerAddress) lib.ErrorI { return p.Dial(a, true) }

// OnPeerError() callback to P2P when a peer errors
func (p *P2P) OnPeerError(err error, publicKey []byte, remoteAddr string) {
	p.log.Warn(PeerError(publicKey, remoteAddr, err))
	// ignore error: peer may have disconnected before added
	peer, _ := p.PeerSet.Remove(publicKey)
	if peer != nil {
		peer.stop.Do(peer.conn.Stop)
	}
	p.book.Remove(publicKey)
}

// NewStreams() creates map of streams for the multiplexing architecture
func (p *P2P) NewStreams() (streams map[lib.Topic]*Stream) {
	streams = make(map[lib.Topic]*Stream)
	for i := lib.Topic(0); i < lib.Topic_INVALID; i++ {
		streams[i] = &Stream{
			topic:     i,
			sendQueue: make(chan []byte, maxQueueSize),
			inbox:     p.Inbox(i),
		}
	}
	return
}

// IsSelf() returns if the peer address public key equals the self public key
func (p *P2P) IsSelf(a *lib.PeerAddress) bool {
	return bytes.Equal(p.privateKey.PublicKey().Bytes(), a.PublicKey)
}

// SelfSend() executes an internal pipe send to self
func (p *P2P) SelfSend(fromPublicKey []byte, topic lib.Topic, payload proto.Message) lib.ErrorI {
	p.log.Debugf("Self sending %s message", topic)
	p.Inbox(topic) <- (&lib.MessageAndMetadata{
		Message: payload,
		Sender:  &lib.PeerInfo{Address: &lib.PeerAddress{PublicKey: fromPublicKey}},
	}).WithHash()
	return nil
}

// MaxPossiblePeers() sums the MaxIn, MaxOut, MaxCommitteeConnects and trusted peer IDs
func (p *P2P) MaxPossiblePeers() int {
	return (p.config.MaxInbound+p.config.MaxOutbound+p.maxMembersPerCommittee)*len(p.config.Plugins) + len(p.config.TrustedPeerIDs)
}

// MaxPossibleInbound() sums the MaxIn, MaxCommitteeConnects and trusted peer IDs
func (p *P2P) MaxPossibleInbound() int {
	return (p.config.MaxInbound+p.maxMembersPerCommittee)*len(p.config.Plugins) + len(p.config.TrustedPeerIDs)
}

// MaxPossibleOutbound() sums the MaxIn, MaxCommitteeConnects and trusted peer IDs
func (p *P2P) MaxPossibleOutbound() int {
	return (p.config.MaxOutbound+p.maxMembersPerCommittee)*len(p.config.Plugins) + len(p.config.TrustedPeerIDs)
}

// Inbox() is a getter for the multiplexed stream with a specific topic
func (p *P2P) Inbox(topic lib.Topic) chan *lib.MessageAndMetadata { return p.channels[topic] }

// MustConnectReceiver() is a getter for the receiver that the controller uses to update 'must connect peers'
func (p *P2P) MustConnectReceiver() chan []*lib.PeerAddress { return p.mustConnectReceiver }

// ListenForMustConnects() is an internal listener that receives 'must connect peers' updates from the controller
func (p *P2P) ListenForMustConnects(receiver chan []*lib.PeerAddress) {
	for mustConnect := range receiver {
		// UpdateMustConnects() removes connections that are already established
		for _, val := range p.UpdateMustConnects(mustConnect) {
			go p.DialWithBackoff(val)
		}
	}
}

// ID() returns the self peer address
func (p *P2P) ID() *lib.PeerAddress {
	return &lib.PeerAddress{
		PublicKey:  p.privateKey.PublicKey().Bytes(),
		NetAddress: p.config.ExternalAddress,
		PeerMeta:   p.meta,
	}
}

var blockedCountries = []string{
	"AF", // Afghanistan
	"BY", // Belarus
	"CF", // Central African Republic
	"CU", // Cuba
	"IR", // Iran
	"KP", // North Korea
	"LB", // Lebanon
	"LY", // Libya
	"ML", // Mali
	"NI", // Nicaragua
	"PR", // Puerto Rico
	"RU", // Russia
	"SD", // Sudan
	"SS", // South Sudan
	"SY", // Syria
	"VE", // Venezuela
	"YE", // Yemen
	"ZW", // Zimbabwe
}

// filterBadIPs() returns the net address string and blocks any undesirable ip addresses
func (p *P2P) filterBadIPs(conn net.Conn) (netAddress string, e lib.ErrorI) {
	remoteAddr := conn.RemoteAddr()
	tcpAddr, ok := remoteAddr.(*net.TCPAddr)
	if !ok {
		return "", ErrNonTCPAddress()
	}
	host := tcpAddr.IP.String()
	ips, err := net.DefaultResolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return "", ErrIPLookup(err)
	}
	for _, ip := range ips {
		for _, bannedIP := range p.bannedIPs {
			if ip.IP.Equal(bannedIP.IP) {
				return "", ErrBannedIP(ip.String())
			}
		}
		originCountry := iploc.Country(ip.IP)
		if slices.Contains(blockedCountries, originCountry) {
			return "", ErrBannedCountry(originCountry)
		}
	}
	return net.JoinHostPort(host, fmt.Sprintf("%d", tcpAddr.Port)), nil
}

// catchPanic() is a programmatic safeguard against panics within the caller
func (p *P2P) catchPanic() {
	if r := recover(); r != nil {
		p.log.Error(string(debug.Stack()))
	}
}

// Dialing a thread safe map that logically maintains
// a count for peering chains that are being dialed
type Dialing struct {
	m map[uint64]int
	l sync.RWMutex
}

// AddDialing() increments the count for a list of chainIds that are currently being dialed
func (d *Dialing) AddDialing(chainId ...uint64) {
	d.l.Lock()
	defer d.l.Unlock()
	for _, chain := range chainId {
		d.m[chain]++
	}
}

// RemoveDialing() decrements the count for a list of chainIds that are currently being dialed
func (d *Dialing) RemoveDialing(chainId ...uint64) {
	d.l.Lock()
	defer d.l.Unlock()
	for _, chain := range chainId {
		d.m[chain]--
	}
}

// NumDialing() returns the count of peers currently being dialed for that id
func (d *Dialing) NumDialing(chainId uint64) int {
	d.l.RLock()
	defer d.l.RUnlock()
	return d.m[chainId]
}
