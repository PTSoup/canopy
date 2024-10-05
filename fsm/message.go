package fsm

import (
	"github.com/ginchuco/ginchu/fsm/types"
	"github.com/ginchuco/ginchu/lib"
	"github.com/ginchuco/ginchu/lib/crypto"
)

// HandleMessage() routes the MessageI to the correct `handler` based on its `type`
func (s *StateMachine) HandleMessage(msg lib.MessageI) lib.ErrorI {
	switch x := msg.(type) {
	case *types.MessageSend:
		return s.HandleMessageSend(x)
	case *types.MessageStake:
		return s.HandleMessageStake(x)
	case *types.MessageEditStake:
		return s.HandleMessageEditStake(x)
	case *types.MessageUnstake:
		return s.HandleMessageUnstake(x)
	case *types.MessagePause:
		return s.HandleMessagePause(x)
	case *types.MessageUnpause:
		return s.HandleMessageUnpause(x)
	case *types.MessageChangeParameter:
		return s.HandleMessageChangeParameter(x)
	case *types.MessageDAOTransfer:
		return s.HandleMessageDAOTransfer(x)
	case *types.MessageCertificateResults:
		return s.HandleMessageCertificateResults(x)
	case *types.MessageSubsidy:
		return s.HandleMessageSubsidy(x)
	case *types.MessageCreateOrder:
		return s.HandleMessageCreateOrder(x)
	case *types.MessageEditOrder:
		return s.HandleMessageEditOrder(x)
	case *types.MessageDeleteOrder:
		return s.HandleMessageDeleteOrder(x)
	default:
		return types.ErrUnknownMessage(x)
	}
}

// GetFeeForMessage() returns the associated cost for processing a specific type of message
func (s *StateMachine) GetFeeForMessage(msg lib.MessageI) (fee uint64, err lib.ErrorI) {
	feeParams, err := s.GetParamsFee()
	if err != nil {
		return 0, err
	}
	switch x := msg.(type) {
	case *types.MessageSend:
		return feeParams.MessageSendFee, nil
	case *types.MessageStake:
		return feeParams.MessageStakeFee, nil
	case *types.MessageEditStake:
		return feeParams.MessageEditStakeFee, nil
	case *types.MessageUnstake:
		return feeParams.MessageUnstakeFee, nil
	case *types.MessageUnpause:
		return feeParams.MessageUnpauseFee, nil
	case *types.MessageChangeParameter:
		return feeParams.MessageChangeParameterFee, nil
	case *types.MessageDAOTransfer:
		return feeParams.MessageDaoTransferFee, nil
	case *types.MessageCertificateResults:
		return feeParams.MessageCertificateResultsFee, nil
	case *types.MessageSubsidy:
		return feeParams.MessageSubsidyFee, nil
	case *types.MessageCreateOrder:
		return feeParams.MessageCreateOrderFee, nil
	case *types.MessageEditOrder:
		return feeParams.MessageEditOrderFee, nil
	case *types.MessageDeleteOrder:
		return feeParams.MessageDeleteOrderFee, nil
	default:
		return 0, types.ErrUnknownMessage(x)
	}
}

