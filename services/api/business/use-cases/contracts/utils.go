package contracts

import (
	ethabi "github.com/consensys/orchestrate/pkg/go-ethereum/v1_9_12/accounts/abi"
)

// returns the selector associated to a signature hash
func sigHashToSelector(data []byte) (res [4]byte) {
	copy(res[:], data)
	return res
}

// returns the count of indexed inputs in the event
func getIndexedCount(event ethabi.Event) (indexedInputCount uint) {
	for i := range event.Inputs {
		if event.Inputs[i].Indexed {
			indexedInputCount++
		}
	}

	return indexedInputCount
}
