package faucet

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/faucet/controllers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/faucet/faucet"
)

const component = "handler.faucet"

var (
	handler  engine.HandlerFunc
	initOnce = &sync.Once{}
)

// Init initialize Crafter Handler
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if handler != nil {
			return
		}

		// Initialize Controlled Faucet
		controllers.Init(ctx)

		// Create Handler
		handler = Faucet(faucet.GlobalFaucet())

		log.Infof("logger: handler ready")
	})
}

// SetGlobalHandler sets global Faucet Handler
func SetGlobalHandler(h engine.HandlerFunc) {
	handler = h
}

// GlobalHandler returns global Faucet handler
func GlobalHandler() engine.HandlerFunc {
	return handler
}
