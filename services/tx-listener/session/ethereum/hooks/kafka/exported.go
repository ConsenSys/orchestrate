package kafka

import (
	"context"
	"sync"

	txscheduler "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/client"

	"github.com/spf13/viper"
	broker "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/broker/sarama"
	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/ethclient/rpc"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/utils"
	crc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/client"
)

var (
	hook     *Hook
	initOnce = &sync.Once{}
)

func initComponent(ctx context.Context) {
	utils.InParallel(
		// Initialize Ethereum Client
		func() { ethclient.Init(ctx) },
		// Initialize Contract Registry Client
		func() { crc.Init(ctx, viper.GetString(crc.ContractRegistryURLViperKey)) },
		// Initialize Sync Producer
		func() { broker.InitSyncProducer(ctx) },
		// Initialize transaction scheduler client
		func() { txscheduler.Init() },
	)
}

// Init Kafka hook
func Init(ctx context.Context) {
	initOnce.Do(func() {
		initComponent(ctx)

		hook = NewHook(
			NewConfig(),
			crc.GlobalClient(),
			ethclient.GlobalClient(),
			broker.GlobalSyncProducer(),
			txscheduler.GlobalClient(),
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
