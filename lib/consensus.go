package lib

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/canopy-network/canopy/lib/crypto"
	"github.com/drand/kyber"
)

// ValidatorSet represents a collection of validators responsible for consensus
// It facilitates the creation and validation of +2/3 Majority agreements using multi-signatures
type ValidatorSet struct {
	ValidatorSet  *ConsensusValidators   // a list of validators participating in the consensus process
	MultiKey      crypto.MultiPublicKeyI // a composite public key derived from the individual public keys of all validators, used for verifying multi-signatures
	TotalPower    uint64                 // the aggregate voting power of all validators in the set, reflecting their influence on the consensus
	MinimumMaj23  uint64                 // the minimum voting power threshold required to achieve a two-thirds majority (2f+1), essential for consensus decisions
	NumValidators uint64                 // the total number of validators in the set, indicating the size of the validator pool
}

// NewValidatorSet() initializes a ValidatorSet from a given set of consensus validators
func NewValidatorSet(validators *ConsensusValidators) (ValidatorSet, ErrorI) {
	// handle empty set
	if validators == nil {
		// exit with error
		return ValidatorSet{}, ErrNoValidators()
	}
	// define tracking variables
	totalPower, count, pointList := uint64(0), uint64(0), make([]kyber.Point, 0)
	// iterate through the ValidatorSet to get the count, total power, and convert
	// the public keys to 'points' on an elliptic curve for the BLS multikey
	for _, v := range validators.ValidatorSet {
		// convert the public key into a BLS point
		point, err := crypto.BytesToBLS12381Point(v.PublicKey)
		// check for an error during the conversion
		if err != nil {
			// exit with error
			return ValidatorSet{}, ErrPubKeyFromBytes(err)
		}
		// add the point to the list
		pointList = append(pointList, point)
		// update total voting power
		totalPower += v.VotingPower
		// increment the count
		count++
	}
	// if the total voting power is 0
	if totalPower == 0 {
		// exit with error
		return ValidatorSet{}, ErrNoValidators()
	}
	// calculate the minimum power for a two-thirds majority (2f+1)
	minPowerFor23Maj := (2*totalPower)/3 + 1
	// create a composite multi-public key out of the public keys (in curve point format)
	multiPublicKey, err := crypto.NewMultiBLSFromPoints(pointList, nil)
	// if an error occurred during the conversion
	if err != nil {
		// exit with error
		return ValidatorSet{}, ErrNewMultiPubKey(err)
	}
	// return the validator set
	return ValidatorSet{
		ValidatorSet:  validators,
		MultiKey:      multiPublicKey,
		TotalPower:    totalPower,
		MinimumMaj23:  minPowerFor23Maj,
		NumValidators: count,
	}, nil
}

// GetValidator() retrieves a validator from the ValidatorSet using the public key
func (vs *ValidatorSet) GetValidator(publicKey []byte) (val *ConsensusValidator, err ErrorI) {
	// retrieve the validator using a public key
	val, _, err = vs.GetValidatorAndIdx(publicKey)
	// exit
	return
}

// GetValidatorAndIdx() retrieves a validator and its index in the ValidatorSet using the public key
func (vs *ValidatorSet) GetValidatorAndIdx(targetPublicKey []byte) (val *ConsensusValidator, idx int, err ErrorI) {
	// if the validator set is empty
	if vs == nil || vs.ValidatorSet == nil {
		// exit with error
		return nil, 0, ErrInvalidValidatorIndex()
	}
	// for each validator in the set
	for i, v := range vs.ValidatorSet.ValidatorSet {
		// if the public key is equal to the target
		if bytes.Equal(v.PublicKey, targetPublicKey) {
			// exit with the validator and the index
			return v, i, nil
		}
	}
	// exit with 'not found' error
	return nil, 0, ErrValidatorNotInSet(targetPublicKey)
}

