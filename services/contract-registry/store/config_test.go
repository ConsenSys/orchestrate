// +build unit

package store

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.NotNil(t, cfg)
}

func TestStoreType(t *testing.T) {
	flgs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	Type(flgs)

	expected := "postgres"
	assert.Equal(t, expected, viper.GetString(TypeViperKey), "Default")

	expected = "env-store"
	_ = os.Setenv(typeEnv, expected)
	assert.Equal(t, expected, viper.GetString(TypeViperKey), "From Environment Variable")
	_ = os.Unsetenv(typeEnv)

	args := []string{
		"--contract-registry-type=flag-store",
	}
	err := flgs.Parse(args)
	assert.NoError(t, err, "No error expected")

	expected = "flag-store"
	assert.Equal(t, expected, viper.GetString(TypeViperKey), "From flag")
}