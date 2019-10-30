package main

import (
	"context"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	handler "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/tx-listener/handler/base"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/tx-listener/listener"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/types"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/engine"
)

// Handler is a engine HandlerFunc
func Handler(txctx *engine.TxContext) {
	// Cast message into sarama.ConsumerMessage
	r, ok := txctx.In.(*types.TxListenerReceipt)
	if !ok {
		panic("loader: expected a types.TxListenerReceipt")
	}

	log.Infof("* New receipt ChainID=%v BlockNumber=%v Txindex=%v TxHash=%v", r.ChainID.Text(10), r.BlockNumber, r.TxIndex, r.TxHash.Hex())
}

func main() {
	// Set log configuration
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)

	// Initialize Listener
	viper.Set("eth.client.url", []string{
		"https://ropsten.infura.io/v3/81e039ce6c8a465180822b525e3644d7",
		"https://mainnet.infura.io/v3/bfc9d6e51fbc4d3db54bea58d1094f9c",
	})

	// Initialize listener
	listener.Init(context.Background())

	// Initialize engine and register handlers
	engine.Init(context.Background())
	engine.Register(Handler)

	// Create handler
	conf, err := handler.NewConfig()
	if err != nil {
		log.WithError(err).Fatalf("listener: could not create config")
	}
	h := handler.NewHandler(engine.GlobalEngine(), conf)

	// Start listening
	_ = listener.Listen(
		context.Background(),
		[]*big.Int{
			big.NewInt(1),
			big.NewInt(3),
		},
		h,
	)
}