// RootChainInfo maintains root-chain data needed for consensus
type RootChainInfo struct {
	RootChainId        uint64             `json:"rootChainId"`      // the id of the root chain
	Height             uint64             `json:"height"`           // the height of the root chain
	ValidatorSet       ValidatorSet       `json:"validatorSet"`     // the validator set for this 'nested chain' as determined by the 'root chain'
	LastValidatorSet   ValidatorSet       `json:"lastValidatorSet"` // the previous validator set for this 'nested chain' as previously determined by the 'root chain'
	LotteryWinner      *LotteryWinner     `json:"lotteryWinner"`    // the 'delegate' lottery winner who may be a 'reward recipient' in the certificate result
	Orders             *OrderBook         `json:"orders"`           // the swap order book from the 'root chain' for the 'nested chain'
	GetRemoteCallbacks GetRemoteCallbacks `json:"remoteCallbacks"`  // the rpc callbacks for the remote chain
	Log                LoggerI            `json:"log"`              // the stout logger
}

// GetRemoteCallbacks is a callback to get the 'rpc client' for the root chain id
type GetRemoteCallbacks func(rootChainId uint64) (*RemoteCallbacks, ErrorI)

// RemoteCallbacks are fallback rpc callbacks to the root-chain
type RemoteCallbacks struct {
	Height              func() (*uint64, ErrorI)                                   // retrieve the height of the root chain
	RootChainInfo       func(height, chainId uint64) (*RootChainInfo, ErrorI)      // retrieve the root chain info in 1 call
	ValidatorSet        func(height, id uint64) (ValidatorSet, ErrorI)             // get the validator set for a chain id using the RPC API
	Lottery             func(height, id uint64) (p *LotteryWinner, err ErrorI)     // get the delegate 'lottery winner' for a chain id
	Orders              func(height, chainId uint64) (p *OrderBooks, err ErrorI)   // get the order book for a specific 'chain id'
	Order               func(height, orderId, chainId uint64) (*SellOrder, ErrorI) // get a specific order from the order book
	IsValidDoubleSigner func(height uint64, address string) (p *bool, err ErrorI)  // check if a double signer is valid for an address for a specific 'double sign height'
	Checkpoint          func(height, id uint64) (blockHash HexBytes, i ErrorI)     // get a checkpoint at a height and chain id combination
	Transaction         func(tx TransactionI) (hash *string, err ErrorI)           // submit a transaction to the 'root chain'
}

// GetHeight() returns the height from the root-chain
func (b *RootChainInfo) GetHeight() uint64 { return b.Height }

// GetValidatorSet() returns the validator set from the root-chain
func (b *RootChainInfo) GetValidatorSet(rootChainId, id, rootHeight uint64) (ValidatorSet, ErrorI) {
	// if the root chain id is the same as the info
	if rootChainId == b.RootChainId {
		// if rootHeight is the same as the RootChainInfo height
		if rootHeight == b.Height {
			// exit with a copy the validator set
			return NewValidatorSet(b.ValidatorSet.ValidatorSet)
		}
		// if rootHeight is 1 before the RootChainInfo height
		if rootHeight == b.Height-1 {
			// exit with a copy of the previous validator set
			return NewValidatorSet(b.LastValidatorSet.ValidatorSet)
		}
	}
	// warn of the remote RPC call to the root chain API
	b.Log.Warnf("Executing remote GetValidatorSet call with requested height=%d for rootChainId=%d", rootHeight, rootChainId)
	// get a rpc client for the root chain id
	rpcClient, err := b.GetRemoteCallbacks(rootChainId)
	// if the rpc client retrieval fails
	if err != nil {
		// exit with error
		return ValidatorSet{}, err
	}
	// execute the remote RPC call to the root chain API
	return rpcClient.ValidatorSet(rootHeight, id)
}

// GetOrders() returns the order book from the root-chain
func (b *RootChainInfo) GetOrders(rootChainId, rootHeight, id uint64) (*OrderBook, ErrorI) {
	// if the root chain id is the same as the info
	if rootChainId == b.RootChainId {
		// if rootHeight is the same as the RootChainInfo height
		if rootHeight == b.Height {
			// exit with the order books from memory
			return b.Orders, nil
		}
	}
	// warn of the remote RPC call to the root chain API
	b.Log.Warnf("Executing remote GetOrders call with requested height=%d for rootChainId=%d", rootHeight, rootChainId)
	// get a rpc client for the root chain id
	rpcClient, err := b.GetRemoteCallbacks(rootChainId)
	// if the rpc client retrieval fails
	if err != nil {
		// exit with error
		return nil, err
	}
	// execute the remote call
	books, err := rpcClient.Orders(rootHeight, id)
	// if an error occurred during the remote call
	if err != nil {
		// exit with error
		return nil, err
	}
	// ensure the order book isn't empty
	if books == nil || len(books.OrderBooks) == 0 {
		// exit with error
		return nil, ErrEmptyOrderBook()
	}
	// exit with the first (and only) order book in the list
	return books.OrderBooks[0], nil
}

