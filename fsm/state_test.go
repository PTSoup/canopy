package fsm

import (
	"encoding/json"
	"github.com/canopy-network/canopy/fsm/types"
	"github.com/canopy-network/canopy/lib"
	"github.com/canopy-network/canopy/lib/crypto"
	"github.com/canopy-network/canopy/store"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestInitialize(t *testing.T) {
	const dataDirPath = "./"
	tests := []struct {
		name          string
		detail        string
		presetBlock   *lib.Block
		presetGenesis *types.GenesisState
		height        uint64
		expected      *types.GenesisState
	}{
		{
			name:        "after genesis",
			detail:      "the block height is after 0, thus it's the non-genesis initialization",
			height:      2,
			presetBlock: &lib.Block{BlockHeader: &lib.BlockHeader{Height: 1, Hash: crypto.Hash([]byte("test")), TotalVdfIterations: 2}},
		},
		{
			name:          "genesis path",
			detail:        "the height is 0 so the genesis path is taken",
			presetGenesis: newTestGenesisState(t),
			expected:      newTestValidateGenesisState(t),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// create the default logger
			log := lib.NewDefaultLogger()
			// create an in-memory store
			db, err := store.NewStoreInMemory(log)
			require.NoError(t, err)
			if test.presetGenesis != nil {
				// marshal genesis file to bytes
				genesisJsonBytes, e := json.MarshalIndent(&test.presetGenesis, "", "  ")
				require.NoError(t, e)
				// write test genesis to file
				require.NoError(t, os.WriteFile("genesis.json", genesisJsonBytes, 0777))
				// remove the test file
				defer os.RemoveAll("genesis.json")
			}
			if test.presetBlock != nil {
				// set the block in state
				require.NoError(t, db.IndexBlock(&lib.BlockResult{
					BlockHeader: test.presetBlock.BlockHeader,
				}))
			}
			if test.height != 0 {
				// increment the db version
				_, _ = db.Commit()
			}
			// create a state machine object
			sm := StateMachine{
				store:  db,
				height: test.height,
				Config: lib.Config{},
				log:    log,
			}
			// set the data dir path
			sm.Config.DataDirPath = dataDirPath
			// execute the function call
			require.NoError(t, sm.Initialize(db))
			// validate the initialization path
			if test.height == 0 {
				// if genesis, validate the state
				validateWithExportedState(t, sm, test.expected)
			} else {
				// if not genesis, validate the VDF iterations
				require.Equal(t, test.presetBlock.BlockHeader.TotalVdfIterations, sm.vdfIterations)
			}
		})
	}
}