// GetAuthorizedSignersFor() returns the addresses that are authorized to sign for this message
func (s *StateMachine) GetAuthorizedSignersFor(msg lib.MessageI) (signers [][]byte, err lib.ErrorI) {
	var validator *types.Validator
	switch x := msg.(type) {
	case *types.MessageSend:
		return [][]byte{x.FromAddress}, nil
	case *types.MessageChangeParameter:
		return [][]byte{x.Signer}, nil
	case *types.MessageDAOTransfer:
		return [][]byte{x.Address}, nil
	case *types.MessageStake:
		pubKey, e := crypto.NewPublicKeyFromBytes(x.PublicKey)
		if e != nil {
			return nil, types.ErrInvalidPublicKey(e)
		}
		validator, err = s.GetValidator(pubKey.Address())
		if err != nil {
			return nil, err
		}
	case *types.MessageEditStake:
		validator, err = s.GetValidator(crypto.NewAddressFromBytes(x.Address))
	case *types.MessageUnstake:
		validator, err = s.GetValidator(crypto.NewAddressFromBytes(x.Address))
	case *types.MessageUnpause:
		validator, err = s.GetValidator(crypto.NewAddressFromBytes(x.Address))
	case *types.MessageSubsidy:
		return [][]byte{x.Address}, nil
	case *types.MessageCreateOrder:
		return [][]byte{x.SellersSellAddress}, nil
	case *types.MessageEditOrder:
		order, e := s.GetOrder(x.OrderId, x.CommitteeId)
		if e != nil {
			return nil, e
		}
		return [][]byte{order.SellersSellAddress}, nil
	case *types.MessageDeleteOrder:
		order, e := s.GetOrder(x.OrderId, x.CommitteeId)
		if e != nil {
			return nil, e
		}
		return [][]byte{order.SellersSellAddress}, nil
	case *types.MessageCertificateResults:
		// any member may submit it - as the reward recipients are locked in the QC.ProposalHash
		committee, e := s.GetCommittee(x.Qc.Header.CommitteeId)
		if e != nil {
			return nil, e
		}
		for _, member := range committee.ValidatorSet.ValidatorSet {
			pk, _ := crypto.NewPublicKeyFromBytes(member.PublicKey)
			signers = append(signers, pk.Address().Bytes())
		}
		return
	default:
		return nil, types.ErrUnknownMessage(x)
	}
	if err != nil {
		return nil, err
	}
	if validator == nil {
		return nil, types.ErrValidatorNotExists()
	}
	return [][]byte{validator.Address, validator.Output}, nil
}

// HandleMessageSend() is the proper handler for a `Send` message
func (s *StateMachine) HandleMessageSend(msg *types.MessageSend) lib.ErrorI {
	// subtract from sender
	if err := s.AccountSub(crypto.NewAddressFromBytes(msg.FromAddress), msg.Amount); err != nil {
		return err
	}
	// add to recipient
	return s.AccountAdd(crypto.NewAddressFromBytes(msg.ToAddress), msg.Amount)
}

// HandleMessageStake() is the proper handler for a `Stake` message (Validator does not yet exist in the state)
func (s *StateMachine) HandleMessageStake(msg *types.MessageStake) lib.ErrorI {
	publicKey, e := crypto.NewPublicKeyFromBytes(msg.PublicKey)
	if e != nil {
		return types.ErrInvalidPublicKey(e)
	}
	address := publicKey.Address()
	// check if validator exists
	exists, err := s.GetValidatorExists(address)
	if err != nil {
		return err
	}
	// fail if validator already exists
	if exists {
		return types.ErrValidatorExists()
	}
	// check if below minimum stake
	params, err := s.GetParamsVal()
	if err != nil {
		return err
	}
	if msg.Amount < params.ValidatorMinStake {
		return types.ErrBelowMinimumStake()
	}
	// subtract from sender
	if err = s.AccountSub(address, msg.Amount); err != nil {
		return err
	}
	// track total staked tokens
	if err = s.AddToStakedSupply(msg.Amount); err != nil {
		return err
	}
	if msg.Delegate {
		// set delegated validator in each committee
		if err = s.SetDelegations(address, msg.Amount, msg.Committees); err != nil {
			return err
		}
	} else {
		// set validator in each committee
		if err = s.SetCommittees(address, msg.Amount, msg.Committees); err != nil {
			return err
		}
	}
	// set validator
	return s.SetValidator(&types.Validator{
		Address:      address.Bytes(),
		PublicKey:    publicKey.Bytes(),
		NetAddress:   msg.NetAddress,
		StakedAmount: msg.Amount,
		Committees:   msg.Committees,
		Output:       msg.OutputAddress,
		Delegate:     msg.Delegate,
		Compound:     msg.Compound,
	})
}

