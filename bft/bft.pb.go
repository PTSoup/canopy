// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: bft.proto

package bft

import (
	lib "github.com/canopy-network/canopy/lib"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header                 *lib.View              `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Vrf                    *lib.Signature         `protobuf:"bytes,2,opt,name=vrf,proto3" json:"vrf,omitempty"`                                                                         // only ELECTION phase from proposer
	Qc                     *lib.QuorumCertificate `protobuf:"bytes,3,opt,name=qc,proto3" json:"qc,omitempty"`                                                                           // Proposers use this to aggregate replica messages, Replicas sign parts of this to vote
	HighQc                 *lib.QuorumCertificate `protobuf:"bytes,4,opt,name=high_qc,json=highQc,proto3" json:"high_qc,omitempty"`                                                     // only if a previous lock was seen: ELECTION-VOTE phase by replica and PROPOSE phase by proposer
	LastDoubleSignEvidence []*DoubleSignEvidence  `protobuf:"bytes,5,rep,name=last_double_sign_evidence,json=lastDoubleSignEvidence,proto3" json:"last_double_sign_evidence,omitempty"` // evidence from height - 1 (required cause last round ds would not be caught otherwise)
	BadProposerEvidence    []*BadProposerEvidence `protobuf:"bytes,6,rep,name=bad_proposer_evidence,json=badProposerEvidence,proto3" json:"bad_proposer_evidence,omitempty"`            // evidence from current height (last round is always good, so we can do same height)
	Vdf                    *lib.VDF               `protobuf:"bytes,7,opt,name=vdf,proto3" json:"vdf,omitempty"`
	Signature              *lib.Signature         `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_bft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_bft_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetHeader() *lib.View {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetVrf() *lib.Signature {
	if x != nil {
		return x.Vrf
	}
	return nil
}

func (x *Message) GetQc() *lib.QuorumCertificate {
	if x != nil {
		return x.Qc
	}
	return nil
}

func (x *Message) GetHighQc() *lib.QuorumCertificate {
	if x != nil {
		return x.HighQc
	}
	return nil
}

func (x *Message) GetLastDoubleSignEvidence() []*DoubleSignEvidence {
	if x != nil {
		return x.LastDoubleSignEvidence
	}
	return nil
}

func (x *Message) GetBadProposerEvidence() []*BadProposerEvidence {
	if x != nil {
		return x.BadProposerEvidence
	}
	return nil
}

func (x *Message) GetVdf() *lib.VDF {
	if x != nil {
		return x.Vdf
	}
	return nil
}

func (x *Message) GetSignature() *lib.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

type DoubleSignEvidences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence     []*DoubleSignEvidence `protobuf:"bytes,1,rep,name=Evidence,proto3" json:"Evidence,omitempty"`
	DeDuplicator map[string]bool       `protobuf:"bytes,2,rep,name=DeDuplicator,proto3" json:"DeDuplicator,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *DoubleSignEvidences) Reset() {
	*x = DoubleSignEvidences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSignEvidences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSignEvidences) ProtoMessage() {}

func (x *DoubleSignEvidences) ProtoReflect() protoreflect.Message {
	mi := &file_bft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSignEvidences.ProtoReflect.Descriptor instead.
func (*DoubleSignEvidences) Descriptor() ([]byte, []int) {
	return file_bft_proto_rawDescGZIP(), []int{1}
}

func (x *DoubleSignEvidences) GetEvidence() []*DoubleSignEvidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *DoubleSignEvidences) GetDeDuplicator() map[string]bool {
	if x != nil {
		return x.DeDuplicator
	}
	return nil
}

type DoubleSignEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteA *lib.QuorumCertificate `protobuf:"bytes,1,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB *lib.QuorumCertificate `protobuf:"bytes,2,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
}

func (x *DoubleSignEvidence) Reset() {
	*x = DoubleSignEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSignEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSignEvidence) ProtoMessage() {}

func (x *DoubleSignEvidence) ProtoReflect() protoreflect.Message {
	mi := &file_bft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSignEvidence.ProtoReflect.Descriptor instead.
func (*DoubleSignEvidence) Descriptor() ([]byte, []int) {
	return file_bft_proto_rawDescGZIP(), []int{2}
}

func (x *DoubleSignEvidence) GetVoteA() *lib.QuorumCertificate {
	if x != nil {
		return x.VoteA
	}
	return nil
}

func (x *DoubleSignEvidence) GetVoteB() *lib.QuorumCertificate {
	if x != nil {
		return x.VoteB
	}
	return nil
}

type BadProposerEvidences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence     []*BadProposerEvidence `protobuf:"bytes,1,rep,name=Evidence,proto3" json:"Evidence,omitempty"`
	DeDuplicator map[string]bool        `protobuf:"bytes,2,rep,name=DeDuplicator,proto3" json:"DeDuplicator,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *BadProposerEvidences) Reset() {
	*x = BadProposerEvidences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadProposerEvidences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadProposerEvidences) ProtoMessage() {}

func (x *BadProposerEvidences) ProtoReflect() protoreflect.Message {
	mi := &file_bft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadProposerEvidences.ProtoReflect.Descriptor instead.
func (*BadProposerEvidences) Descriptor() ([]byte, []int) {
	return file_bft_proto_rawDescGZIP(), []int{3}
}

func (x *BadProposerEvidences) GetEvidence() []*BadProposerEvidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

func (x *BadProposerEvidences) GetDeDuplicator() map[string]bool {
	if x != nil {
		return x.DeDuplicator
	}
	return nil
}

type BadProposerEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionVoteQc *lib.QuorumCertificate `protobuf:"bytes,1,opt,name=election_vote_qc,json=electionVoteQc,proto3" json:"election_vote_qc,omitempty"`
}

func (x *BadProposerEvidence) Reset() {
	*x = BadProposerEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadProposerEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadProposerEvidence) ProtoMessage() {}

func (x *BadProposerEvidence) ProtoReflect() protoreflect.Message {
	mi := &file_bft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadProposerEvidence.ProtoReflect.Descriptor instead.
func (*BadProposerEvidence) Descriptor() ([]byte, []int) {
	return file_bft_proto_rawDescGZIP(), []int{4}
}

func (x *BadProposerEvidence) GetElectionVoteQc() *lib.QuorumCertificate {
	if x != nil {
		return x.ElectionVoteQc
	}
	return nil
}

var File_bft_proto protoreflect.FileDescriptor

var file_bft_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x08, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x03,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x03, 0x76, 0x72, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x03, 0x76,
	0x72, 0x66, 0x12, 0x28, 0x0a, 0x02, 0x71, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x02, 0x71, 0x63, 0x12, 0x31, 0x0a, 0x07,
	0x68, 0x69, 0x67, 0x68, 0x5f, 0x71, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x06, 0x68, 0x69, 0x67, 0x68, 0x51, 0x63, 0x12,
	0x54, 0x0a, 0x19, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x16, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x62, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x64,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x13, 0x62, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x64, 0x66, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x44, 0x46, 0x52, 0x03,
	0x76, 0x64, 0x66, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x13, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x45, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x76, 0x0a, 0x12, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x12, 0x2f, 0x0a, 0x06,
	0x76, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x22, 0xe2, 0x01,
	0x0a, 0x14, 0x42, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x0c, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x64,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x1a, 0x3f, 0x0a, 0x11, 0x44, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x59, 0x0a, 0x13, 0x42, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x71, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x51, 0x63, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f,
	0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70,
	0x79, 0x2f, 0x62, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bft_proto_rawDescOnce sync.Once
	file_bft_proto_rawDescData = file_bft_proto_rawDesc
)

func file_bft_proto_rawDescGZIP() []byte {
	file_bft_proto_rawDescOnce.Do(func() {
		file_bft_proto_rawDescData = protoimpl.X.CompressGZIP(file_bft_proto_rawDescData)
	})
	return file_bft_proto_rawDescData
}

var file_bft_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bft_proto_goTypes = []interface{}{
	(*Message)(nil),               // 0: types.Message
	(*DoubleSignEvidences)(nil),   // 1: types.DoubleSignEvidences
	(*DoubleSignEvidence)(nil),    // 2: types.DoubleSignEvidence
	(*BadProposerEvidences)(nil),  // 3: types.BadProposerEvidences
	(*BadProposerEvidence)(nil),   // 4: types.BadProposerEvidence
	nil,                           // 5: types.DoubleSignEvidences.DeDuplicatorEntry
	nil,                           // 6: types.BadProposerEvidences.DeDuplicatorEntry
	(*lib.View)(nil),              // 7: types.View
	(*lib.Signature)(nil),         // 8: types.Signature
	(*lib.QuorumCertificate)(nil), // 9: types.QuorumCertificate
	(*lib.VDF)(nil),               // 10: types.VDF
}
var file_bft_proto_depIdxs = []int32{
	7,  // 0: types.Message.header:type_name -> types.View
	8,  // 1: types.Message.vrf:type_name -> types.Signature
	9,  // 2: types.Message.qc:type_name -> types.QuorumCertificate
	9,  // 3: types.Message.high_qc:type_name -> types.QuorumCertificate
	2,  // 4: types.Message.last_double_sign_evidence:type_name -> types.DoubleSignEvidence
	4,  // 5: types.Message.bad_proposer_evidence:type_name -> types.BadProposerEvidence
	10, // 6: types.Message.vdf:type_name -> types.VDF
	8,  // 7: types.Message.signature:type_name -> types.Signature
	2,  // 8: types.DoubleSignEvidences.Evidence:type_name -> types.DoubleSignEvidence
	5,  // 9: types.DoubleSignEvidences.DeDuplicator:type_name -> types.DoubleSignEvidences.DeDuplicatorEntry
	9,  // 10: types.DoubleSignEvidence.vote_a:type_name -> types.QuorumCertificate
	9,  // 11: types.DoubleSignEvidence.vote_b:type_name -> types.QuorumCertificate
	4,  // 12: types.BadProposerEvidences.Evidence:type_name -> types.BadProposerEvidence
	6,  // 13: types.BadProposerEvidences.DeDuplicator:type_name -> types.BadProposerEvidences.DeDuplicatorEntry
	9,  // 14: types.BadProposerEvidence.election_vote_qc:type_name -> types.QuorumCertificate
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_bft_proto_init() }
func file_bft_proto_init() {
	if File_bft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_bft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSignEvidences); i {
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
		file_bft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSignEvidence); i {
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
		file_bft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadProposerEvidences); i {
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
		file_bft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadProposerEvidence); i {
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
			RawDescriptor: file_bft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bft_proto_goTypes,
		DependencyIndexes: file_bft_proto_depIdxs,
		MessageInfos:      file_bft_proto_msgTypes,
	}.Build()
	File_bft_proto = out.File
	file_bft_proto_rawDesc = nil
	file_bft_proto_goTypes = nil
	file_bft_proto_depIdxs = nil
}
