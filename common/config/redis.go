package config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	redisAddressFlag     = "redis-address"
	redisAddressViperKey = "redis.address"
	redisAddressDefault  = "localhost:6379"
	redisAddressEnv      = "REDIS_ADDRESS"
)

// RedisAddress register a flag for Redis server address
func RedisAddress(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Address of Redis server to connect to.
Environment variable: %q`, redisAddressEnv)
	f.String(redisAddressFlag, redisAddressDefault, desc)
	viper.SetDefault(redisAddressViperKey, redisAddressDefault)
	viper.BindPFlag(redisAddressViperKey, f.Lookup(redisAddressFlag))
	viper.BindEnv(redisAddressViperKey, redisAddressEnv)
}

var (
	redisLockTimeoutFlag     = "redis-lock-timeout"
	redisLockTimeoutViperKey = "redis.lock.timeout"
	redisLockTimeoutDefault  = 1500
	redisLockTimeoutEnv      = "REDIS_LOCKTIMEOUT"
)

// RedisLockTimeout register a flag for Redis lock timeout
func RedisLockTimeout(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Redis lock timeout.
Environment variable: %q`, redisLockTimeoutEnv)
	f.Int(redisLockTimeoutFlag, redisLockTimeoutDefault, desc)
	viper.SetDefault(redisLockTimeoutViperKey, redisLockTimeoutDefault)
	viper.BindPFlag(redisLockTimeoutViperKey, f.Lookup(redisLockTimeoutFlag))
	viper.BindEnv(redisLockTimeoutViperKey, redisLockTimeoutEnv)
}
