package dataagents

import (
	"context"
	"fmt"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/secretstore"

	log "github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
)

const ethereumDAComponent = "data-agents.ethereum"

// HashicorpEthereum is a job data agent for Ethereum Hashicorp Vault
type HashicorpEthereum struct {
	vault       secretstore.SecretStore
	generateKey func(address, namespace string) string
}

// NewHashicorpEthereum creates a new HashicorpEthereum
func NewHashicorpEthereum(vault secretstore.SecretStore) *HashicorpEthereum {
	return &HashicorpEthereum{vault: vault, generateKey: generateKey}
}

func (agent *HashicorpEthereum) Insert(ctx context.Context, address, privKey, namespace string) error {
	key := agent.generateKey(address, namespace)
	err := agent.vault.Store(ctx, key, privKey)
	if err != nil {
		errMessage := "failed to store privateKey in Hashicorp Vault"
		log.WithContext(ctx).WithError(err).WithField("key", key).Error(errMessage)
		return errors.HashicorpVaultConnectionError(errMessage).ExtendComponent(ethereumDAComponent)
	}

	return nil
}

func (agent *HashicorpEthereum) FindOne(ctx context.Context, address, namespace string) (string, error) {
	key := agent.generateKey(address, namespace)
	logger := log.WithContext(ctx).WithField("key", key)

	privKey, ok, err := agent.vault.Load(ctx, key)
	if err != nil {
		errMessage := "failed to load privateKey from Hashicorp Vault"
		logger.WithError(err).Error(errMessage)
		return "", errors.HashicorpVaultConnectionError(errMessage).ExtendComponent(ethereumDAComponent)
	}

	if !ok {
		warnMessage := "account does not exist"
		logger.Warn(warnMessage)
		return "", errors.NotFoundError(warnMessage).ExtendComponent(ethereumDAComponent)
	}

	return privKey, nil
}

func generateKey(address, namespace string) string {
	return fmt.Sprintf("%v%v", namespace, address)
}
