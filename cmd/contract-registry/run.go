package contractregistry

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	ethclient "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/ethereum/ethclient/rpc"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/http"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/tracing/opentracing/jaeger"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/pkg/utils"
	contractregistry "gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/contract-registry"
	"gitlab.com/ConsenSys/client/fr/core-stack/corestack.git/services/contract-registry/redis"
)

func newRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run application",
		Run:   run,
	}

	// Register OpenTracing flags
	jaeger.InitFlags(runCmd.Flags())

	// Register HTTP server flags
	http.Hostname(runCmd.Flags())

	// EthClient flag
	ethclient.URLs(runCmd.Flags())

	// ContractRegistry flag
	contractregistry.Type(runCmd.Flags())
	contractregistry.ABIs(runCmd.Flags())

	// Redis ContractRegistry flag
	redis.InitFlags(runCmd.Flags())

	return runCmd
}

func run(_ *cobra.Command, _ []string) {
	// Process signals
	sig := utils.NewSignalListener(func(signal os.Signal) { Close(context.Background()) })
	defer sig.Close()

	// Start application
	Start(context.Background())
}
