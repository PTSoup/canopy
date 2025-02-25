// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: block.proto

package lib

import (
	crypto "github.com/canopy-network/canopy/lib/crypto"
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

// *****************************************************************************************************
// This file is auto-generated from source files in `/lib/.proto/*` using Protocol Buffers (protobuf)
//
// Protobuf is a language-neutral, platform-neutral serialization format. It allows users
// to define objects in a way that’s both efficient to store and fast to transmit over the network.
// These definitions are compiled into code that *enables different systems and programming languages
// to communicate in a byte-perfect manner*
//
// To update these structures, make changes to the source .proto files, then recompile
// to regenerate this file.
// These auto-generated files are easily recognized by checking for a `.pb.go` ending
// *****************************************************************************************************
// _
// _
// _
// A block_header contains essential summary information about a block, like its unique identifier (hash), the previous
// block’s hash, timestamp, and block height. All of which help secure the block and connect it to the chain.
type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// height: the number of blocks in the blockchain before a specific block, indicating its position in the chain
	// and representing the blockchain's length
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// hash: a unique identifier for a block, generated by applying a cryptographic hash function to the block's contents,
	// ensuring its integrity and connecting it to the previous block in the chain.
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// network_id: a unique identifier used to distinguish different canopy blockchain networks, ensuring that
	// transactions and blocks are only processed within the correct network
	NetworkId uint32 `protobuf:"varint,3,opt,name=network_id,json=networkId,proto3" json:"networkID"` // @gotags: json:"networkID"
	// time: specific time recorded in a block when it is created by the proposer, indicating when the block was proposed
	Time uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	// num_txs: the count of transactions in the block
	NumTxs uint64 `protobuf:"varint,5,opt,name=num_txs,json=numTxs,proto3" json:"numTxs"` // @gotags: json:"numTxs"
	// total_txs: the count of transactions in the blockchain
	TotalTxs uint64 `protobuf:"varint,6,opt,name=total_txs,json=totalTxs,proto3" json:"totalTxs"` // @gotags: json:"totalTxs"
	// total_vdf_iterations: the total number of verifiable random delay function iterations in the blockchain
	TotalVdfIterations uint64 `protobuf:"varint,7,opt,name=total_vdf_iterations,json=totalVdfIterations,proto3" json:"totalVDFIterations"` // @gotags: json:"totalVDFIterations"
	// last_block_hash: the unique identifier of the previous block, chaining this block to the previous, ensuring the
	// continuity of the blockchain
	LastBlockHash []byte `protobuf:"bytes,8,opt,name=last_block_hash,json=lastBlockHash,proto3" json:"lastBlockHash"` // @gotags: json:"lastBlockHash"
	// state_root: the merkle root of the 'state commit store' representing  the entire state of the blockchain at
	// this height
	StateRoot []byte `protobuf:"bytes,9,opt,name=state_root,json=stateRoot,proto3" json:"stateRoot"` // @gotags: json:"stateRoot"
	// transaction_root: the merkle root of the 'transactions' included in this block
	TransactionRoot []byte `protobuf:"bytes,10,opt,name=transaction_root,json=transactionRoot,proto3" json:"transactionRoot"` // @gotags: json:"transactionRoot"
	// validator_root: the merkle root of the validators that signed the quorum certificate for this height
	// ensuring that the list of validators (who are responsible for proposing and validating blocks) is consistent
	// and tamper-proof, allowing participants to independently verify the integrity of the validator set at any point
	// in the blockchain
	ValidatorRoot []byte `protobuf:"bytes,11,opt,name=validator_root,json=validatorRoot,proto3" json:"validatorRoot"` // @gotags: json:"validatorRoot"
	// next_validator_root: the merkle root of the validators who are responsible for proposing and validating the next
	// block - ensuring smooth transitions between validator sets. This design allows lite-node operations as just with
	// the block_headers a node is able to validate the signers of the entire blockchain
	NextValidatorRoot []byte `protobuf:"bytes,12,opt,name=next_validator_root,json=nextValidatorRoot,proto3" json:"nextValidatorRoot"` // @gotags: json:"nextValidatorRoot"
	// proposer_address: is the short version of the public key of the Validator who proposed this block
	ProposerAddress []byte `protobuf:"bytes,13,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposerAddress"` // @gotags: json:"proposerAddress"
	// vdf: the verifiable delay proof for this block. The VDF serves as a protection mechanism against historical
	// forking attacks
	Vdf *crypto.VDF `protobuf:"bytes,14,opt,name=vdf,proto3" json:"vdf,omitempty"`
	// last_quorum_certificate: The quorum certificate from the previous block is included in the block header to ensure
	// all nodes have the same record of blockchain certificates, preventing discrepancies where the same block
	// could be accepted with different super-majorities of signatures across nodes.
	LastQuorumCertificate *QuorumCertificate `protobuf:"bytes,15,opt,name=last_quorum_certificate,json=lastQuorumCertificate,proto3" json:"lastQuorumCertificate"` // @gotags: json:"lastQuorumCertificate"
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHeader) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockHeader) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockHeader) GetNetworkId() uint32 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *BlockHeader) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *BlockHeader) GetNumTxs() uint64 {
	if x != nil {
		return x.NumTxs
	}
	return 0
}

