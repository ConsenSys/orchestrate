package dialer

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/grpc"
	svc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/contract-registry/proto"
)

func DialContextWithDefaultOptions(ctx context.Context, url string) (svc.ContractRegistryClient, error) {
	conn, err := grpc.DialContextWithDefaultOptions(
		ctx,
		url,
	)
	if err != nil {
		return nil, err
	}

	return svc.NewContractRegistryClient(conn), nil
}