func TestApplyBlock(t *testing.T) {
	var timestamp = uint64(time.Date(2024, 02, 01, 0, 0, 0, 0, time.UTC).UnixMicro())
	// define a key group to use in testing
	kg := newTestKeyGroup(t)
	// predefine a send-transaction to insert into the block
	sendTx, err := types.NewSendTransaction(kg.PrivateKey, newTestAddress(t), 1, 1, 1, 1, "")
	txn := sendTx.(*lib.Transaction)
	// set the timestamp to a fixed time for validity checking
	txn.Time = timestamp
	// re-sign the tx
	require.NoError(t, txn.Sign(kg.PrivateKey))
	// ensure no error
	require.NoError(t, err)
	// convert the object to bytes
	sendTxBytes, err := lib.Marshal(sendTx)
	// ensure no error
	require.NoError(t, err)
	// define test cases
	tests := []struct {
		name            string
		detail          string
		accountPreset   uint64
		storeError      bool
		beginBlockError bool
		block           *lib.Block
		expectedHeader  *lib.BlockHeader
		expectedResults *lib.TxResult
		error           string
	}{
		{
			name:       "store error",
			detail:     "an error occurred in casting the store to lib.Store",
			storeError: true,
			error:      "wrong store type",
		},
		{
			name:            "begin_block error",
			detail:          "an error occurred in begin block",
			block:           &lib.Block{BlockHeader: &lib.BlockHeader{}, Transactions: [][]byte{sendTxBytes}},
			beginBlockError: true,
			error:           "invalid protocol version",
		},
		{
			name:   "transaction error",
			detail: "an error occurred in the transaction",
			block:  &lib.Block{BlockHeader: &lib.BlockHeader{}, Transactions: [][]byte{sendTxBytes}},
			error:  "insufficient funds",
		},
		{
			name:          "successful apply block",
			detail:        "the happy path with apply block without a 'last quorum certificate'",
			accountPreset: 2,
			block: &lib.Block{
				BlockHeader: &lib.BlockHeader{
					Height:          2,
					NumTxs:          1,
					Time:            timestamp,
					TotalTxs:        1,
					LastBlockHash:   crypto.Hash([]byte("block_hash")),
					ProposerAddress: newTestAddressBytes(t),
				},
				Transactions: [][]byte{sendTxBytes},
			},
			expectedHeader: &lib.BlockHeader{
				Height:                3,
				NetworkId:             1,
				Time:                  timestamp,
				NumTxs:                1,
				TotalTxs:              1,
				TotalVdfIterations:    0,
				Hash:                  []byte{0x29, 0xe1, 0x86, 0x8a, 0x7d, 0x26, 0x2e, 0xd9, 0x88, 0xc6, 0x2b, 0xa8, 0xef, 0xab, 0x48, 0x9a, 0x4, 0x82, 0x1c, 0x2d, 0xcd, 0x18, 0x73, 0x21, 0x26, 0xc7, 0x91, 0x15, 0x18, 0x49, 0x7, 0x5e},
				LastBlockHash:         []byte{0x26, 0x46, 0xe, 0xd3, 0x76, 0x17, 0x95, 0x7c, 0x96, 0xd9, 0xab, 0xf5, 0x94, 0xa1, 0xac, 0x86, 0x5a, 0x43, 0x11, 0x2, 0xfc, 0x38, 0x77, 0x71, 0xa8, 0xc7, 0x6d, 0xa0, 0x2e, 0x6f, 0x1, 0xe8},
				StateRoot:             []byte{0x8c, 0x3a, 0x7f, 0xbe, 0xb6, 0x26, 0x63, 0xd0, 0xfa, 0x12, 0x64, 0x8a, 0xc4, 0x7a, 0x97, 0xa5, 0xa0, 0x6d, 0x45, 0xb5},
				TransactionRoot:       []byte{0x99, 0x1f, 0x2c, 0x1c, 0x67, 0x25, 0xfd, 0x97, 0x9e, 0x4c, 0x5e, 0xbd, 0xdb, 0x38, 0x23, 0xe, 0x8f, 0x68, 0x4d, 0xb5, 0xc4, 0x1b, 0xcc, 0xfe, 0x7d, 0x29, 0xb9, 0x57, 0x62, 0xe3, 0x4b, 0x6d},
				ValidatorRoot:         []byte{0x24, 0xa5, 0xf1, 0x5d, 0xdd, 0x13, 0xdd, 0x75, 0x33, 0x2a, 0xe4, 0xf6, 0x2b, 0x3f, 0xa, 0x8c, 0xdf, 0x90, 0x1d, 0x9f, 0xaa, 0xb0, 0x5d, 0xae, 0x7a, 0x47, 0xa7, 0x59, 0x98, 0x64, 0xb3, 0x7c},
				NextValidatorRoot:     []byte{0x24, 0xa5, 0xf1, 0x5d, 0xdd, 0x13, 0xdd, 0x75, 0x33, 0x2a, 0xe4, 0xf6, 0x2b, 0x3f, 0xa, 0x8c, 0xdf, 0x90, 0x1d, 0x9f, 0xaa, 0xb0, 0x5d, 0xae, 0x7a, 0x47, 0xa7, 0x59, 0x98, 0x64, 0xb3, 0x7c},
				ProposerAddress:       newTestAddressBytes(t),
				Vdf:                   nil,
				LastQuorumCertificate: nil,
			},
			expectedResults: &lib.TxResult{
				Sender:      newTestAddressBytes(t),
				Recipient:   newTestAddressBytes(t),
				MessageType: "send",
				Height:      3,
				Index:       0,
				Transaction: sendTx.(*lib.Transaction),
				TxHash:      crypto.HashString(sendTxBytes),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// create a state machine instance with default parameters
			sm := newTestStateMachine(t)
			if test.storeError {
				// set the store to the wrong type
				sm.store = lib.RWStoreI(nil)
			} else {
				// preset the 'last block' in state
				require.NoError(t, sm.store.(lib.StoreI).IndexBlock(&lib.BlockResult{
					BlockHeader: &lib.BlockHeader{
						Height: 2,
						Hash:   test.block.BlockHeader.LastBlockHash,
						Time:   timestamp,
					},
				}))
				// set the minimum fee to 1 for send transactions
				require.NoError(t, sm.UpdateParam("fee", types.ParamMessageSendFee, &lib.UInt64Wrapper{Value: 1}))
				// preset the account with funds
				require.NoError(t, sm.AccountAdd(newTestAddress(t), test.accountPreset))
				qc := &lib.QuorumCertificate{
					Header: &lib.View{Height: 2},
					Results: &lib.CertificateResult{RewardRecipients: &lib.RewardRecipients{
						PaymentPercents: []*lib.PaymentPercents{{
							Address: newTestAddressBytes(t),
							Percent: 100,
						}},
					}},
				}
				// track the supply
				supply := &types.Supply{}
				// for 4 validators
				for i := 0; i < 4; i++ {
					// set the validator
					require.NoError(t, sm.SetValidators([]*types.Validator{{
						Address:      newTestAddressBytes(t, i),
						PublicKey:    newTestPublicKeyBytes(t, i),
						StakedAmount: 100,
						Committees:   []uint64{lib.CanopyCommitteeId},
						Output:       newTestAddressBytes(t),
					}}, supply))
					// set the committee member
					require.NoError(t, sm.SetCommitteeMember(newTestAddress(t, i), lib.CanopyCommitteeId, 100))
				}
				// set the supply in state
				require.NoError(t, sm.SetSupply(supply))
				// create an aggregate signature
				// get the committee members
				committee, er := sm.GetCommitteeMembers(lib.CanopyCommitteeId, true)
				require.NoError(t, er)
				// create a copy of the multikey
				mk := committee.MultiKey.Copy()
				// only sign with 3/4 to test the non-signer reduction
				for i := 0; i < 3; i++ {
					privateKey := newTestKeyGroup(t, i).PrivateKey
					// search for the proper index for the signer
					for j, pubKey := range mk.PublicKeys() {
						// if found, add the signer
						if privateKey.PublicKey().Equals(pubKey) {
							// sign the qc
							require.NoError(t, mk.AddSigner(privateKey.Sign(qc.SignBytes()), j))
						}
					}
				}
				// aggregate the signature
				aggSig, e := mk.AggregateSignatures()
				require.NoError(t, e)
				// attach the signature to the message
				qc.Signature = &lib.AggregateSignature{
					Signature: aggSig,
					Bitmap:    mk.Bitmap(),
				}
				require.NoError(t, sm.store.(lib.StoreI).IndexQC(qc))
				// setup for a 'last validator set' for apply block
				sm.height = 3
				// ommit here to have a 'last validator set' for apply block
				_, err = sm.store.(lib.StoreI).Commit()
				require.NoError(t, err)
			}
			if !test.beginBlockError {
				// set the protocol version to not trigger an error
				sm.ProtocolVersion = 1
			}
			// execute the function call
			header, txResults, e := sm.ApplyBlock(test.block)
			// validate the expected error
			require.Equal(t, test.error != "", e != nil, e)
			if e != nil {
				require.ErrorContains(t, e, test.error)
				return
			}
			// validate got vs expected block header
			require.EqualExportedValues(t, test.expectedHeader, header)
			// validate got vs expected tx results
			require.EqualExportedValues(t, test.expectedResults, txResults[0])
		})
	}
}

