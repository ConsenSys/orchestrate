package keystore

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/multi-vault.git/keystore/base"
	"gitlab.com/ConsenSys/client/fr/core-stack/service/multi-vault.git/secretstore"
)

var (
	keyStore KeyStore
	initOnce = &sync.Once{}
)

// Init initialize Key Store
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if keyStore != nil {
			return
		}

		secretstore.Init(ctx)
		keyStore = base.NewKeyStore(secretstore.GlobalSecretStore())

		err := ImportPrivateKey(keyStore)
		if err != nil {
			log.Fatalf("Key Store: Cannot import private keys, got error: %q", err)
		}

		log.Info("Key Store: ready")
	})
}

// SetGlobalKeyStore sets global Key Store
func SetGlobalKeyStore(k KeyStore) {
	keyStore = k
}

// GlobalKeyStore returns global Key Store
func GlobalKeyStore() KeyStore {
	return keyStore
}

// ImportPrivateKey create new Key Store
func ImportPrivateKey(k KeyStore) error {

	// Pre-Import Pkeys
	for _, pkey := range viper.GetStringSlice(secretPkeyViperKey) {
		err := k.ImportPrivateKey(pkey)
		if err != nil {
			return err
		}
	}

	return nil
}
