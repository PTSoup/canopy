// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: gov.proto

package types

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

type GovProposalVoteConfig int32

const (
	GovProposalVoteConfig_ACCEPT_ALL   GovProposalVoteConfig = 0
	GovProposalVoteConfig_APPROVE_LIST GovProposalVoteConfig = 1
	GovProposalVoteConfig_REJECT_ALL   GovProposalVoteConfig = 2
)

// Enum value maps for GovProposalVoteConfig.
var (
	GovProposalVoteConfig_name = map[int32]string{
		0: "ACCEPT_ALL",
		1: "APPROVE_LIST",
		2: "REJECT_ALL",
	}
	GovProposalVoteConfig_value = map[string]int32{
		"ACCEPT_ALL":   0,
		"APPROVE_LIST": 1,
		"REJECT_ALL":   2,
	}
)

func (x GovProposalVoteConfig) Enum() *GovProposalVoteConfig {
	p := new(GovProposalVoteConfig)
	*p = x
	return p
}

func (x GovProposalVoteConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GovProposalVoteConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_gov_proto_enumTypes[0].Descriptor()
}

func (GovProposalVoteConfig) Type() protoreflect.EnumType {
	return &file_gov_proto_enumTypes[0]
}

func (x GovProposalVoteConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GovProposalVoteConfig.Descriptor instead.
func (GovProposalVoteConfig) EnumDescriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{0}
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consensus  *ConsensusParams  `protobuf:"bytes,1,opt,name=Consensus,proto3" json:"Consensus,omitempty"`
	Validator  *ValidatorParams  `protobuf:"bytes,2,opt,name=Validator,proto3" json:"Validator,omitempty"`
	Fee        *FeeParams        `protobuf:"bytes,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Governance *GovernanceParams `protobuf:"bytes,4,opt,name=Governance,proto3" json:"Governance,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetConsensus() *ConsensusParams {
	if x != nil {
		return x.Consensus
	}
	return nil
}

func (x *Params) GetValidator() *ValidatorParams {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *Params) GetFee() *FeeParams {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Params) GetGovernance() *GovernanceParams {
	if x != nil {
		return x.Governance
	}
	return nil
}

type ProtocolVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height  uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ProtocolVersion) Reset() {
	*x = ProtocolVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersion) ProtoMessage() {}

func (x *ProtocolVersion) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolVersion.ProtoReflect.Descriptor instead.
func (*ProtocolVersion) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{1}
}

func (x *ProtocolVersion) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ProtocolVersion) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type ConsensusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockSize       uint64 `protobuf:"varint,1,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	ProtocolVersion string `protobuf:"bytes,2,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
}

func (x *ConsensusParams) Reset() {
	*x = ConsensusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusParams) ProtoMessage() {}

func (x *ConsensusParams) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusParams.ProtoReflect.Descriptor instead.
func (*ConsensusParams) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{2}
}

func (x *ConsensusParams) GetBlockSize() uint64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *ConsensusParams) GetProtocolVersion() string {
	if x != nil {
		return x.ProtocolVersion
	}
	return ""
}

type ValidatorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorUnstakingBlocks                uint64 `protobuf:"varint,1,opt,name=validator_unstaking_blocks,json=validatorUnstakingBlocks,proto3" json:"validator_unstaking_blocks,omitempty"`
	ValidatorMaxPauseBlocks                 uint64 `protobuf:"varint,2,opt,name=validator_max_pause_blocks,json=validatorMaxPauseBlocks,proto3" json:"validator_max_pause_blocks,omitempty"`
	ValidatorDoubleSignSlashPercentage      uint64 `protobuf:"varint,3,opt,name=validator_double_sign_slash_percentage,json=validatorDoubleSignSlashPercentage,proto3" json:"validator_double_sign_slash_percentage,omitempty"`
	ValidatorBadProposalSlashPercentage     uint64 `protobuf:"varint,4,opt,name=validator_bad_proposal_slash_percentage,json=validatorBadProposalSlashPercentage,proto3" json:"validator_bad_proposal_slash_percentage,omitempty"`
	ValidatorNonSignSlashPercentage         uint64 `protobuf:"varint,5,opt,name=validator_non_sign_slash_percentage,json=validatorNonSignSlashPercentage,proto3" json:"validator_non_sign_slash_percentage,omitempty"`
	ValidatorMaxNonSign                     uint64 `protobuf:"varint,6,opt,name=validator_max_non_sign,json=validatorMaxNonSign,proto3" json:"validator_max_non_sign,omitempty"`
	ValidatorNonSignWindow                  uint64 `protobuf:"varint,7,opt,name=validator_non_sign_window,json=validatorNonSignWindow,proto3" json:"validator_non_sign_window,omitempty"`
	ValidatorMaxCommittees                  uint64 `protobuf:"varint,8,opt,name=validator_max_committees,json=validatorMaxCommittees,proto3" json:"validator_max_committees,omitempty"`
	ValidatorMaxCommitteeSize               uint64 `protobuf:"varint,9,opt,name=validator_max_committee_size,json=validatorMaxCommitteeSize,proto3" json:"validator_max_committee_size,omitempty"`
	ValidatorBlockReward                    uint64 `protobuf:"varint,10,opt,name=validator_block_reward,json=validatorBlockReward,proto3" json:"validator_block_reward,omitempty"`
	ValidatorEarlyWithdrawalPenalty         uint64 `protobuf:"varint,11,opt,name=validator_early_withdrawal_penalty,json=validatorEarlyWithdrawalPenalty,proto3" json:"validator_early_withdrawal_penalty,omitempty"`
	ValidatorDelegateUnstakingBlocks        uint64 `protobuf:"varint,12,opt,name=validator_delegate_unstaking_blocks,json=validatorDelegateUnstakingBlocks,proto3" json:"validator_delegate_unstaking_blocks,omitempty"`
	ValidatorMinimumOrderSize               uint64 `protobuf:"varint,13,opt,name=validator_minimum_order_size,json=validatorMinimumOrderSize,proto3" json:"validator_minimum_order_size,omitempty"`
	ValidatorMinimumPercentForPaidCommittee uint64 `protobuf:"varint,14,opt,name=validator_minimum_percent_for_paid_committee,json=validatorMinimumPercentForPaidCommittee,proto3" json:"validator_minimum_percent_for_paid_committee,omitempty"`
	ValidatorMaxSlashPerCommittee           uint64 `protobuf:"varint,15,opt,name=validator_max_slash_per_committee,json=validatorMaxSlashPerCommittee,proto3" json:"validator_max_slash_per_committee,omitempty"`
}

func (x *ValidatorParams) Reset() {
	*x = ValidatorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorParams) ProtoMessage() {}

func (x *ValidatorParams) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorParams.ProtoReflect.Descriptor instead.
func (*ValidatorParams) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{3}
}