// IsValidDoubleSigner() returns if an address is a valid double signer for a specific 'double sign height'
func (b *RootChainInfo) IsValidDoubleSigner(rootChainId, height uint64, address string) (*bool, ErrorI) {
	// warn of the remote RPC call to the API of the 'root chain'
	b.Log.Warnf("Executing remote IsValidDoubleSigner call for address=%s at double sign height=%d and rootChainId=%d", address, height, rootChainId)
	// get a rpc client for the root chain id
	rpcClient, err := b.GetRemoteCallbacks(rootChainId)
	// if the rpc client retrieval fails
	if err != nil {
		// exit with error
		return nil, err
	}
	// exit with the results of the remote RPC call to the API of the 'root chain'
	return rpcClient.IsValidDoubleSigner(height, address)
}

// GetCheckpoint() returns the checkpoint if any for a specific chain height
// TODO should be able to get these from the file or the root-chain upon independence
func (b *RootChainInfo) GetCheckpoint(rootChainId, height, chainId uint64) (blockHash HexBytes, err ErrorI) {
	// warn of the remote RPC call to the API of the 'root chain'
	b.Log.Warnf("Executing remote GetCheckpoint call for height=%d and chainId=%d and rootChainId=%d", height, chainId, rootChainId)
	// get a rpc client for the root chain id
	rpcClient, err := b.GetRemoteCallbacks(rootChainId)
	// if the rpc client retrieval fails
	if err != nil {
		// exit with error
		return nil, err
	}
	// exit with the results of the remote RPC call to the API of the 'root chain'
	return rpcClient.Checkpoint(height, chainId)
}

// GetLotteryWinner() returns the winner of the delegate lottery from the root-chain
func (b *RootChainInfo) GetLotteryWinner(rootChainId, height, id uint64) (*LotteryWinner, ErrorI) {
	// if the root chain id is the same as the info
	if rootChainId == b.RootChainId {
		// if rootHeight is the same as the RootChainInfo height
		if height == b.Height {
			// exit with the lottery winner
			return b.LotteryWinner, nil
		}
	}
	// warn of the remote RPC call to the API of the 'root chain'
	b.Log.Warnf("Executing remote Lottery call with requested height=%d and rootChainId=%d", height, rootChainId)
	// get a rpc client for the root chain id
	rpcClient, err := b.GetRemoteCallbacks(rootChainId)
	// if the rpc client retrieval fails
	if err != nil {
		// exit with error
		return nil, err
	}
	// exit with the results of the remote RPC call to the API of the 'root chain'
	return rpcClient.Lottery(height, id)
}

// rootChainInfoJSON is the encoding structure used for json for RootChainInfo
type rootChainInfoJSON struct {
	RootChainId   uint64               `json:"rootChainId"`
	Height        uint64               `json:"height"`
	Committee     *ConsensusValidators `json:"committee"`
	LastCommittee *ConsensusValidators `json:"last_committee"`
	LotteryWinner *LotteryWinner       `json:"lottery_winner"`
	Orders        *OrderBook           `json:"orders"`
}

// MarshalJSON() implements the json.Marshaller for RootChainInfo
func (b RootChainInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(rootChainInfoJSON{
		RootChainId:   b.RootChainId,
		Height:        b.Height,
		Committee:     b.ValidatorSet.ValidatorSet,
		LastCommittee: b.LastValidatorSet.ValidatorSet,
		LotteryWinner: b.LotteryWinner,
		Orders:        b.Orders,
	})
}

// UnmarshalJSON() implements the json.Unmarshaler for RootChainInfo
func (b *RootChainInfo) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new object reference to ensure a non nil result
	j := new(rootChainInfoJSON)
	// populate the object ref with the json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit with error
		return
	}
	// create a validator set from the committee
	validatorSet, err := NewValidatorSet(j.Committee)
	// if an error occurred during the conversion
	if err != nil {
		// exit with error
		return
	}
	// create a validator set from the previous committee - ignore error
	lastValidatorSet, _ := NewValidatorSet(j.LastCommittee)
	// set the underlying object
	*b = RootChainInfo{
		RootChainId:        j.RootChainId,
		Height:             j.Height,
		ValidatorSet:       validatorSet,
		LastValidatorSet:   lastValidatorSet,
		LotteryWinner:      j.LotteryWinner,
		Orders:             j.Orders,
		GetRemoteCallbacks: nil,
		Log:                nil,
	}
	// exit
	return
}

