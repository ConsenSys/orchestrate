package kafka

import (
	"context"
	"sync"

	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethereum/ethclient/rpc"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	crc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
	storeclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/envelope-store/client"
)

var (
	hook     *Hook
	initOnce = &sync.Once{}
)

func initComponent(ctx context.Context) {
	utils.InParallel(
		// Initialize Ethereum Client
		func() { rpc.Init(ctx) },
		// Initialize Contract Registry Client
		func() { crc.Init(ctx, viper.GetString(crc.ContractRegistryURLViperKey)) },
		// Initialize Envelope store client
		func() { storeclient.Init(ctx) },
		// Initialize Sync Producer
		func() { broker.InitSyncProducer(ctx) },
	)
}

// Init Kafka hook
func Init(ctx context.Context) {
	initOnce.Do(func() {
		initComponent(ctx)

		hook = NewHook(
			NewConfig(),
			crc.GlobalClient(),
			rpc.GlobalClient(),
			storeclient.GlobalEnvelopeStoreClient(),
			broker.GlobalSyncProducer(),
		)
	})
}

// SetGlobalHook set global Kafka hook
func SetGlobalHook(hk *Hook) {
	hook = hk
}

// GlobalHook return global Kafka hook
func GlobalHook() *Hook {
	return hook
}
