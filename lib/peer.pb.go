// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: peer.proto

package lib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ordered by stream priority
type Topic int32

const (
	Topic_CONSENSUS           Topic = 0
	Topic_CERTIFICATE         Topic = 1
	Topic_CERTIFICATE_REQUEST Topic = 2
	Topic_TX                  Topic = 3
	Topic_PEERS_RESPONSE      Topic = 4
	Topic_PEERS_REQUEST       Topic = 5
	Topic_INVALID             Topic = 6
)

// Enum value maps for Topic.
var (
	Topic_name = map[int32]string{
		0: "CONSENSUS",
		1: "CERTIFICATE",
		2: "CERTIFICATE_REQUEST",
		3: "TX",
		4: "PEERS_RESPONSE",
		5: "PEERS_REQUEST",
		6: "INVALID",
	}
	Topic_value = map[string]int32{
		"CONSENSUS":           0,
		"CERTIFICATE":         1,
		"CERTIFICATE_REQUEST": 2,
		"TX":                  3,
		"PEERS_RESPONSE":      4,
		"PEERS_REQUEST":       5,
		"INVALID":             6,
	}
)

func (x Topic) Enum() *Topic {
	p := new(Topic)
	*p = x
	return p
}

func (x Topic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Topic) Descriptor() protoreflect.EnumDescriptor {
	return file_peer_proto_enumTypes[0].Descriptor()
}

func (Topic) Type() protoreflect.EnumType {
	return &file_peer_proto_enumTypes[0]
}

func (x Topic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Topic.Descriptor instead.
func (Topic) EnumDescriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       *PeerAddress `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	IsOutbound    bool         `protobuf:"varint,2,opt,name=is_outbound,json=isOutbound,proto3" json:"is_outbound,omitempty"`
	IsMustConnect bool         `protobuf:"varint,3,opt,name=is_must_connect,json=isMustConnect,proto3" json:"is_must_connect,omitempty"`
	IsTrusted     bool         `protobuf:"varint,4,opt,name=is_trusted,json=isTrusted,proto3" json:"is_trusted,omitempty"`
	Reputation    int32        `protobuf:"varint,5,opt,name=reputation,proto3" json:"reputation,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

func (x *PeerInfo) GetAddress() *PeerAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PeerInfo) GetIsOutbound() bool {
	if x != nil {
		return x.IsOutbound
	}
	return false
}

func (x *PeerInfo) GetIsMustConnect() bool {
	if x != nil {
		return x.IsMustConnect
	}
	return false
}

func (x *PeerInfo) GetIsTrusted() bool {
	if x != nil {
		return x.IsTrusted
	}
	return false
}

func (x *PeerInfo) GetReputation() int32 {
	if x != nil {
		return x.Reputation
	}
	return 0
}

type PeerAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey  []byte    `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	NetAddress string    `protobuf:"bytes,2,opt,name=net_address,json=netAddress,proto3" json:"net_address,omitempty"`
	PeerMeta   *PeerMeta `protobuf:"bytes,3,opt,name=peer_meta,json=peerMeta,proto3" json:"peer_meta,omitempty"`
}

func (x *PeerAddress) Reset() {
	*x = PeerAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAddress) ProtoMessage() {}

func (x *PeerAddress) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerAddress.ProtoReflect.Descriptor instead.
func (*PeerAddress) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{1}
}

func (x *PeerAddress) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *PeerAddress) GetNetAddress() string {
	if x != nil {
		return x.NetAddress
	}
	return ""
}

func (x *PeerAddress) GetPeerMeta() *PeerMeta {
	if x != nil {
		return x.PeerMeta
	}
	return nil
}

type PeerMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId uint64   `protobuf:"varint,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Chains    []uint64 `protobuf:"varint,2,rep,packed,name=chains,proto3" json:"chains,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *PeerMeta) Reset() {
	*x = PeerMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerMeta) ProtoMessage() {}

func (x *PeerMeta) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerMeta.ProtoReflect.Descriptor instead.
func (*PeerMeta) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{2}
}

func (x *PeerMeta) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *PeerMeta) GetChains() []uint64 {
	if x != nil {
		return x.Chains
	}
	return nil
}

func (x *PeerMeta) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ConsensusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitteeId uint64 `protobuf:"varint,1,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	Message     []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConsensusMessage) Reset() {
	*x = ConsensusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusMessage) ProtoMessage() {}

func (x *ConsensusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusMessage.ProtoReflect.Descriptor instead.
func (*ConsensusMessage) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{3}
}

func (x *ConsensusMessage) GetCommitteeId() uint64 {
	if x != nil {
		return x.CommitteeId
	}
	return 0
}

func (x *ConsensusMessage) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type BlockRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitteeId uint64 `protobuf:"varint,1,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	Height      uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	HeightOnly  bool   `protobuf:"varint,3,opt,name=height_only,json=heightOnly,proto3" json:"height_only,omitempty"`
}

func (x *BlockRequestMessage) Reset() {
	*x = BlockRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequestMessage) ProtoMessage() {}

