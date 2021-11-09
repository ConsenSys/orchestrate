package jose

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	_ = viper.BindEnv(issuerURLViperKey, issuerURLEnv)
	_ = viper.BindEnv(audienceViperKey, audienceEnv)
}

func Flags(f *pflag.FlagSet) {
	issuerURLFlags(f)
	audienceFlags(f)
}

const (
	issuerURLFlag     = "auth-jwt-issuer-url"
	issuerURLViperKey = "auth.jwt.issuer-url"
	issuerURLEnv      = "AUTH_JWT_ISSUER_URL"
)

const (
	audienceFlag     = "auth-jwt-audience"
	audienceViperKey = "auth.jwt.audience"
	audienceEnv      = "AUTH_JWT_AUDIENCE"
)

func issuerURLFlags(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`JWT issuer server domain (ie. https://orchestrate.eu.auth0.com).
Environment variable: %q`, issuerURLEnv)
	f.String(issuerURLFlag, "", desc)
	_ = viper.BindPFlag(issuerURLViperKey, f.Lookup(issuerURLFlag))
}

func audienceFlags(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Expected audience ("aud" field) of JWT tokens.
Environment variable: %q`, audienceEnv)
	f.StringSlice(audienceFlag, []string{}, desc)
	_ = viper.BindPFlag(audienceViperKey, f.Lookup(audienceFlag))
}

type Config struct {
	IssuerURL string
	CacheTTL  time.Duration
	Audience  []string
}

func NewConfig(vipr *viper.Viper) *Config {
	issuerURL := vipr.GetString(issuerURLViperKey)

	if issuerURL != "" {
		return &Config{
			IssuerURL: issuerURL,
			Audience:  vipr.GetStringSlice(audienceViperKey),
			CacheTTL:  5 * time.Minute, // TODO: Make the cache ttl an ENV var if needed
		}
	}

	return nil
}