// LotteryWinner is a structure that holds the subject of a pseudorandom selection and their % cut of the reward
// This is used for delegation + sub-delegation + sub-validator earnings
type LotteryWinner struct {
	Winner HexBytes `json:"winner"` // the winner address in hex bytes
	Cut    uint64   `json:"cut"`    // the % of the reward that is allocated to the winner
}

// CheckBasic() validates the basic structure and length of the AggregateSignature
func (x *AggregateSignature) CheckBasic() ErrorI {
	// ensure the aggregate signature is not nil
	if x == nil {
		// exit with empty error
		return ErrEmptyAggregateSignature()
	}
	// ensure the signature bytes are the proper size
	if len(x.Signature) != crypto.BLS12381SignatureSize {
		// exit with signature size error
		return ErrInvalidAggrSignatureLength()
	}
	// ensure the bitmap is not empty
	if len(x.Bitmap) == 0 {
		// exit with empty bitmap error
		return ErrEmptySignerBitmap()
	}
	// exit
	return nil
}

// Check() validates a +2/3 majority of the signature using the payload bytes and the ValidatorSet
// NOTE: "partialQC" means the signature is valid but does not reach a +2/3 majority
func (x *AggregateSignature) Check(sb SignByte, vs ValidatorSet) (isPartialQC bool, err ErrorI) {
	// perform basic 'sanity checks' on the aggregate signature structure
	if err = x.CheckBasic(); err != nil {
		// exit with error
		return false, err
	}
	// create a copy of the mutli-public-key of the validator set
	key := vs.MultiKey.Copy()
	// indicate which validator indexes have purportedly signed the payload
	// and are included in the aggregated signature
	if er := key.SetBitmap(x.Bitmap); er != nil {
		// exit with error
		return false, ErrInvalidSignerBitmap(er)
	}
	// use the composite public key to verify the aggregate signature
	if !key.VerifyBytes(sb.SignBytes(), x.Signature) {
		// exit with error
		return false, ErrInvalidAggrSignature()
	}
	// get the total power and the min +2/3 majority from the bitmap and ValSet
	_, totalSignedPower, err := x.GetSigners(vs)
	// if an error occured when retrieving the signers
	if err != nil {
		// exit with error
		return false, err
	}
	// ensure the signers reach a +2/3 majority
	if totalSignedPower < vs.MinimumMaj23 {
		// exit with isPartialQC
		return true, nil
	}
	// exit
	return false, nil
}

// GetSigners() returns the public keys and corresponding combined voting power of those who signed
func (x *AggregateSignature) GetSigners(vs ValidatorSet) (signers [][]byte, signedPower uint64, err ErrorI) {
	// retrieve the signers for the aggregate signatures
	signers, signedPower, err = x.getSigners(vs, false)
	// exit
	return
}

// GetNonSigners() returns the public keys and corresponding percentage of voting power who are not included in the AggregateSignature
func (x *AggregateSignature) GetNonSigners(validatorList *ConsensusValidators) (nonSignerPubKeys [][]byte, nonSignerPercent int, err ErrorI) {
	// convert the consensus validators list into a validator set
	vs, err := NewValidatorSet(validatorList)
	// if an error occurred during the conversion
	if err != nil {
		// exit with error
		return
	}
	// get the non-signers from the aggregate signature using the validator set
	nonSignerPubKeys, nonSignerPower, err := x.getSigners(vs, true)
	// if an error occurred
	if err != nil {
		// exit with error
		return
	}
	// set the non signer percent (Non Signer Power / Total Power)
	nonSignerPercent = int(Uint64PercentageDiv(nonSignerPower, vs.TotalPower))
	// exit
	return
}