// HandleMessageEditStake() is the proper handler for a `Edit-Stake` message (Validator already exists in the state)
func (s *StateMachine) HandleMessageEditStake(msg *types.MessageEditStake) lib.ErrorI {
	address := crypto.NewAddressFromBytes(msg.Address)
	// get the validator
	val, err := s.GetValidator(address)
	if err != nil {
		return err
	}
	// check unstaking
	if val.UnstakingHeight != 0 {
		return types.ErrValidatorUnstaking()
	}
	var amountToAdd uint64
	switch {
	case msg.Amount < val.StakedAmount: // amount less than stake
		return types.ErrInvalidAmount()
	case msg.Amount == val.StakedAmount: // amount equals stake
	case msg.Amount > val.StakedAmount: // amount greater than stake
		amountToAdd = msg.Amount - val.StakedAmount
	}
	// subtract from sender
	if err = s.AccountSub(address, amountToAdd); err != nil {
		return err
	}
	if err = s.AddToStakedSupply(amountToAdd); err != nil {
		return err
	}
	// update validator stake amount
	newStakedAmount := val.StakedAmount + amountToAdd
	// use validator.delegate value -> not allowed to change delegation status
	if val.Delegate {
		if err = s.UpdateDelegations(address, val, msg.Amount, msg.Committees); err != nil {
			return err
		}
	} else {
		if err = s.UpdateCommittees(address, val, msg.Amount, msg.Committees); err != nil {
			return err
		}
	}
	// set validator
	return s.SetValidator(&types.Validator{
		Address:         val.Address,
		PublicKey:       val.PublicKey,
		NetAddress:      msg.NetAddress,
		StakedAmount:    newStakedAmount,
		Committees:      msg.Committees,
		MaxPausedHeight: val.MaxPausedHeight,
		UnstakingHeight: val.UnstakingHeight,
		Output:          msg.OutputAddress,
		Delegate:        val.Delegate,
		Compound:        msg.Compound,
	})
}

// HandleMessageUnstake() is the proper handler for an `Unstake` message
func (s *StateMachine) HandleMessageUnstake(msg *types.MessageUnstake) lib.ErrorI {
	address := crypto.NewAddressFromBytes(msg.Address)
	// get validator
	validator, err := s.GetValidator(address)
	if err != nil {
		return err
	}
	// check if already unstaking
	if validator.UnstakingHeight != 0 {
		return types.ErrValidatorUnstaking()
	}
	// get params for unstaking blocks
	p, err := s.GetParamsVal()
	if err != nil {
		return err
	}
	// get unstaking blocks parameter
	var unstakingBlocks uint64
	if !validator.Delegate {
		unstakingBlocks = p.ValidatorUnstakingBlocks
	} else {
		unstakingBlocks = p.ValidatorDelegateUnstakingBlocks
	}
	unstakingHeight := s.Height() + unstakingBlocks
	// set validator unstaking
	return s.SetValidatorUnstaking(address, validator, unstakingHeight)
}

// HandleMessagePause() is the proper handler for an `Pause` message
func (s *StateMachine) HandleMessagePause(msg *types.MessagePause) lib.ErrorI {
	address := crypto.NewAddressFromBytes(msg.Address)
	// get validator
	validator, err := s.GetValidator(address)
	if err != nil {
		return err
	}
	// ensure not already paused
	if validator.MaxPausedHeight != 0 {
		return types.ErrValidatorPaused()
	}
	if validator.UnstakingHeight != 0 {
		return types.ErrValidatorUnstaking()
	}
	if validator.Delegate {
		return types.ErrInvalidDelegationStatus()
	}
	// get max pause parameter
	params, err := s.GetParamsVal()
	if err != nil {
		return err
	}
	maxPausedHeight := s.Height() + params.ValidatorMaxPauseBlocks
	// set validator paused
	return s.SetValidatorPaused(address, validator, maxPausedHeight)
}

// HandleMessageUnpause() is the proper handler for an `Unpause` message
func (s *StateMachine) HandleMessageUnpause(msg *types.MessageUnpause) lib.ErrorI {
	address := crypto.NewAddressFromBytes(msg.Address)
	// get validator
	validator, err := s.GetValidator(address)
	if err != nil {
		return err
	}
	// ensure already paused
	if validator.MaxPausedHeight == 0 {
		return types.ErrValidatorNotPaused()
	}
	if validator.UnstakingHeight != 0 {
		return types.ErrValidatorUnstaking()
	}
	// set validator unpaused
	return s.SetValidatorUnpaused(address, validator)
}

// HandleMessageChangeParameter() is the proper handler for an `Change-Parameter` message
func (s *StateMachine) HandleMessageChangeParameter(msg *types.MessageChangeParameter) lib.ErrorI {
	// requires explicit approval from +2/3 maj of the validator set
	if err := s.ApproveProposal(msg); err != nil {
		return types.ErrRejectProposal()
	}
	// update the parameter
	protoMsg, err := lib.FromAny(msg.ParameterValue)
	if err != nil {
		return err
	}
	return s.UpdateParam(msg.ParameterSpace, msg.ParameterKey, protoMsg)
}

