package chainregistry

import (
	"context"
	"sync"

	chainregistry "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/chain-registry/client"
)

var (
	provider *Provider
	initOnce = &sync.Once{}
)

// Init hook
func Init(ctx context.Context) {
	initOnce.Do(func() {
		if provider != nil {
			return
		}

		chainregistry.Init(ctx)

		provider = &Provider{
			Client: chainregistry.GlobalClient(),
			conf:   NewConfig(),
		}
	})
}

// SetGlobalProvider sets global a chain-registry provider
func SetGlobalProvider(p *Provider) {
	provider = p
}

// GlobalSetProvider returns global a chain-registry provider
func GlobalProvider() *Provider {
	return provider
}
