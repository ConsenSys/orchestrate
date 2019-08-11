package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/multi-vault.git/keystore"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-signer.git/handlers/vault/signer/generic"
)

// Signer produce a handler executing Tessera signature
func Signer(k keystore.KeyStore) engine.HandlerFunc {
	return generic.GenerateSignerHandler(
		signTx,
		k,
		"Successfully signed transaction for ethereum public transaction",
		"Public ethereum signer: could not sign the transaction: ",
	)
}

func signTx(s keystore.KeyStore, txctx *engine.TxContext, sender common.Address, t *ethtypes.Transaction) ([]byte, *common.Hash, error) {
	b, hash, err := s.SignTx(txctx.Envelope.GetChain(), sender, t)
	if err != nil {
		return b, hash, errors.FromError(err).ExtendComponent(component)
	}
	return b, hash, nil
}