func (x *ValidatorParams) GetValidatorUnstakingBlocks() uint64 {
	if x != nil {
		return x.ValidatorUnstakingBlocks
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMaxPauseBlocks() uint64 {
	if x != nil {
		return x.ValidatorMaxPauseBlocks
	}
	return 0
}

func (x *ValidatorParams) GetValidatorDoubleSignSlashPercentage() uint64 {
	if x != nil {
		return x.ValidatorDoubleSignSlashPercentage
	}
	return 0
}

func (x *ValidatorParams) GetValidatorBadProposalSlashPercentage() uint64 {
	if x != nil {
		return x.ValidatorBadProposalSlashPercentage
	}
	return 0
}

func (x *ValidatorParams) GetValidatorNonSignSlashPercentage() uint64 {
	if x != nil {
		return x.ValidatorNonSignSlashPercentage
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMaxNonSign() uint64 {
	if x != nil {
		return x.ValidatorMaxNonSign
	}
	return 0
}

func (x *ValidatorParams) GetValidatorNonSignWindow() uint64 {
	if x != nil {
		return x.ValidatorNonSignWindow
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMaxCommittees() uint64 {
	if x != nil {
		return x.ValidatorMaxCommittees
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMaxCommitteeSize() uint64 {
	if x != nil {
		return x.ValidatorMaxCommitteeSize
	}
	return 0
}

func (x *ValidatorParams) GetValidatorBlockReward() uint64 {
	if x != nil {
		return x.ValidatorBlockReward
	}
	return 0
}

func (x *ValidatorParams) GetValidatorEarlyWithdrawalPenalty() uint64 {
	if x != nil {
		return x.ValidatorEarlyWithdrawalPenalty
	}
	return 0
}

func (x *ValidatorParams) GetValidatorDelegateUnstakingBlocks() uint64 {
	if x != nil {
		return x.ValidatorDelegateUnstakingBlocks
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMinimumOrderSize() uint64 {
	if x != nil {
		return x.ValidatorMinimumOrderSize
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMinimumPercentForPaidCommittee() uint64 {
	if x != nil {
		return x.ValidatorMinimumPercentForPaidCommittee
	}
	return 0
}

func (x *ValidatorParams) GetValidatorMaxSlashPerCommittee() uint64 {
	if x != nil {
		return x.ValidatorMaxSlashPerCommittee
	}
	return 0
}

type FeeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSendFee               uint64 `protobuf:"varint,1,opt,name=message_send_fee,json=messageSendFee,proto3" json:"message_send_fee,omitempty"`
	MessageStakeFee              uint64 `protobuf:"varint,2,opt,name=message_stake_fee,json=messageStakeFee,proto3" json:"message_stake_fee,omitempty"`
	MessageEditStakeFee          uint64 `protobuf:"varint,3,opt,name=message_edit_stake_fee,json=messageEditStakeFee,proto3" json:"message_edit_stake_fee,omitempty"`
	MessageUnstakeFee            uint64 `protobuf:"varint,4,opt,name=message_unstake_fee,json=messageUnstakeFee,proto3" json:"message_unstake_fee,omitempty"`
	MessagePauseFee              uint64 `protobuf:"varint,5,opt,name=message_pause_fee,json=messagePauseFee,proto3" json:"message_pause_fee,omitempty"`
	MessageUnpauseFee            uint64 `protobuf:"varint,6,opt,name=message_unpause_fee,json=messageUnpauseFee,proto3" json:"message_unpause_fee,omitempty"`
	MessageChangeParameterFee    uint64 `protobuf:"varint,7,opt,name=message_change_parameter_fee,json=messageChangeParameterFee,proto3" json:"message_change_parameter_fee,omitempty"`
	MessageDaoTransferFee        uint64 `protobuf:"varint,8,opt,name=message_dao_transfer_fee,json=messageDaoTransferFee,proto3" json:"message_dao_transfer_fee,omitempty"`
	MessageCertificateResultsFee uint64 `protobuf:"varint,9,opt,name=message_certificate_results_fee,json=messageCertificateResultsFee,proto3" json:"message_certificate_results_fee,omitempty"`
	MessageSubsidyFee            uint64 `protobuf:"varint,10,opt,name=message_subsidy_fee,json=messageSubsidyFee,proto3" json:"message_subsidy_fee,omitempty"`
	MessageCreateOrderFee        uint64 `protobuf:"varint,11,opt,name=message_create_order_fee,json=messageCreateOrderFee,proto3" json:"message_create_order_fee,omitempty"`
	MessageEditOrderFee          uint64 `protobuf:"varint,12,opt,name=message_edit_order_fee,json=messageEditOrderFee,proto3" json:"message_edit_order_fee,omitempty"`
	MessageDeleteOrderFee        uint64 `protobuf:"varint,13,opt,name=message_delete_order_fee,json=messageDeleteOrderFee,proto3" json:"message_delete_order_fee,omitempty"`
}

func (x *FeeParams) Reset() {
	*x = FeeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeParams) ProtoMessage() {}

func (x *FeeParams) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeParams.ProtoReflect.Descriptor instead.
func (*FeeParams) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{4}
}

func (x *FeeParams) GetMessageSendFee() uint64 {
	if x != nil {
		return x.MessageSendFee
	}
	return 0
}

func (x *FeeParams) GetMessageStakeFee() uint64 {
	if x != nil {
		return x.MessageStakeFee
	}
	return 0
}

func (x *FeeParams) GetMessageEditStakeFee() uint64 {
	if x != nil {
		return x.MessageEditStakeFee
	}
	return 0
}

func (x *FeeParams) GetMessageUnstakeFee() uint64 {
	if x != nil {
		return x.MessageUnstakeFee
	}
	return 0
}

func (x *FeeParams) GetMessagePauseFee() uint64 {
	if x != nil {
		return x.MessagePauseFee
	}
	return 0
}

func (x *FeeParams) GetMessageUnpauseFee() uint64 {
	if x != nil {
		return x.MessageUnpauseFee
	}
	return 0
}

func (x *FeeParams) GetMessageChangeParameterFee() uint64 {
	if x != nil {
		return x.MessageChangeParameterFee
	}
	return 0
}

func (x *FeeParams) GetMessageDaoTransferFee() uint64 {
	if x != nil {
		return x.MessageDaoTransferFee
	}
	return 0
}

func (x *FeeParams) GetMessageCertificateResultsFee() uint64 {
	if x != nil {
		return x.MessageCertificateResultsFee
	}
	return 0
}

func (x *FeeParams) GetMessageSubsidyFee() uint64 {
	if x != nil {
		return x.MessageSubsidyFee
	}
	return 0
}

func (x *FeeParams) GetMessageCreateOrderFee() uint64 {
	if x != nil {
		return x.MessageCreateOrderFee
	}
	return 0
}

func (x *FeeParams) GetMessageEditOrderFee() uint64 {
	if x != nil {
		return x.MessageEditOrderFee
	}
	return 0
}

func (x *FeeParams) GetMessageDeleteOrderFee() uint64 {
	if x != nil {
		return x.MessageDeleteOrderFee
	}
	return 0
}

type GovernanceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaoRewardPercentage uint64 `protobuf:"varint,1,opt,name=dao_reward_percentage,json=daoRewardPercentage,proto3" json:"dao_reward_percentage,omitempty"`
}

func (x *GovernanceParams) Reset() {
	*x = GovernanceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gov_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GovernanceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GovernanceParams) ProtoMessage() {}

func (x *GovernanceParams) ProtoReflect() protoreflect.Message {
	mi := &file_gov_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GovernanceParams.ProtoReflect.Descriptor instead.
func (*GovernanceParams) Descriptor() ([]byte, []int) {
	return file_gov_proto_rawDescGZIP(), []int{5}
}

func (x *GovernanceParams) GetDaoRewardPercentage() uint64 {
	if x != nil {
		return x.DaoRewardPercentage
	}
	return 0
}

var File_gov_proto protoreflect.FileDescriptor

var file_gov_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x6f, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x34, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x09,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x03, 0x46, 0x65, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46,
	0x65, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x03, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x47, 0x6f, 0x76, 0x65,
	0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xab, 0x08, 0x0a, 0x0f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3c, 0x0a, 0x1a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x18, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x50, 0x61, 0x75, 0x73,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x52, 0x0a, 0x26, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x22, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x27, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x23, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x4c, 0x0a, 0x23, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e,
	0x53, 0x6c, 0x61, 0x73, 0x68, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x33, 0x0a, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x4e, 0x6f, 0x6e,
	0x53, 0x69, 0x67, 0x6e, 0x12, 0x39, 0x0a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x38, 0x0a, 0x18, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x4b, 0x0a, 0x22, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x61,
	0x72, 0x6c, 0x79, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x70,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x4d, 0x0a,
	0x23, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3f, 0x0a, 0x1c,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x69, 0x6e,
	0x69, 0x6d, 0x75, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x5d, 0x0a,
	0x2c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70,
	0x61, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x27, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x50,
	0x61, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12, 0x48, 0x0a, 0x21,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x4d, 0x61, 0x78, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x22, 0xba, 0x05, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x46, 0x65, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x46, 0x65, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x46, 0x65, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x6e, 0x73, 0x74, 0x61,
	0x6b, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x46, 0x65, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x75, 0x73, 0x65,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x75, 0x73, 0x65, 0x46, 0x65, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x66,
	0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x55, 0x6e, 0x70, 0x61, 0x75, 0x73, 0x65, 0x46, 0x65, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a, 0x18,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x46, 0x65, 0x65, 0x12, 0x45, 0x0a, 0x1f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x46, 0x65, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x69, 0x64, 0x79, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x75, 0x62, 0x73, 0x69, 0x64, 0x79, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a, 0x18,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x65, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x64,
	0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x61, 0x6f, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x64, 0x61, 0x6f, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x2a, 0x49, 0x0a, 0x15, 0x47,
	0x6f, 0x76, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f, 0x41,
	0x4c, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x5f,
	0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x6e, 0x63, 0x68, 0x75, 0x63, 0x6f, 0x2f, 0x67, 0x69,
	0x6e, 0x63, 0x68, 0x75, 0x2f, 0x66, 0x73, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gov_proto_rawDescOnce sync.Once
	file_gov_proto_rawDescData = file_gov_proto_rawDesc
)

func file_gov_proto_rawDescGZIP() []byte {
	file_gov_proto_rawDescOnce.Do(func() {
		file_gov_proto_rawDescData = protoimpl.X.CompressGZIP(file_gov_proto_rawDescData)
	})
	return file_gov_proto_rawDescData
}

var file_gov_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gov_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gov_proto_goTypes = []interface{}{
	(GovProposalVoteConfig)(0), // 0: types.GovProposalVoteConfig
	(*Params)(nil),             // 1: types.Params
	(*ProtocolVersion)(nil),    // 2: types.ProtocolVersion
	(*ConsensusParams)(nil),    // 3: types.ConsensusParams
	(*ValidatorParams)(nil),    // 4: types.ValidatorParams
	(*FeeParams)(nil),          // 5: types.FeeParams
	(*GovernanceParams)(nil),   // 6: types.GovernanceParams
}
var file_gov_proto_depIdxs = []int32{
	3, // 0: types.Params.Consensus:type_name -> types.ConsensusParams
	4, // 1: types.Params.Validator:type_name -> types.ValidatorParams
	5, // 2: types.Params.Fee:type_name -> types.FeeParams
	6, // 3: types.Params.Governance:type_name -> types.GovernanceParams
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gov_proto_init() }
func file_gov_proto_init() {
	if File_gov_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gov_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_gov_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersion); i {
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
		file_gov_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusParams); i {
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
		file_gov_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorParams); i {
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
		file_gov_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeParams); i {
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
		file_gov_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GovernanceParams); i {
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
			RawDescriptor: file_gov_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gov_proto_goTypes,
		DependencyIndexes: file_gov_proto_depIdxs,
		EnumInfos:         file_gov_proto_enumTypes,
		MessageInfos:      file_gov_proto_msgTypes,
	}.Build()
	File_gov_proto = out.File
	file_gov_proto_rawDesc = nil
	file_gov_proto_goTypes = nil
	file_gov_proto_depIdxs = nil
}
