package chainregistry

import (
	"math/big"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Faucet struct {
	UUID       string
	Amount     *big.Int
	MaxBalance *big.Int
	Cooldown   time.Duration
	Creditor   ethcommon.Address
}
