// +build unit

package dynamic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/http/config/dynamic"
)

func TestBuilder(t *testing.T) {
	builder := NewBuilder()
	srv := &dynamic.Service{
		HealthCheck: &dynamic.HealthCheck{},
	}

	h, err := builder.Build(context.Background(), "", srv, nil)
	require.NoError(t, err)
	assert.NotNil(t, h)
}