func newSingleAccountStateMachine(t *testing.T) StateMachine {
	sm := newTestStateMachine(t)
	keyGroup := newTestKeyGroup(t)
	require.NoError(t, sm.SetParams(types.DefaultParams()))
	require.NoError(t, sm.SetAccount(&types.Account{
		Address: keyGroup.Address.Bytes(),
		Amount:  1000000,
	}))
	require.NoError(t, sm.HandleMessageStake(&types.MessageStake{
		PublicKey:     keyGroup.PublicKey.Bytes(),
		Amount:        1000000,
		Committees:    []uint64{lib.CanopyCommitteeId},
		NetAddress:    "http://localhost:80",
		OutputAddress: keyGroup.Address.Bytes(),
		Delegate:      false,
		Compound:      true,
	}))
	require.NoError(t, sm.SetParams(types.DefaultParams()))
	return sm
}

func newTestStateMachine(t *testing.T) StateMachine {
	log := lib.NewDefaultLogger()
	db, err := store.NewStoreInMemory(log)
	require.NoError(t, err)
	sm := StateMachine{
		store:             db,
		ProtocolVersion:   0,
		NetworkID:         1,
		height:            2,
		vdfIterations:     0,
		slashTracker:      types.NewSlashTracker(),
		proposeVoteConfig: types.AcceptAllProposals,
		Config: lib.Config{
			MainConfig: lib.DefaultMainConfig(),
		},
		log: log,
	}
	require.NoError(t, sm.SetParams(types.DefaultParams()))
	db.Commit()
	require.NoError(t, sm.SetParams(types.DefaultParams()))
	return sm
}

