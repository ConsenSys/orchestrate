package handlers

import (
	"context"
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/handlers/opentracing/jaeger"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-nonce.git/handlers/nonce"
	"gitlab.com/ConsenSys/client/fr/core-stack/worker/tx-nonce.git/handlers/producer"
)

// Init inialize handlers
func Init(ctx context.Context) {
	wg := sync.WaitGroup{}

	// Initialize Jaeger tracer
	wg.Add(1)
	go func() {
		jaeger.Init(ctx)
		wg.Done()
	}()

	// Initialize Nonce manager
	wg.Add(1)
	go func() {
		nonce.Init(ctx)
		wg.Done()
	}()

	// Initialize PrepareMsg
	wg.Add(1)
	go func() {
		producer.Init(ctx)
		wg.Done()
	}()

	// Wait for all handlers to be ready
	wg.Wait()
}
