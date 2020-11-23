package utils

import (
	"encoding/json"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	proto "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/ethereum"
)

type ProcessResultFunc func(result json.RawMessage) error

func ProcessResult(v interface{}) ProcessResultFunc {
	return func(result json.RawMessage) error {
		err := json.Unmarshal(result, &v)
		if err != nil {
			return errors.EncodingError(err.Error())
		}

		return nil
	}
}

func ProcessReceiptResult(receipt **proto.Receipt) ProcessResultFunc {
	return func(result json.RawMessage) error {
		err := ProcessResult(&receipt)(result)
		if err != nil {
			return err
		}

		if receipt == nil || *receipt == nil {
			// Receipt was not found
			return errors.NotFoundError("receipt not found")
		}

		return nil
	}
}