// getSigners() returns the public keys and corresponding combined voting power of signers or non-signers
func (x *AggregateSignature) getSigners(vs ValidatorSet, nonSigners bool) (pubkeys [][]byte, power uint64, err ErrorI) {
	// create a copy of the multi-public-key from the validator set
	key := vs.MultiKey.Copy()
	// set the 'who signed' bitmap in a copy of the key
	if e := key.SetBitmap(x.Bitmap); e != nil {
		// convert the error to a lib.ErrorI
		err = ErrInvalidSignerBitmap(e)
		// exit with error
		return
	}
	// iterate through the ValSet to and see if the validator signed
	for i, val := range vs.ValidatorSet.ValidatorSet {
		// did they sign?
		signed, er := key.SignerEnabledAt(i)
		// if an error occurred during this check
		if er != nil {
			// convert the error to a lib.ErrorI
			err = ErrInvalidSignerBitmap(er)
			// exit
			return
		}
		// if so, add to the pubkeys and add to the power
		if signed && !nonSigners || !signed && nonSigners {
			// add to list
			pubkeys = append(pubkeys, val.PublicKey)
			// update power
			power += val.VotingPower
		}
	}
	// exit
	return
}

// GetDoubleSigners() compares the signers of two signatures and return who signed both
func (x *AggregateSignature) GetDoubleSigners(y *AggregateSignature, vs ValidatorSet) (doubleSigners [][]byte, err ErrorI) {
	// create 2 copies of the multi public key for the validator set
	key, key2 := vs.MultiKey.Copy(), vs.MultiKey.Copy()
	// set the 'who signed' bitmap in key 1
	if er := key.SetBitmap(x.Bitmap); er != nil {
		// exit with error
		return nil, ErrInvalidSignerBitmap(er)
	}
	// set the 'who signed' bitmap in key 2
	if er := key2.SetBitmap(y.Bitmap); er != nil {
		// exit with error
		return nil, ErrInvalidSignerBitmap(er)
	}
	// iterate through the ValSet to and see if the validator signed
	for i, val := range vs.ValidatorSet.ValidatorSet {
		// see if the signer is enabled at index
		signed, e := key.SignerEnabledAt(i)
		// if an error occurred during this check
		if e != nil {
			// exit with error
			return nil, ErrInvalidSignerBitmap(e)
		}
		// if signed 1, check if they signed 2 as well
		if signed {
			// see if the signer is enabled at index
			signed, e = key2.SignerEnabledAt(i)
			// if an error occurred during this check
			if e != nil {
				// exit with error
				return nil, ErrInvalidSignerBitmap(e)
			}
			// if signed both, save as a double signer
			if signed {
				// add to the double signers list
				doubleSigners = append(doubleSigners, val.PublicKey)
			}
		}
	}
	// exit
	return
}

// jsonAggregateSig represents the json.Marshaller and json.Unmarshaler implementation of AggregateSignature
type jsonAggregateSig struct {
	Signature HexBytes `json:"signature,omitempty"`
	Bitmap    HexBytes `json:"bitmap,omitempty"`
}

// MarshalJSON() implements the json.Marshaller interface
func (x AggregateSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonAggregateSig{
		Signature: x.Signature,
		Bitmap:    x.Bitmap,
	})
}

// UnmarshalJSON() implements the json.Unmarshaler interface
func (x *AggregateSignature) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new json object reference to ensure a non nil result
	j := new(jsonAggregateSig)
	// populate the object using json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit with error
		return
	}
	// populate the underlying object using the json object reference
	x.Signature, x.Bitmap = j.Signature, j.Bitmap
	// exit
	return
}

// CONSENSUS VALIDATOR LOGIC BELOW

// Root() calculates the Merkle root of the ConsensusValidators
func (x *ConsensusValidators) Root() ([]byte, ErrorI) {
	// if the list of Consensus Validators is empty
	if x == nil || len(x.ValidatorSet) == 0 {
		// exit with an empty root
		return nil, nil
	}
	// create a list of bytes for the validators
	var valBytesList [][]byte
	// for each validator in the set
	for _, val := range x.ValidatorSet {
		// convert the object to bytes
		validatorBytes, err := Marshal(val)
		// if an error occurred
		if err != nil {
			// exit with error
			return nil, err
		}
		// add the bytes to the list
		valBytesList = append(valBytesList, validatorBytes)
	}
	// convert the list to a merkle tree and extract the root
	root, _, err := MerkleTree(valBytesList)
	// exit with the root
	return root, err
}