func (x *BlockRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequestMessage.ProtoReflect.Descriptor instead.
func (*BlockRequestMessage) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{4}
}

func (x *BlockRequestMessage) GetCommitteeId() uint64 {
	if x != nil {
		return x.CommitteeId
	}
	return 0
}

func (x *BlockRequestMessage) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockRequestMessage) GetHeightOnly() bool {
	if x != nil {
		return x.HeightOnly
	}
	return false
}

type BlockResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitteeId         uint64             `protobuf:"varint,1,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	MaxHeight           uint64             `protobuf:"varint,2,opt,name=max_height,json=maxHeight,proto3" json:"max_height,omitempty"`
	TotalVdfIterations  uint64             `protobuf:"varint,3,opt,name=total_vdf_iterations,json=totalVdfIterations,proto3" json:"total_vdf_iterations,omitempty"`
	BlockAndCertificate *QuorumCertificate `protobuf:"bytes,4,opt,name=BlockAndCertificate,proto3" json:"BlockAndCertificate,omitempty"`
}

func (x *BlockResponseMessage) Reset() {
	*x = BlockResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponseMessage) ProtoMessage() {}

func (x *BlockResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponseMessage.ProtoReflect.Descriptor instead.
func (*BlockResponseMessage) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{5}
}

func (x *BlockResponseMessage) GetCommitteeId() uint64 {
	if x != nil {
		return x.CommitteeId
	}
	return 0
}

func (x *BlockResponseMessage) GetMaxHeight() uint64 {
	if x != nil {
		return x.MaxHeight
	}
	return 0
}

func (x *BlockResponseMessage) GetTotalVdfIterations() uint64 {
	if x != nil {
		return x.TotalVdfIterations
	}
	return 0
}

func (x *BlockResponseMessage) GetBlockAndCertificate() *QuorumCertificate {
	if x != nil {
		return x.BlockAndCertificate
	}
	return nil
}

type TxMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitteeId uint64 `protobuf:"varint,1,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	Tx          []byte `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *TxMessage) Reset() {
	*x = TxMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxMessage) ProtoMessage() {}

func (x *TxMessage) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxMessage.ProtoReflect.Descriptor instead.
func (*TxMessage) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{6}
}

func (x *TxMessage) GetCommitteeId() uint64 {
	if x != nil {
		return x.CommitteeId
	}
	return 0
}

func (x *TxMessage) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

var File_peer_proto protoreflect.FileDescriptor

var file_peer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x4d, 0x75, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4f, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x71, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x64,
	0x66, 0x5f, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x64, 0x66, 0x49, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x13, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x22, 0x3e, 0x0a, 0x09, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x74, 0x78, 0x2a, 0x7c, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4e, 0x53, 0x45, 0x4e, 0x53, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x45,
	0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43,
	0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x58, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x45, 0x45, 0x52, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x50, 0x45, 0x45, 0x52, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x06,
	0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x6e, 0x63, 0x68, 0x75, 0x63, 0x6f, 0x2f, 0x67, 0x69, 0x6e, 0x63, 0x68, 0x75, 0x2f, 0x6c,
	0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proto_rawDescOnce sync.Once
	file_peer_proto_rawDescData = file_peer_proto_rawDesc
)

func file_peer_proto_rawDescGZIP() []byte {
	file_peer_proto_rawDescOnce.Do(func() {
		file_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proto_rawDescData)
	})
	return file_peer_proto_rawDescData
}

var file_peer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_peer_proto_goTypes = []interface{}{
	(Topic)(0),                   // 0: types.Topic
	(*PeerInfo)(nil),             // 1: types.PeerInfo
	(*PeerAddress)(nil),          // 2: types.PeerAddress
	(*PeerMeta)(nil),             // 3: types.PeerMeta
	(*ConsensusMessage)(nil),     // 4: types.ConsensusMessage
	(*BlockRequestMessage)(nil),  // 5: types.BlockRequestMessage
	(*BlockResponseMessage)(nil), // 6: types.BlockResponseMessage
	(*TxMessage)(nil),            // 7: types.TxMessage
	(*QuorumCertificate)(nil),    // 8: types.QuorumCertificate
}
var file_peer_proto_depIdxs = []int32{
	2, // 0: types.PeerInfo.Address:type_name -> types.PeerAddress
	3, // 1: types.PeerAddress.peer_meta:type_name -> types.PeerMeta
	8, // 2: types.BlockResponseMessage.BlockAndCertificate:type_name -> types.QuorumCertificate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_peer_proto_init() }
func file_peer_proto_init() {
	if File_peer_proto != nil {
		return
	}
	file_consensus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_proto_goTypes,
		DependencyIndexes: file_peer_proto_depIdxs,
		EnumInfos:         file_peer_proto_enumTypes,
		MessageInfos:      file_peer_proto_msgTypes,
	}.Build()
	File_peer_proto = out.File
	file_peer_proto_rawDesc = nil
	file_peer_proto_goTypes = nil
	file_peer_proto_depIdxs = nil
}
