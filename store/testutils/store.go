package testutils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/api/context-store.git/store"
	common "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/common"
	envelope "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/envelope"
	ethereum "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/ethereum"
)

// EnvelopeStoreTestSuite is a test suit for EnvelopeStore
type EnvelopeStoreTestSuite struct {
	suite.Suite
	Store store.EnvelopeStore
}

// TestEnvelopeStore test envelope store
func (s *EnvelopeStoreTestSuite) TestEnvelopeStore() {
	txData := (&ethereum.TxData{}).
		SetNonce(10).
		SetTo(ethcommon.HexToAddress("0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C")).
		SetValue(big.NewInt(100000)).
		SetGas(2000).
		SetGasPrice(big.NewInt(200000)).
		SetData(hexutil.MustDecode("0xabcd"))

	tr := &envelope.Envelope{
		Chain:    &common.Chain{Id: "0x3"},
		Metadata: &envelope.Metadata{Id: "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11"},
		Tx: &ethereum.Transaction{
			TxData: txData,
			Raw:    "0xf86c0184ee6b280082529094ff778b716fc07d98839f48ddb88d8be583beb684872386f26fc1000082abcd29a0d1139ca4c70345d16e00f624622ac85458d450e238a48744f419f5345c5ce562a05bd43c512fcaf79e1756b2015fec966419d34d2a87d867b9618a48eca33a1a80",
			Hash:   "0x0a0cafa26ca3f411e6629e9e02c53f23713b0033d7a72e534136104b5447a210",
		},
	}

	// Store Envelope
	status, storedAt, err := s.Store.Store(context.Background(), tr)
	assert.Nil(s.T(), err, "Should properly store envelope")
	assert.Equal(s.T(), "stored", status, "Default status should be correct")
	assert.True(s.T(), time.Since(storedAt) < 5*time.Second, "Stored date should be close")

	// Load Envelope
	tr = &envelope.Envelope{}
	status, _, err = s.Store.LoadByTxHash(context.Background(), "0x3", "0x0a0cafa26ca3f411e6629e9e02c53f23713b0033d7a72e534136104b5447a210", tr)
	assert.Nil(s.T(), err, "Should properly store envelope")
	assert.Equal(s.T(), "stored", status, "Status should be correct")
	assert.Equal(s.T(), "0x3", tr.GetChain().GetId(), "ChainID should be correct")
	assert.Equal(s.T(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11", tr.GetMetadata().GetId(), "MetadataID should be correct")

	// Set Status
	err = s.Store.SetStatus(context.Background(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11", "pending")
	assert.Nil(s.T(), err, "Setting status to %q", "pending")

	status, sentAt, err := s.Store.GetStatus(context.Background(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11")
	assert.Nil(s.T(), err, "Should not error")
	assert.Equal(s.T(), "pending", status, "Status should be correct")
	assert.True(s.T(), sentAt.Sub(storedAt) > 0, "Stored should be older than sent date")

	// Stores an already existing
	tr = &envelope.Envelope{
		Chain:    &common.Chain{Id: "0x3"},
		Metadata: &envelope.Metadata{Id: "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11"},
		Tx: &ethereum.Transaction{
			TxData: txData,
			Raw:    "0xf86c0184ee6b280082529094ff778b716fc07d98839f48ddb88d8be583beb684872386f26fc1000082abcd29a0d1139ca4c70345d16e00f624622ac85458d450e238a48744f419f5345c5ce562a05bd43c512fcaf79e1756b2015fec966419d34d2a87d867b9618a48eca33a1a80",
			Hash:   "0x0a0cafa26ca3f411e6629e9e02c53f23713b0033d7a72e534136104b5447a210",
		},
	}

	status, _, err = s.Store.Store(context.Background(), tr)
	assert.Nil(s.T(), err, "Should update")
	assert.Equal(s.T(), "pending", status, "Status should be correct")

	// Set status to error
	err = s.Store.SetStatus(context.Background(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11", "error")
	assert.Nil(s.T(), err, "Setting status to %q", "error")

	status, errorAt, err := s.Store.GetStatus(context.Background(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11")
	assert.Nil(s.T(), err, "Should not error")
	assert.Equal(s.T(), "error", status, "Status should be correct")
	assert.True(s.T(), errorAt.Sub(sentAt) > 0, "Stored date should be close")

	// Test to Load By ID
	status, _, err = s.Store.LoadByID(context.Background(), "a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a11", tr)
	assert.Nil(s.T(), err, "Should not error")
	assert.Equal(s.T(), "error", status, "Status should be correct")
}

// TestLoadPending test load pending envelopes
func (s *EnvelopeStoreTestSuite) TestLoadPending() {
	txData := (&ethereum.TxData{}).
		SetNonce(10).
		SetTo(ethcommon.HexToAddress("0xAf84242d70aE9D268E2bE3616ED497BA28A7b62C")).
		SetValue(big.NewInt(100000)).
		SetGas(2000).
		SetGasPrice(big.NewInt(200000)).
		SetData(hexutil.MustDecode("0xabcd"))

	for i, chain := range []string{"0x1", "0x2", "0x3", "0xa2", "0x42", "0xab"} {
		tr := &envelope.Envelope{
			Chain:    &common.Chain{Id: chain},
			Metadata: &envelope.Metadata{Id: fmt.Sprintf("a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a1%v", i)},
			Tx: &ethereum.Transaction{
				TxData: txData,
				Raw:    "0xf86c0184ee6b280082529094ff778b716fc07d98839f48ddb88d8be583beb684872386f26fc1000082abcd29a0d1139ca4c70345d16e00f624622ac85458d450e238a48744f419f5345c5ce562a05bd43c512fcaf79e1756b2015fec966419d34d2a87d867b9618a48eca33a1a80",
				Hash:   "0x0a0cafa26ca3f411e6629e9e02c53f23713b0033d7a72e534136104b5447a210",
			},
		}

		_, _, err := s.Store.Store(context.Background(), tr)
		assert.Nil(s.T(), err, "No error expected")
		time.Sleep(100 * time.Millisecond)

		if i%2 == 0 {
			err = s.Store.SetStatus(context.Background(), fmt.Sprintf("a0ee-bc99-9c0b-4ef8-bb6d-6bb9-bd38-0a1%v", i), "pending")
			assert.Nil(s.T(), err, "No error expected")
		}
	}

	envelopes, err := s.Store.LoadPending(context.Background(), 0)
	assert.Nil(s.T(), err, "No error expected on LoadPending")
	assert.Len(s.T(), envelopes, 3, "Count of envelope pending incorrect")

	envelopes, err = s.Store.LoadPending(context.Background(), 300*time.Millisecond)
	assert.Nil(s.T(), err, "No error expected on LoadPending")
	assert.Len(s.T(), envelopes, 2, "Count of envelope pending incorrect")

	envelopes, err = s.Store.LoadPending(context.Background(), 500*time.Millisecond)
	assert.Nil(s.T(), err, "No error expected on LoadPending")
	assert.Len(s.T(), envelopes, 1, "Count of envelope pending incorrect")

	envelopes, err = s.Store.LoadPending(context.Background(), 700*time.Millisecond)
	assert.Nil(s.T(), err, "No error expected on LoadPending")
	assert.Len(s.T(), envelopes, 0, "Count of envelope pending incorrect")
}
