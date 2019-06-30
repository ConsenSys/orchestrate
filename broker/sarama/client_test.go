package sarama

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	errors "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/errors"
	err "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
)

func TestNewClient(t *testing.T) {
	_, e := NewClient([]string{"unkown"}, sarama.NewConfig())
	assert.NotNil(t, e, "Client should error")
	ie, ok := e.(*err.Error)
	assert.True(t, ok, "Error should cast to internal error")
	assert.Equal(t, "broker.sarama", ie.GetComponent(), "Component should be correct")
	assert.True(t, errors.IsConnectionError(ie), "Error should be a connection error")
}