// marshalling utility structure for the ConsensusValidator
// allows easy hex byte marshalling of the public key
type jsonConsValidator struct {
	PublicKey   HexBytes `json:"publicKey,omitempty"`
	VotingPower uint64   `json:"votingPower,omitempty"`
	NetAddress  string   `json:"netAddress,omitempty"`
}

// MarshalJSON() overrides and implements the json.Marshaller interface
func (x ConsensusValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonConsValidator{
		PublicKey:   x.PublicKey,
		VotingPower: x.VotingPower,
		NetAddress:  x.NetAddress,
	})
}

// UnmarshalJSON() overrides and implements the json.Unmarshaller interface
func (x *ConsensusValidator) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new json object reference to ensure a non-nil result
	j := new(jsonConsValidator)
	// populate the json object ref using the json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit with error
		return
	}
	// populate the underlying object using the json object
	*x = ConsensusValidator{
		PublicKey:   j.PublicKey,
		VotingPower: j.VotingPower,
		NetAddress:  j.NetAddress,
	}
	// exit
	return
}

// ValidatorFilters are used to filter types of validators from a ValidatorPage
type ValidatorFilters struct {
	Unstaking FilterOption `json:"unstaking"` // validators are currently unstaking
	Paused    FilterOption `json:"paused"`    // validators are currently paused
	Delegate  FilterOption `json:"delegate"`  // validators are set as delegates
	Committee uint64       `json:"committee"` // validators are staked for this chain id (committee id)
}

// On() returns whether there exists any filters
func (v ValidatorFilters) On() bool {
	return v.Unstaking != FilterOption_Off || v.Paused != FilterOption_Off || v.Delegate != FilterOption_Off || v.Committee != 0
}

// FilterOption symbolizes 'condition must be true (yes)' 'condition must be false (no)' or 'filter off (both)' for filters
type FilterOption int

// nolint:all
const (
	FilterOption_Off     FilterOption = 0 // true or false condition
	FilterOption_MustBe               = 1 // condition must be true
	FilterOption_Exclude              = 2 // condition must be false
)

// VIEW CODE BELOW

func (x *View) CheckBasic() (err ErrorI) {
	// check if the view is empty
	if x == nil {
		// exit with empty error
		return ErrEmptyView()
	}
	// round and phase are not further checked,
	// because peers may be sending valid messages
	// asynchronously from different views
	return
}

// Check() checks the validity of the view and optionally enforce *heights* (plugin height and committee height)
func (x *View) Check(view *View, enforceHeights bool) ErrorI {
	// do a basic sanity check on the view
	if err := x.CheckBasic(); err != nil {
		// exit with err
		return err
	}
	// ensure the network id in the view matches expected
	if view.NetworkId != x.NetworkId {
		// exit with network id error
		return ErrWrongNetworkID()
	}
	// ensure the chain id in the view matches expected
	if view.ChainId != x.ChainId {
		// exit with chain id error
		return ErrWrongChainId()
	}
	// if enforcing heights, ensure the chain height is correct
	if enforceHeights && x.Height != view.Height {
		// exit with wrong height error
		return ErrWrongHeight()
	}
	// if enforcing heights, ensure root height is correct
	if enforceHeights && x.RootHeight != view.RootHeight {
		// exit with wrong root height error
		return ErrWrongRootHeight()
	}
	// exit
	return nil
}

// Copy() returns a reference to a clone of the View
func (x *View) Copy() *View {
	// clone the view
	return &View{
		Height:     x.Height,
		Round:      x.Round,
		Phase:      x.Phase,
		RootHeight: x.RootHeight,
		NetworkId:  x.NetworkId,
		ChainId:    x.ChainId,
	}
}

// Equals() returns true if this view is equal to the parameter view
// nil views are always false
func (x *View) Equals(v *View) bool {
	// if either of the views are empty
	if x == nil || v == nil {
		// exit with 'unequal'
		return false
	}
	// if the heights are not equal
	if x.Height != v.Height {
		// exit with 'unequal'
		return false
	}
	// if the root heights are not equal
	if x.RootHeight != v.RootHeight {
		// exit with 'unequal'
		return false
	}
	// if the chain ids are not equal
	if x.ChainId != v.ChainId {
		// exit with 'unequal'
		return false
	}
	// if the network ids are not equal
	if x.NetworkId != v.NetworkId {
		// exit with 'unequal'
		return false
	}
	// if the rounds are not equal
	if x.Round != v.Round {
		// exit with 'unequal'
		return false
	}
	// if the phases are not equal
	if x.Phase != v.Phase {
		// exit with 'unequal'
		return false
	}
	// exit with 'equal
	return true
}

