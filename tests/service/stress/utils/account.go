package utils

import (
	"context"

	"github.com/containous/traefik/v2/pkg/log"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/identitymanager"
	identitymanager "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/identity-manager/client"
)

type ctxKey string

var accountCtxKey ctxKey = "accounts"

func CreateNewAccount(ctx context.Context, identity identitymanager.IdentityManagerClient) (string, error) {
	log.FromContext(ctx).Debugf("Registering new account...")
	resp, err := identity.CreateAccount(ctx, &types.CreateAccountRequest{})
	if err != nil {
		return "", nil
	}

	log.FromContext(ctx).WithField("address", resp.Address).Info("Account has been registered")
	return resp.Address, nil
}

func ContextWithAccounts(ctx context.Context, accounts []string) context.Context {
	return context.WithValue(ctx, accountCtxKey, accounts)
}

func ContextAccounts(ctx context.Context) []string {
	v, ok := ctx.Value(accountCtxKey).([]string)
	if !ok {
		return []string{}
	}
	return v
}