// HandleMessageDAOTransfer() is the proper handler for a `DAO-Transfer` message
func (s *StateMachine) HandleMessageDAOTransfer(msg *types.MessageDAOTransfer) lib.ErrorI {
	// requires explicit approval from +2/3 maj of the validator set
	if err := s.ApproveProposal(msg); err != nil {
		return types.ErrRejectProposal()
	}
	// remove from DAO fund
	if err := s.PoolSub(lib.DAOPoolID, msg.Amount); err != nil {
		return err
	}
	// add to account
	return s.AccountAdd(crypto.NewAddressFromBytes(msg.Address), msg.Amount)
}

// HandleMessageCertificateResults() is the proper handler for a `CertificateResults` message
func (s *StateMachine) HandleMessageCertificateResults(msg *types.MessageCertificateResults) lib.ErrorI {
	results, committeeId, store := msg.Qc.Results, msg.Qc.Header.CommitteeId, s.store.(lib.StoreI)
	// get the proper reward Pool
	pool, err := s.GetPool(committeeId)
	if err != nil {
		return err
	}
	// ensure subsidized
	if pool.Amount == 0 {
		return types.ErrNonSubsidizedCommittee()
	}
	// validate the height of the CertificateResults Transaction
	height := msg.Qc.Header.CommitteeHeight
	if height != s.Height() && height != s.Height()-1 {
		return lib.ErrInvalidQCCommitteeHeight()
	}
	// get committee for the QC
	committee, err := s.LoadCommittee(committeeId, height)
	if err != nil {
		return err
	}
	// ensure it's a valid QC
	// NOTE: max block size is 0 here because there should not be a block attached to this QC
	isPartialQC, err := msg.Qc.Check(committee, 0, &lib.View{NetworkId: uint64(s.NetworkID), CommitteeId: committeeId}, false)
	if err != nil {
		return err
	}
	if isPartialQC {
		return lib.ErrNoMaj23()
	}
	// handle the token swaps from the
	if err = s.HandleCommitteeBuyOrders(results.Orders, committeeId); err != nil {
		return err
	}
	// handle checkpoint-as-a-service functionality
	if results.Checkpoint != nil {
		// ensure checkpoint isn't older than the most recent
		mostRecentCheckpoint, e := store.GetMostRecentCheckpoint(committeeId)
		if e != nil {
			return e
		}
		if results.Checkpoint.Height <= mostRecentCheckpoint.Height {
			return types.ErrInvalidCheckpoint()
		}
		// index the checkpoint
		if err = store.IndexCheckpoint(committeeId, results.Checkpoint); err != nil {
			return err
		}
	}
	// reduce payment percentages by non-signer-percent
	validatorParams, err := s.GetParamsVal()
	if err != nil {
		return err
	}
	nonSignerPercent, err := s.HandleByzantine(msg.Qc, committee.ValidatorSet, validatorParams)
	if err != nil {
		return err
	}
	for i, p := range results.RewardRecipients.PaymentPercents {
		results.RewardRecipients.PaymentPercents[i].Percent = lib.Uint64ReducePercentage(p.Percent, float64(nonSignerPercent))
	}
	// upsert a fund
	return s.UpsertCommitteeData(&types.CommitteeData{
		CommitteeId:     committeeId,
		CommitteeHeight: msg.Qc.Header.CommitteeHeight,
		ChainHeight:     msg.Qc.Header.Height,
		PaymentPercents: results.RewardRecipients.PaymentPercents,
	})
}

// HandleMessageSubsidy() is the proper handler for a `Subsidy` message
func (s *StateMachine) HandleMessageSubsidy(msg *types.MessageSubsidy) lib.ErrorI {
	// subtract from sender
	if err := s.AccountSub(crypto.NewAddressFromBytes(msg.Address), msg.Amount); err != nil {
		return err
	}
	// add to recipient committee
	return s.PoolAdd(msg.CommitteeId, msg.Amount)
}

