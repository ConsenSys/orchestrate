package enricher

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethclient/rpc"
	crc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
)

const (
	component = "handler.listener.enricher"
)

var (
	handler  engine.HandlerFunc
	initOnce = &sync.Once{}
)

// Init initialize handler
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if handler != nil {
			return
		}

		// Initialize Ethereum Client
		ethclient.Init(ctx)

		// Initialize Contract Registry Client
		crc.Init(ctx, viper.GetString(crc.ContractRegistryURLViperKey))

		// Create Handler
		handler = Enricher(crc.GlobalClient(), ethclient.GlobalClient())

		log.Infof("enricher: handler ready")
	})
}

// SetGlobalHandler sets global handler
func SetGlobalHandler(h engine.HandlerFunc) {
	handler = h
}

// GlobalHandler returns global handler
func GlobalHandler() engine.HandlerFunc {
	return handler
}
