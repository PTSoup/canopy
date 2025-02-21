// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: swap.proto

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
// SWAP PROTOCOL: Facilitates a token swap between Bob (with Token A) and Alice (with Token B).
// The committee oversees the process while controlling Blockchain A and observing Blockchain B.
//
//  1. Bob creates a 'SellOrder' with the amount of "Token A" he wants to sell, the 'exchange rate', and his
//     'Token B address'. Token A is escrowed in a committee-controlled address. Bob can reverse this order by submitting
//     a transaction on Blockchain A.
//
//  2. Alice accepts Bob's offer by sending a transaction on Blockchain B, referencing Bob’s offer hash and providing her
//     'Token A address' in the memo field.
//
//  3. The committee updates the recipient of Bob’s sell order to Alice’s "Token A" address, verifying that Alice has
//     enough "aged Token B" in her Blockchain B address.
//
// 4. Alice sends "Token B" to Bob, with a memo linking to the 'Request to Sell'.
//
// 5. The committee witnesses Alice’s transaction and releases Bob’s "Token A" to Alice.
//
//  6. If Alice does not send "Token B" within N Blockchain B blocks, the committee resets Bob’s 'Request to Sell'
//     recipient.
//
// SellOrder is a structure that holds relevant data to complete a token swap. It's created by a 'request to sell'
// then populated by an 'intent to buy', and finally closed when the committee witnesses the transfer of funds.
type SellOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id: the unique identifier of the order
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"id"` // @gotags: json:"id"
	// committee: the id of the committee that is in-charge of escrow for the swap
	Committee uint64 `protobuf:"varint,2,opt,name=Committee,proto3" json:"committee"` // @gotags: json:"committee"
	// amount_for_sale: amount of CNPY for sale
	AmountForSale uint64 `protobuf:"varint,3,opt,name=AmountForSale,proto3" json:"amountForSalee"` // @gotags: json:"amountForSalee"
	// requested_amount: amount of 'counter-asset' to receive
	RequestedAmount uint64 `protobuf:"varint,4,opt,name=RequestedAmount,proto3" json:"requestedAmount"` // @gotags: json:"requestedAmount"
	// seller_receive_address: the external chain address to receive the 'counter-asset'
	SellerReceiveAddress []byte `protobuf:"bytes,5,opt,name=SellerReceiveAddress,proto3" json:"sellerReceiveAddress"` // @gotags: json:"sellerReceiveAddress"
	// buyer_send_address: the address the buyer will be transferring the funds from
	BuyerSendAddress []byte `protobuf:"bytes,6,opt,name=BuyerSendAddress,proto3" json:"buyerSendAddress"` // @gotags: json:"buyerSendAddress"
	// buyer_receive_address: the buyer Canopy address to receive the CNPY
	BuyerReceiveAddress []byte `protobuf:"bytes,7,opt,name=BuyerReceiveAddress,proto3" json:"buyerReceiveAddress"` // @gotags: json:"buyerReceiveAddress"
	// buyer_chain_deadline: the external chain height deadline to send the 'tokens' to SellerReceiveAddress
	BuyerChainDeadline uint64 `protobuf:"varint,8,opt,name=BuyerChainDeadline,proto3" json:"buyerChainDeadline"` // @gotags: json:"buyerChainDeadline"
	// sellers_send_address: the signing address of seller who is selling the CNPY
	SellersSendAddress []byte `protobuf:"bytes,9,opt,name=SellersSendAddress,proto3" json:"sellersSendAddress"` // @gotags: json:"sellersSendAddress"
}

func (x *SellOrder) Reset() {
	*x = SellOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellOrder) ProtoMessage() {}

func (x *SellOrder) ProtoReflect() protoreflect.Message {
	mi := &file_swap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellOrder.ProtoReflect.Descriptor instead.
func (*SellOrder) Descriptor() ([]byte, []int) {
	return file_swap_proto_rawDescGZIP(), []int{0}
}

func (x *SellOrder) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SellOrder) GetCommittee() uint64 {
	if x != nil {
		return x.Committee
	}
	return 0
}

func (x *SellOrder) GetAmountForSale() uint64 {
	if x != nil {
		return x.AmountForSale
	}
	return 0
}

func (x *SellOrder) GetRequestedAmount() uint64 {
	if x != nil {
		return x.RequestedAmount
	}
	return 0
}

func (x *SellOrder) GetSellerReceiveAddress() []byte {
	if x != nil {
		return x.SellerReceiveAddress
	}
	return nil
}

func (x *SellOrder) GetBuyerSendAddress() []byte {
	if x != nil {
		return x.BuyerSendAddress
	}
	return nil
}

func (x *SellOrder) GetBuyerReceiveAddress() []byte {
	if x != nil {
		return x.BuyerReceiveAddress
	}
	return nil
}

func (x *SellOrder) GetBuyerChainDeadline() uint64 {
	if x != nil {
		return x.BuyerChainDeadline
	}
	return 0
}

func (x *SellOrder) GetSellersSendAddress() []byte {
	if x != nil {
		return x.SellersSendAddress
	}
	return nil
}

// OrderBooks: is a list of order book objects held in the blockchain state
type OrderBooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OrderBooks: the actual list of order book objects
	OrderBooks []*OrderBook `protobuf:"bytes,1,rep,name=OrderBooks,proto3" json:"orderBooks"` // @gotags: json:"orderBooks"
}

func (x *OrderBooks) Reset() {
	*x = OrderBooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBooks) ProtoMessage() {}

func (x *OrderBooks) ProtoReflect() protoreflect.Message {
	mi := &file_swap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBooks.ProtoReflect.Descriptor instead.
func (*OrderBooks) Descriptor() ([]byte, []int) {
	return file_swap_proto_rawDescGZIP(), []int{1}
}

func (x *OrderBooks) GetOrderBooks() []*OrderBook {
	if x != nil {
		return x.OrderBooks
	}
	return nil
}

// OrderBook: a list of sell orders associated with a particular committee held in the blockchain state
type OrderBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// chain_id: the unique identifier of the 'counter asset' committee
	ChainId uint64 `protobuf:"varint,1,opt,name=chainId,proto3" json:"chainID"` // @gotags: json:"chainID"
	// orders: the actual list of sell orders
	Orders []*SellOrder `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderBook) Reset() {
	*x = OrderBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBook) ProtoMessage() {}

