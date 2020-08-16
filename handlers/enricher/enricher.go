package enricher

import (
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethclient"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/proxy"
	svc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/proto"
)

// Enricher is a Middleware engine.HandlerFunc
func Enricher(r svc.ContractRegistryClient, ec ethclient.ChainStateReader) engine.HandlerFunc {
	return func(txctx *engine.TxContext) {
		if txctx.Envelope.GetReceipt().GetContractAddress() != "" {
			url, err := proxy.GetURL(txctx)
			if err != nil {
				return
			}

			code, err := ec.CodeAt(txctx.Context(),
				url,
				txctx.Envelope.GetReceipt().GetContractAddr(),
				nil)
			if err != nil {
				_ = txctx.AbortWithError(errors.InternalError(
					"could not read account code for chain %s and account %s",
					url,
					txctx.Envelope.GetReceipt().GetContractAddr().Hex(),
				)).SetComponent(component)
				return
			}

			_, err = r.SetAccountCodeHash(txctx.Context(),
				&svc.SetAccountCodeHashRequest{
					AccountInstance: &common.AccountInstance{},
					CodeHash:        crypto.Keccak256Hash(code).String(),
				},
			)
			if err != nil {
				_ = txctx.AbortWithError(errors.InternalError("invalid input message format")).
					SetComponent(component)
				return
			}
			txctx.Logger.Debugf("%s successfully SetAccountCodeHash in Contract Registry for chain %s and account %s with codehash",
				txctx.Envelope.GetChainIDString(),
				txctx.Envelope.GetReceipt().GetContractAddress(),
				crypto.Keccak256Hash(code))
		}
	}
}