func newTestAddress(t *testing.T, variation ...int) crypto.AddressI {
	kg := newTestKeyGroup(t, variation...)
	return kg.Address
}

func newTestAddressBytes(t *testing.T, variation ...int) []byte {
	return newTestAddress(t, variation...).Bytes()
}

func newTestPublicKey(t *testing.T, variation ...int) crypto.PublicKeyI {
	kg := newTestKeyGroup(t, variation...)
	return kg.PublicKey
}

func newTestPublicKeyBytes(t *testing.T, variation ...int) []byte {
	return newTestPublicKey(t, variation...).Bytes()
}

func newTestKeyGroup(t *testing.T, variation ...int) *crypto.KeyGroup {
	var (
		key  crypto.PrivateKeyI
		err  error
		keys = []string{
			"01553a101301cd7019b78ffa1186842dd93923e563b8ae22e2ab33ae889b23ee",
			"1b6b244fbdf614acb5f0d00a2b56ffcbe2aa23dabd66365dffcd3f06491ae50a",
			"2ee868f74134032eacba191ca529115c64aa849ac121b75ca79b37420a623036",
			"3e3ab94c10159d63a12cb26aca4b0e76070a987d49dd10fc5f526031e05801da",
			"479839d3edbd0eefa60111db569ded6a1a642cc84781600f0594bd8d4a429319",
			"51eb5eb6eca0b47c8383652a6043aadc66ddbcbe240474d152f4d9a7439eae42",
			"637cb8e916bba4c1773ed34d89ebc4cb86e85c145aea5653a58de930590a2aa4",
			"7235e5757e6f52e6ae4f9e20726d9c514281e58e839e33a7f667167c524ff658"}
	)

	if len(variation) == 1 {
		key, err = crypto.StringToBLS12381PrivateKey(keys[variation[0]])
	} else {
		key, err = crypto.StringToBLS12381PrivateKey(keys[0])
	}
	require.NoError(t, err)
	return crypto.NewKeyGroup(key)
}

func newTestKeyGroups(t *testing.T, count int) (groups []*crypto.KeyGroup) {
	for i := 0; i < count; i++ {
		groups = append(groups, newTestKeyGroup(t, i))
	}
	return
}

// testQCParams are the associate parameters needed to generate a testQC
type testQCParams struct {
	height        uint64
	idxSigned     map[int]bool
	committeeKeys []*crypto.KeyGroup
	committee     []*types.Validator
	results       *lib.CertificateResult
}

// newTestQC is a utility function for this test to generate various quorum certificates in the test cases
func newTestQC(t *testing.T, params testQCParams) (qc *lib.QuorumCertificate) {
	// convert committee members to consensus validators
	var vals []*lib.ConsensusValidator
	for _, m := range params.committee {
		vals = append(vals, &lib.ConsensusValidator{PublicKey: m.PublicKey, VotingPower: m.StakedAmount})
	}
	// create a validator set object in order to generate a multi-public key for the set
	validatorSet, err := lib.NewValidatorSet(&lib.ConsensusValidators{ValidatorSet: vals})
	require.NoError(t, err)
	// create the 'justification' object
	justification := validatorSet.MultiKey.Copy()
	// create the certificate results object to put in the QC
	// create the QC object
	qc = &lib.QuorumCertificate{
		Header: &lib.View{
			Height:       params.height,
			CanopyHeight: params.height,
			CommitteeId:  lib.CanopyCommitteeId,
		},
		Results:     params.results,
		ResultsHash: params.results.Hash(),
		BlockHash:   crypto.Hash([]byte("some block that's not included here")),
	}
	// generate the bytes to be signed in the justification (multi-key)
	bytesToBeSigned := qc.SignBytes()
	// have the 'signers' sign the justification (multi-key)
	for i, s := range params.committeeKeys {
		if params.idxSigned[i] {
			require.NoError(t, justification.AddSigner(s.PrivateKey.Sign(bytesToBeSigned), i))
		}
	}
	// aggregate the signature
	aggregateSignatures, e := justification.AggregateSignatures()
	require.NoError(t, e)
	// wrap in object
	qc.Signature = &lib.AggregateSignature{
		Signature: aggregateSignatures,
		Bitmap:    justification.Bitmap(),
	}
	return
}
