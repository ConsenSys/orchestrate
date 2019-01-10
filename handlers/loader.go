package handlers

import (
	"sync"

	"gitlab.com/ConsenSys/client/fr/core-stack/core.git/protobuf"
	tracepb "gitlab.com/ConsenSys/client/fr/core-stack/core.git/protobuf/trace"
	"gitlab.com/ConsenSys/client/fr/core-stack/core.git/services"
	"gitlab.com/ConsenSys/client/fr/core-stack/core.git/types"
)

// Loader creates an handler loading input
func Loader(u services.Unmarshaller) types.HandlerFunc {
	pool := &sync.Pool{
		New: func() interface{} { return &tracepb.Trace{} },
	}

	return func(ctx *types.Context) {
		pb := pool.Get().(*tracepb.Trace)
		defer pool.Put(pb)

		pb.Reset()

		// Unmarshal message
		err := u.Unmarshal(ctx.Msg, pb)
		if err != nil {
			// TODO: handle error
			ctx.AbortWithError(err)
			return
		}

		// Load Trace from protobuffer
		protobuf.LoadTrace(pb, ctx.T)
	}
}