// HandleMessageCreateOrder() is the proper handler for a `CreateOrder` message
func (s *StateMachine) HandleMessageCreateOrder(msg *types.MessageCreateOrder) (err lib.ErrorI) {
	valParams, err := s.GetParamsVal()
	if err != nil {
		return
	}
	// ensure order isn't below the minimum size
	if msg.RequestedAmount < valParams.ValidatorMinimumOrderSize {
		return types.ErrMinimumOrderSize()
	}
	// subtract from account balance
	address := crypto.NewAddress(msg.SellersSellAddress)
	if err = s.AccountSub(address, msg.AmountForSale); err != nil {
		return
	}
	// add to committee escrow pool
	if err = s.PoolAdd(msg.CommitteeId+uint64(types.EscrowPoolAddend), msg.AmountForSale); err != nil {
		return
	}
	// count it in the 'supply' structure
	if err = s.AddToEscrowedSupply(msg.CommitteeId, msg.AmountForSale); err != nil {
		return
	}
	// save the order in state
	_, err = s.CreateOrder(&types.SellOrder{
		Committee:             msg.CommitteeId,
		AmountForSale:         msg.AmountForSale,
		RequestedAmount:       msg.RequestedAmount,
		SellerReceiveAddress:  msg.SellerReceiveAddress,
		SellersSellAddress:    msg.SellersSellAddress,
		OrderExpirationHeight: s.Height() + valParams.ValidatorOrderExpirationBlocks,
	}, msg.CommitteeId)
	return
}

// HandleMessageEditOrder() is the proper handler for a `EditOrder` message
func (s *StateMachine) HandleMessageEditOrder(msg *types.MessageEditOrder) (err lib.ErrorI) {
	order, err := s.GetOrder(msg.OrderId, msg.CommitteeId)
	if err != nil {
		return
	}
	if order.BuyerReceiveAddress != nil {
		return types.ErrOrderAlreadyAccepted()
	}
	valParams, err := s.GetParamsVal()
	if err != nil {
		return
	}
	// ensure order isn't below the minimum size
	if msg.RequestedAmount < valParams.ValidatorMinimumOrderSize {
		return types.ErrMinimumOrderSize()
	}
	difference, address := int(msg.AmountForSale-order.AmountForSale), crypto.NewAddress(order.SellersSellAddress)
	if difference > 0 {
		amountDifference := uint64(difference)
		if err = s.AccountSub(address, amountDifference); err != nil {
			return
		}
		// add to committee escrow pool
		if err = s.PoolAdd(msg.CommitteeId+uint64(types.EscrowPoolAddend), amountDifference); err != nil {
			return
		}
		if err = s.AddToEscrowedSupply(order.Committee, amountDifference); err != nil {
			return
		}
	} else if difference < 0 {
		amountDifference := uint64(difference * -1)
		if err = s.SubFromEscrowedSupply(order.Committee, amountDifference); err != nil {
			return
		}
		// subtract from the committee escrow pool
		if err = s.PoolSub(msg.CommitteeId+uint64(types.EscrowPoolAddend), amountDifference); err != nil {
			return
		}
		if err = s.AccountAdd(address, amountDifference); err != nil {
			return
		}
	}
	err = s.EditOrder(&types.SellOrder{
		Id:                    order.Id,
		Committee:             msg.CommitteeId,
		AmountForSale:         msg.AmountForSale,
		RequestedAmount:       msg.RequestedAmount,
		SellerReceiveAddress:  msg.SellerReceiveAddress,
		OrderExpirationHeight: s.Height() + valParams.ValidatorOrderExpirationBlocks,
		SellersSellAddress:    order.SellersSellAddress,
	}, msg.CommitteeId)
	return
}

// HandleMessageDeleteOrder() is the proper handler for a `DeleteOrder` message
func (s *StateMachine) HandleMessageDeleteOrder(msg *types.MessageDeleteOrder) (err lib.ErrorI) {
	order, err := s.GetOrder(msg.OrderId, msg.CommitteeId)
	if err != nil {
		return
	}
	if order.BuyerReceiveAddress != nil {
		return types.ErrOrderAlreadyAccepted()
	}
	if err = s.SubFromEscrowedSupply(order.Committee, order.AmountForSale); err != nil {
		return
	}
	// subtract from the committee escrow pool
	if err = s.PoolSub(msg.CommitteeId+uint64(types.EscrowPoolAddend), order.AmountForSale); err != nil {
		return
	}
	if err = s.AccountAdd(crypto.NewAddress(order.SellersSellAddress), order.AmountForSale); err != nil {
		return
	}
	err = s.DeleteOrder(msg.OrderId, msg.CommitteeId)
	return
}