// Less() returns true if this View is less than the parameter View
func (x *View) Less(v *View) bool {
	// if the comparable view is nil
	if v == nil {
		// exit with this view cannot be less than an empty view
		return false
	}
	// if this view is nil
	if x == nil {
		// exit with this view must be less than an empty view
		return true
	}
	// if height is less than the comparable
	if x.Height < v.Height {
		// exit with 'is less'
		return true
	}
	// if height is less than the comparable
	if x.Height > v.Height {
		// exit with 'not less'
		return false
	}
	// if root height is less than the comparable
	if x.RootHeight < v.RootHeight {
		// exit with 'is less'
		return true
	}
	// if the root height is greater than the comparable
	if x.RootHeight > v.RootHeight {
		// exit with 'not less'
		return false
	}
	// if round is less than the comparable
	if x.Round < v.Round {
		// exit with 'is less'
		return true
	}
	// if round is greater than the comparable
	if x.Round > v.Round {
		// exit with 'not less'
		return false
	}
	// if phase is less
	if x.Phase < v.Phase {
		return true
	}
	// exit with 'not less' (equal)
	return false
}

// ToString() returns the log string format of View
func (x *View) ToString() string {
	return fmt.Sprintf("(rH:%d, H:%d, R:%d, P:%s)", x.RootHeight, x.Height, x.Round, x.Phase)
}

// jsonView represents the json.Marshaller and json.Unmarshaler implementation of View
type jsonView struct {
	Height     uint64 `json:"height"`
	RootHeight uint64 `json:"committeeHeight"`
	Round      uint64 `json:"round"`
	Phase      string `json:"phase"` // string version of phase
	NetworkID  uint64 `json:"networkID"`
	ChainId    uint64 `json:"chainId"`
}

// MarshalJSON() implements the json.Marshaller interface
func (x View) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonView{
		Height:     x.Height,
		RootHeight: x.RootHeight,
		Round:      x.Round,
		Phase:      Phase_name[int32(x.Phase)],
		NetworkID:  x.NetworkId,
		ChainId:    x.ChainId,
	})
}

// MarshalJSON() implements the json.Marshaller interface
func (x *View) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new json object reference to ensure a non nil result
	j := new(jsonView)
	// populate the json object using json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit with error
		return
	}
	// populate the underlying object with the json object
	*x = View{
		NetworkId:  j.NetworkID,
		ChainId:    j.ChainId,
		Height:     j.Height,
		RootHeight: j.RootHeight,
		Round:      j.Round,
		Phase:      Phase(Phase_value[j.Phase]),
	}
	// exit
	return
}

// MarshalJSON() implements the json.Marshaller interface
func (x Phase) MarshalJSON() ([]byte, error) {
	return json.Marshal(Phase_name[int32(x)])
}

// UnmarshalJSON() implements the json.Unmarshaler interface
func (x *Phase) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new json string ref to ensure a non nil result
	j := new(string)
	// populate the json object using the json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit with error
		return
	}
	// populate the underlying object using the json string
	*x = Phase(Phase_value[*j])
	// exit
	return
}

// DOUBLE SIGNER CODE BELOW

// AddHeight() adds a height to the DoubleSigner
func (x *DoubleSigner) AddHeight(height uint64) {
	// ensure no duplicates
	if slices.Contains(x.Heights, height) {
		// exit without adding
		return
	}
	// add it to the list
	x.Heights = append(x.Heights, height)
}

// Equals() compares this DoubleSigner against the passed DoubleSigner
func (x *DoubleSigner) Equals(d *DoubleSigner) bool {
	// if both double signer objects are empty
	if x == nil && d == nil {
		// exit with 'equal'
		return true
	}
	// if either of the double signer objects are empty
	if x == nil || d == nil {
		// exit with 'unequal'
		return false
	}
	// if the double signer ids aren't equal
	if !bytes.Equal(x.Id, d.Id) {
		// exit with 'unequal'
		return false
	}
	// exit with if the height slices are equal
	return slices.Equal(x.Heights, d.Heights)
}