func (x *OrderBook) ProtoReflect() protoreflect.Message {
	mi := &file_swap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBook.ProtoReflect.Descriptor instead.
func (*OrderBook) Descriptor() ([]byte, []int) {
	return file_swap_proto_rawDescGZIP(), []int{2}
}

func (x *OrderBook) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *OrderBook) GetOrders() []*SellOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_swap_proto protoreflect.FileDescriptor

var file_swap_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0xfb, 0x02, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x53, 0x61, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f,
	0x72, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x32, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x42, 0x75, 0x79, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x30, 0x0a, 0x13, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2e, 0x0a, 0x12, 0x42, 0x75, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x53,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x3e, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12,
	0x30, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x22, 0x4f, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2f, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_swap_proto_rawDescOnce sync.Once
	file_swap_proto_rawDescData = file_swap_proto_rawDesc
)

func file_swap_proto_rawDescGZIP() []byte {
	file_swap_proto_rawDescOnce.Do(func() {
		file_swap_proto_rawDescData = protoimpl.X.CompressGZIP(file_swap_proto_rawDescData)
	})
	return file_swap_proto_rawDescData
}

var file_swap_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_swap_proto_goTypes = []interface{}{
	(*SellOrder)(nil),  // 0: types.SellOrder
	(*OrderBooks)(nil), // 1: types.OrderBooks
	(*OrderBook)(nil),  // 2: types.OrderBook
}
var file_swap_proto_depIdxs = []int32{
	2, // 0: types.OrderBooks.OrderBooks:type_name -> types.OrderBook
	0, // 1: types.OrderBook.orders:type_name -> types.SellOrder
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_swap_proto_init() }
func file_swap_proto_init() {
	if File_swap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_swap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellOrder); i {
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
		file_swap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBooks); i {
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
		file_swap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBook); i {
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
			RawDescriptor: file_swap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_swap_proto_goTypes,
		DependencyIndexes: file_swap_proto_depIdxs,
		MessageInfos:      file_swap_proto_msgTypes,
	}.Build()
	File_swap_proto = out.File
	file_swap_proto_rawDesc = nil
	file_swap_proto_goTypes = nil
	file_swap_proto_depIdxs = nil
}
