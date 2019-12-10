package grpcserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
)

func TestRecoverPanicHandler(t *testing.T) {
	err := RecoverPanicHandler("test")
	assert.True(t, errors.IsInternalError(err), "Error should be an internal error")
}