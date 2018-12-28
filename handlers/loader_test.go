package handlers

import (
	"fmt"
	"sync"
	"testing"

	"gitlab.com/ConsenSys/client/fr/core-stack/core/infra"
	tracepb "gitlab.com/ConsenSys/client/fr/core-stack/core/protobuf/trace"
)

type MockUnmarshaller struct {
	t *testing.T
}

func (u *MockUnmarshaller) Unmarshal(msg interface{}, pb *tracepb.Trace) error {
	if msg.(string) == "error" {
		return fmt.Errorf("Could not unmarshall")
	}
	return nil
}

func makeLoaderContext(i int) *infra.Context {
	ctx := infra.NewContext()
	ctx.Reset()
	switch i % 2 {
	case 0:
		ctx.Msg = "error"
		ctx.Keys["errors"] = 1
	case 1:
		ctx.Msg = "valid"
		ctx.Keys["errors"] = 0
	}
	return ctx
}

func TestLoader(t *testing.T) {
	mu := MockUnmarshaller{t: t}
	loader := Loader(&mu)

	rounds := 10
	outs := make(chan *infra.Context, rounds)
	wg := &sync.WaitGroup{}
	for i := 0; i < rounds; i++ {
		wg.Add(1)
		ctx := makeLoaderContext(i)
		go func(ctx *infra.Context) {
			defer wg.Done()
			loader(ctx)
			outs <- ctx
		}(ctx)
	}
	wg.Wait()
	close(outs)
	if len(outs) != rounds {
		t.Errorf("Loader: expected %v outs but got %v", rounds, len(outs))
	}

	for out := range outs {
		errCount := out.Keys["errors"].(int)
		if len(out.T.Errors) != errCount {
			t.Errorf("Loader: expected %v errors but got %v", errCount, out.T.Errors)
		}
	}
}
