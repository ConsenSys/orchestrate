package handlers

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/core/worker"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/protos/ethereum"
)

var noErrorChainID = int64(0)
var error1ChainID = int64(1)
var error2ChainID = int64(2)
var error3ChainID = int64(3)
var error4ChainID = int64(4)
var error5ChainID = int64(5)

var chainIDs = []int64{
	noErrorChainID,
	error1ChainID,
	error2ChainID,
	error3ChainID,
	error4ChainID,
	error5ChainID,
}

var nonceInCacheAddress = "0x1234608A02a7A15fd6689D6DaD560C44E9ab61Ff"
var nonceNotInCacheAddress = "0xfF778b716FC07D98839f48DdB88D8bE583BEB684"
var nonceTooOldAddress = "0x6009608A02a7A15fd6689D6DaD560C44E9ab61Ff"

var addresses = []string{
	nonceInCacheAddress,
	nonceNotInCacheAddress,
	nonceTooOldAddress,
}

var cacheNonce = uint64(53)
var chainNonce = uint64(42)

type MockNonceGetter struct {
	counter uint64
}

func (g *MockNonceGetter) GetNonce(ctx context.Context, chainID *big.Int, a ethcommon.Address) (uint64, error) {
	atomic.AddUint64(&g.counter, 1)
	if chainID.Int64() == error1ChainID {
		// Simulate error on chain 0
		return 0, fmt.Errorf("Unknwon chain")
	}
	return chainNonce, nil
}

type MockNonceManager struct {
	mux   *sync.Mutex
	nonce *sync.Map
}

func (nm *MockNonceManager) GetNonce(chainID *big.Int, a *ethcommon.Address) (uint64, int, error) {
	if chainID.Int64() == error2ChainID {
		// Simulate error
		return 0, 0, fmt.Errorf("Error retrieving nonce")
	}

	if a.Hex() == nonceNotInCacheAddress {
		// Simulate unknown nonce
		return 0, -1, nil
	}

	if a.Hex() == nonceTooOldAddress {
		// Simulate nonce that is too old
		return 0, 1000, nil
	}

	return cacheNonce, 1, nil
}

func (nm *MockNonceManager) SetNonce(chainID *big.Int, a *ethcommon.Address, value uint64) error {
	if chainID.Int64() == error3ChainID {
		// Simulate error
		return fmt.Errorf("Error setting nonce")
	}
	return nil
}

func (nm *MockNonceManager) Lock(chainID *big.Int, a *ethcommon.Address) (string, error) {
	if chainID.Int64() == error4ChainID {
		// Simulate error
		return "", fmt.Errorf("Error locking nonce")
	}
	nm.mux.Lock()
	return "random", nil
}

func (nm *MockNonceManager) Unlock(chainID *big.Int, a *ethcommon.Address, lockSig string) error {
	nm.mux.Unlock()
	if chainID.Int64() == error5ChainID {
		// Simulate error
		return fmt.Errorf("Error unlocking nonce")
	}
	return nil
}

func makeNonceContext(chainID int64, address string) *worker.Context {
	ctx := worker.NewContext()
	ctx.Reset()
	ctx.Logger = log.NewEntry(log.StandardLogger())
	ctx.T.Chain = (&common.Chain{}).SetID(big.NewInt(chainID))
	ctx.T.Sender = &common.Account{Addr: address}
	ctx.T.Tx = &ethereum.Transaction{TxData: &ethereum.TxData{}}

	if chainID == noErrorChainID {
		ctx.Keys["expectedErrorCount"] = 0
	} else if chainID == error1ChainID && address == nonceInCacheAddress {
		// If nonce is in cache, there is no calibration
		// Thus an error when getting nonce from chain is not seen
		ctx.Keys["expectedErrorCount"] = 0
	} else {
		ctx.Keys["expectedErrorCount"] = 1
	}

	if address == nonceInCacheAddress {
		ctx.Keys["expectedNonce"] = cacheNonce
	} else {
		ctx.Keys["expectedNonce"] = chainNonce
	}

	return ctx
}

func TestNonceHandler(t *testing.T) {
	viper.Set("redis.nonce.expiration.time", "3")
	nm := MockNonceManager{
		mux: &sync.Mutex{},
	}
	ng := MockNonceGetter{}
	nh := NonceHandler(&nm, ng.GetNonce)

	rounds := 10
	outs := make(chan *worker.Context, rounds*len(addresses)*len(chainIDs))
	wg := &sync.WaitGroup{}
	for i := 0; i < rounds; i++ {
		for _, a := range addresses {
			for _, c := range chainIDs {
				wg.Add(1)
				ctx := makeNonceContext(c, a)
				go func(ctx *worker.Context) {
					defer wg.Done()
					nh(ctx)
					outs <- ctx
				}(ctx)
			}
		}
	}
	wg.Wait()
	close(outs)

	if len(outs) != rounds*len(addresses)*len(chainIDs) {
		t.Errorf("NonceHandler: expected %v outs but got %v", rounds*len(addresses)*len(chainIDs), len(outs))
	}

	for ctx := range outs {
		if ctx.Keys["expectedErrorCount"].(int) > 0 {
			if len(ctx.T.GetErrors()) != ctx.Keys["expectedErrorCount"].(int) {
				t.Errorf("Expected %v errors but got %v %v", ctx.Keys["expectedErrorCount"].(int), ctx.T.GetErrors(), ctx.T.GetSender().Addr)
			}
		} else {
			if ctx.T.Tx.TxData.GetNonce() != ctx.Keys["expectedNonce"].(uint64) {
				t.Errorf("Expected Nonce to be %v but got %v", ctx.Keys["expectedNonce"].(uint64), ctx.T.Tx.TxData.GetNonce())
			}
		}
	}
}