func (x *BlockHeader) GetTotalTxs() uint64 {
	if x != nil {
		return x.TotalTxs
	}
	return 0
}

func (x *BlockHeader) GetTotalVdfIterations() uint64 {
	if x != nil {
		return x.TotalVdfIterations
	}
	return 0
}

func (x *BlockHeader) GetLastBlockHash() []byte {
	if x != nil {
		return x.LastBlockHash
	}
	return nil
}

func (x *BlockHeader) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *BlockHeader) GetTransactionRoot() []byte {
	if x != nil {
		return x.TransactionRoot
	}
	return nil
}

func (x *BlockHeader) GetValidatorRoot() []byte {
	if x != nil {
		return x.ValidatorRoot
	}
	return nil
}

func (x *BlockHeader) GetNextValidatorRoot() []byte {
	if x != nil {
		return x.NextValidatorRoot
	}
	return nil
}

func (x *BlockHeader) GetProposerAddress() []byte {
	if x != nil {
		return x.ProposerAddress
	}
	return nil
}

func (x *BlockHeader) GetVdf() *crypto.VDF {
	if x != nil {
		return x.Vdf
	}
	return nil
}

func (x *BlockHeader) GetLastQuorumCertificate() *QuorumCertificate {
	if x != nil {
		return x.LastQuorumCertificate
	}
	return nil
}

// A block is a batch of transactions and essential header information that marks a unit of committed information in the
// blockchain. Blocks are linked together in a sequence, creating a secure, verifiable chain of transaction history.
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// block_header: is the essential summary information of the block
	BlockHeader *BlockHeader `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"blockHeader"` // @gotags: json:"blockHeader"
	// transactions: is the batch of transactions in this block
	Transactions [][]byte `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetBlockHeader() *BlockHeader {
	if x != nil {
		return x.BlockHeader
	}
	return nil
}

func (x *Block) GetTransactions() [][]byte {
	if x != nil {
		return x.Transactions
	}
	return nil
}

// a block_result is a batch of transaction results, essential header information, and non-essential metadata about
// processing that block. A block result is made by applying a 'block' against a state machine.
type BlockResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// block_header: is the essential summary information of the block
	BlockHeader *BlockHeader `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"blockHeader"` // @gotags: json:"blockHeader"
	// transactions: is a batch of transaction results in this block
	Transactions []*TxResult `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// meta: is non-essential metadata about the processing of this block
	Meta *BlockResultMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *BlockResult) Reset() {
	*x = BlockResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResult) ProtoMessage() {}

func (x *BlockResult) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResult.ProtoReflect.Descriptor instead.
func (*BlockResult) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockResult) GetBlockHeader() *BlockHeader {
	if x != nil {
		return x.BlockHeader
	}
	return nil
}

func (x *BlockResult) GetTransactions() []*TxResult {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockResult) GetMeta() *BlockResultMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// block_result_meta is non-essential information about the processing of a block
type BlockResultMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// size: number of bytes in the block
	Size uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// took: duration string of the block in milliseconds
	Took uint64 `protobuf:"varint,2,opt,name=took,proto3" json:"took,omitempty"`
}

func (x *BlockResultMeta) Reset() {
	*x = BlockResultMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResultMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResultMeta) ProtoMessage() {}

func (x *BlockResultMeta) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResultMeta.ProtoReflect.Descriptor instead.
func (*BlockResultMeta) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{3}
}

func (x *BlockResultMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlockResultMeta) GetTook() uint64 {
	if x != nil {
		return x.Took
	}
	return 0
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x08, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb8, 0x04, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x78, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x54, 0x78, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x74, 0x78, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x54, 0x78, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x64,
	0x66, 0x5f, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x64, 0x66, 0x49, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x6e, 0x65,
	0x78, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x64,
	0x66, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x56, 0x44, 0x46, 0x52, 0x03, 0x76, 0x64, 0x66, 0x12, 0x50, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x62, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa5,
	0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35,
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x6f,
	0x6b, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63,
	0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2f, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_block_proto_goTypes = []interface{}{
	(*BlockHeader)(nil),       // 0: types.BlockHeader
	(*Block)(nil),             // 1: types.Block
	(*BlockResult)(nil),       // 2: types.BlockResult
	(*BlockResultMeta)(nil),   // 3: types.BlockResultMeta
	(*crypto.VDF)(nil),        // 4: types.VDF
	(*QuorumCertificate)(nil), // 5: types.QuorumCertificate
	(*TxResult)(nil),          // 6: types.TxResult
}
var file_block_proto_depIdxs = []int32{
	4, // 0: types.BlockHeader.vdf:type_name -> types.VDF
	5, // 1: types.BlockHeader.last_quorum_certificate:type_name -> types.QuorumCertificate
	0, // 2: types.Block.block_header:type_name -> types.BlockHeader
	0, // 3: types.BlockResult.block_header:type_name -> types.BlockHeader
	6, // 4: types.BlockResult.transactions:type_name -> types.TxResult
	3, // 5: types.BlockResult.meta:type_name -> types.BlockResultMeta
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	file_tx_proto_init()
	file_certificate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResult); i {
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
		file_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResultMeta); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
