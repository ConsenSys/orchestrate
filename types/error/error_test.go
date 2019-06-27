package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {

	err := NewError("Test Error").
		SetCode([]byte{0xab}).
		SetComponent("test-component")

	assert.Equal(t, "Test Error", err.Error(), "Error message should be valid")
	assert.Equal(t, []byte{0xab}, err.GetCode(), "Codee should be valid")
	assert.Equal(t, "test-component", err.GetComponent(), "Component should be valid")
}
