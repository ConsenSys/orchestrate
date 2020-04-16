// +build unit

package transactionscheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	runCmd := newRunCommand()
	assert.NotNil(t, runCmd, "run cmd should not be nil")
}