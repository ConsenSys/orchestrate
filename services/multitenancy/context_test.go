package multitenancy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTenantID(t *testing.T) {
	assert.Equal(t, "", TenantIDFromContext(context.Background()))
	ctx := WithTenantID(context.Background(), "test-tenant")
	assert.Equal(t, "test-tenant", TenantIDFromContext(ctx))
}