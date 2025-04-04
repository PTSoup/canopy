// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: genesis.proto

package fsm

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
// GenesisState is an importable and exportable structure that is representative of the blockchain state at a
// certain height. The main use-cases for the Genesis state are acting as block 0 and analyzing the state db in
// json form. Please note that Block 0 is ofter referred to as the 'genesis' block
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time: the start time of the blockchain network
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// pools: a list of all the auto controlled accounts
	Pools []*Pool `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	// accounts: a list of all the user controlled accounts
	Accounts []*Account `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// validators: a list of all the staked validators
	Validators []*Validator `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators,omitempty"`
	// non_signers: a count of validators who haven't signed in some 'non-sign-window' of blocks
	// this field is not importable
	NonSigners []*NonSigner `protobuf:"bytes,5,rep,name=non_signers,json=nonSigners,proto3" json:"nonSigners"` // @gotags: json:"nonSigners"
	// double_signers: a list of those who 'double-signed' and the heights they double signed on
	// this field is not importable
	DoubleSigners []*lib.DoubleSigner `protobuf:"bytes,6,rep,name=double_signers,json=doubleSigners,proto3" json:"doubleSigners"` // @gotags: json:"doubleSigners"
	// order_books: is a list of active sell orders
	OrderBooks *lib.OrderBooks `protobuf:"bytes,7,opt,name=order_books,json=orderBooks,proto3" json:"orderBooks"` // @gotags: json:"orderBooks"
	// params: is a list of governance parameters and their respective values
	Params *Params `protobuf:"bytes,8,opt,name=params,proto3" json:"params,omitempty"`
	// supply: keeps track of the total amount of tokens or funds available across the entire blockchain
	// this field is not importable (auto generated using pools, accounts, validators etc.)
	Supply *Supply `protobuf:"bytes,9,opt,name=supply,proto3" json:"supply,omitempty"`
	// committees_data: is current status information about the current status of all committees
	// this field is not importable
	Committees *lib.CommitteesData `protobuf:"bytes,10,opt,name=committees,proto3" json:"committees,omitempty"`
	// retired committees: a list of the chain ids that are permanently retired
	RetiredCommittees []uint64 `protobuf:"varint,11,rep,packed,name=retired_committees,json=retiredCommittees,proto3" json:"retiredComittees"` // @gotags: json:"retiredComittees"
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *GenesisState) GetPools() []*Pool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *GenesisState) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *GenesisState) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GenesisState) GetNonSigners() []*NonSigner {
	if x != nil {
		return x.NonSigners
	}
	return nil
}

func (x *GenesisState) GetDoubleSigners() []*lib.DoubleSigner {
	if x != nil {
		return x.DoubleSigners
	}
	return nil
}

func (x *GenesisState) GetOrderBooks() *lib.OrderBooks {
	if x != nil {
		return x.OrderBooks
	}
	return nil
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetSupply() *Supply {
	if x != nil {
		return x.Supply
	}
	return nil
}

func (x *GenesisState) GetCommittees() *lib.CommitteesData {
	if x != nil {
		return x.Committees
	}
	return nil
}

func (x *GenesisState) GetRetiredCommittees() []uint64 {
	if x != nil {
		return x.RetiredCommittees
	}
	return nil
}

var File_genesis_proto protoreflect.FileDescriptor

var file_genesis_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x67, 0x6f, 0x76, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfa, 0x03, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x6e,
	0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x0d, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x25, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x04, 0x52, 0x11, 0x72, 0x65, 0x74, 0x69,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x73, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f,
	0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70,
	0x79, 0x2f, 0x66, 0x73, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genesis_proto_rawDescOnce sync.Once
	file_genesis_proto_rawDescData = file_genesis_proto_rawDesc
)

func file_genesis_proto_rawDescGZIP() []byte {
	file_genesis_proto_rawDescOnce.Do(func() {
		file_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_genesis_proto_rawDescData)
	})
	return file_genesis_proto_rawDescData
}

var file_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),       // 0: types.GenesisState
	(*Pool)(nil),               // 1: types.Pool
	(*Account)(nil),            // 2: types.Account
	(*Validator)(nil),          // 3: types.Validator
	(*NonSigner)(nil),          // 4: types.NonSigner
	(*lib.DoubleSigner)(nil),   // 5: types.DoubleSigner
	(*lib.OrderBooks)(nil),     // 6: types.OrderBooks
	(*Params)(nil),             // 7: types.Params
	(*Supply)(nil),             // 8: types.Supply
	(*lib.CommitteesData)(nil), // 9: types.CommitteesData
}
var file_genesis_proto_depIdxs = []int32{
	1, // 0: types.GenesisState.pools:type_name -> types.Pool
	2, // 1: types.GenesisState.accounts:type_name -> types.Account
	3, // 2: types.GenesisState.validators:type_name -> types.Validator
	4, // 3: types.GenesisState.non_signers:type_name -> types.NonSigner
	5, // 4: types.GenesisState.double_signers:type_name -> types.DoubleSigner
	6, // 5: types.GenesisState.order_books:type_name -> types.OrderBooks
	7, // 6: types.GenesisState.params:type_name -> types.Params
	8, // 7: types.GenesisState.supply:type_name -> types.Supply
	9, // 8: types.GenesisState.committees:type_name -> types.CommitteesData
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_genesis_proto_init() }
func file_genesis_proto_init() {
	if File_genesis_proto != nil {
		return
	}
	file_account_proto_init()
	file_validator_proto_init()
	file_gov_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genesis_proto_goTypes,
		DependencyIndexes: file_genesis_proto_depIdxs,
		MessageInfos:      file_genesis_proto_msgTypes,
	}.Build()
	File_genesis_proto = out.File
	file_genesis_proto_rawDesc = nil
	file_genesis_proto_goTypes = nil
	file_genesis_proto_depIdxs = nil
}