// proposersJSON implements the json.Marshaller and json.Unmarshaler interfaces for Proposer
type proposersJSON struct {
	Addresses []HexBytes `json:"addresses,omitempty"`
}

// MarshalJSON() implements the json.Marshaller interface for Proposers
func (x Proposers) MarshalJSON() ([]byte, error) {
	// create a new json object
	j := proposersJSON{}
	// for each address in the list
	for _, address := range x.Addresses {
		// add to the json object using the address
		j.Addresses = append(j.Addresses, address)
	}
	// convert the json object to bytes
	return json.Marshal(j)
}

// UnmarshalJSON() implements the json.Unmarshaler interface for Proposers
func (x *Proposers) UnmarshalJSON(jsonBytes []byte) (err error) {
	// create a new json object reference to ensure a non nil result
	j := new(proposersJSON)
	// populate the json object reference using the json bytes
	if err = json.Unmarshal(jsonBytes, j); err != nil {
		// exit
		return
	}
	// for each address in the json object reference
	for _, address := range j.Addresses {
		// add each address to the list of addresses
		x.Addresses = append(x.Addresses, address)
	}
	// exit
	return
}

// SortitionData is the seed data for the IsCandidate and VRF functions
type SortitionData struct {
	LastProposerAddresses [][]byte // the last N proposers addresses prevents any grinding attacks
	RootHeight            uint64   // the height of the root (optional) ensures leader rotation in a chain halt
	Height                uint64   // the height ensures unique proposer selection for each height
	Round                 uint64   // the round ensures unique proposer selection for each round
	TotalValidators       uint64   // the count of validators in the set
	TotalPower            uint64   // the total power of all validators in the set
	VotingPower           uint64   // the amount of voting power the node has
}

// PseudorandomParams are the input params to run the Stake-Weighted-Pseudorandom fallback leader selection algorithm
type PseudorandomParams struct {
	*SortitionData                      // seed data the peer used for sortition
	ValidatorSet   *ConsensusValidators // the set of validators
}

// WeightedPseudorandom() generates an index for the 'token' that the winner has in their stake
func WeightedPseudorandom(p *PseudorandomParams) (publicKey crypto.PublicKeyI) {
	// convert the seed data to a 16 byte hash, so it may fit in a uint64 type
	seed := FormatInputIntoSeed(p.LastProposerAddresses, p.RootHeight, p.Height, p.Round)[:16]
	// convert the seedBytes into a uint64 number
	seedUint64 := binary.BigEndian.Uint64(seed)
	// ensure that number falls within our 'Total Power'
	powerIndex := seedUint64 % p.TotalPower
	// create a power count variable to track the power index
	powerCount := uint64(0)
	// with this deterministically ordered validator set, iterate until exceeding the power index
	// as that Validator has the exact randomly chosen 'token' that is the lottery winner
	for _, v := range p.ValidatorSet.ValidatorSet {
		// add the voting power to the count
		powerCount += v.VotingPower
		// if exceed the powerIndex, that Validator has the exact 'token'
		if powerCount > powerIndex {
			// set the winner and exit
			publicKey, _ = crypto.BytesToBLS12381Public(v.PublicKey)
			// exit
			return
		}
	}
	// failsafe: should not happen - use the last validator from the set as the winner
	publicKey, _ = crypto.BytesToBLS12381Public(p.ValidatorSet.ValidatorSet[len(p.ValidatorSet.ValidatorSet)-1].PublicKey)
	// exit
	return
}

// FormatInputIntoSeed() returns the 'seed data' for the VRF function
// `seed = lastProposerAddresses + height + round`
func FormatInputIntoSeed(lastProposerAddresses [][]byte, rootHeight, height, round uint64) []byte {
	// create a string to hold the vrf input
	var input string
	// for each proposer address
	for _, address := range lastProposerAddresses {
		// add to the string using hex representation with a delimiter
		input += BytesToString(address) + "/"
	}
	// add the height and round to the end of the input
	input += fmt.Sprintf("%d/%d/%d", rootHeight, height, round)
	// hash the result
	return crypto.Hash([]byte(input))
}
