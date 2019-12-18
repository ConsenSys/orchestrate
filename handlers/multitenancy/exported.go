package multitenancy

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/authentication/token"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/engine"
)

const component = "handler.multitenancy"

var (
	handler  engine.HandlerFunc
	initOnce = &sync.Once{}
)

// Init initialize Multi Tenancy Handler
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if handler != nil {
			return
		}

		// Initialize Authentication Manager
		token.Init(ctx)

		// Create Handler
		handler = ExtractTenant(token.GlobalAuth())

		log.Infof("authentication multi-tenancy: handler ready")
	})
}

// SetGlobalHandler sets global Gas Estimator Handler
func SetGlobalHandler(h engine.HandlerFunc) {
	handler = h
}

// GlobalHandler returns global Gas Estimator handler
func GlobalHandler() engine.HandlerFunc {
	return handler
}