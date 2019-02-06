package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Log holds data about a log
type Log struct {
	types.Log
	DecodedData map[string]string
}

// Receipt holds Data about a receipt
type Receipt struct {
	types.Receipt
	Logs []*Log

	BlockHash   common.Hash
	BlockNumber uint64
	TxIndex     uint64
}

// NewReceipt creates a new receipt
func newReceipt(root []byte, failed bool, cumulativeGasUsed uint64) *Receipt {
	return &Receipt{
		Receipt: *types.NewReceipt(root, failed, cumulativeGasUsed),
		Logs:    make([]*Log, 0),
	}
}

func (r *Receipt) reset() {
	r.PostState = r.PostState[0:0]
	r.Status = 0
	r.CumulativeGasUsed = 0
	r.Bloom.SetBytes([]byte{})
	r.Logs = r.Logs[0:0]
	r.TxHash.SetBytes([]byte{})
	r.ContractAddress.SetBytes([]byte{})
	r.GasUsed = 0
	r.BlockHash.SetBytes([]byte{})
	r.BlockNumber = 0
	r.TxIndex = 0
}

// SetDecodedData set DecodedData to log
func (l *Log) SetDecodedData(m map[string]string) {
	l.DecodedData = m
}
