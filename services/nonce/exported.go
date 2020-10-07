package nonce

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database/redis"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/nonce/memory"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	m        Manager
	initOnce = &sync.Once{}
)

const (
	redisOpt    = "redis"
	inMemoryOpt = "in-memory"
)

// Init initializes Nonce
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if m != nil {
			return
		}

		switch viper.GetString(typeViperKey) {
		case redisOpt:
			// Initialize Redis Nonce Manager
			redis.Init()

			// Set Nonce
			m = redis.GlobalNonceManager()
		case inMemoryOpt:
			// Initialize Mock Nonce
			memory.Init(ctx)

			// Set Nonce
			m = memory.GlobalNonceManager()
		default:
			log.WithFields(log.Fields{
				"type": viper.GetString(typeViperKey),
			}).Fatalf("nonce: unknown storage type")
		}
	})
}

// GlobalManager returns globalNonceManager
func GlobalManager() Manager {
	return m
}

// SetGlobalManager sets global Manager
func SetGlobalManager(mngr Manager) {
	m = mngr
}
