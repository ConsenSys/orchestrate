package http

import (
	"fmt"

	traefikstatic "github.com/containous/traefik/v2/pkg/config/static"
	traefiktypes "github.com/containous/traefik/v2/pkg/types"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/logger"
)

func init() {
	viper.SetDefault(hostnameViperKey, hostnameDefault)
	_ = viper.BindEnv(hostnameViperKey, hostnameEnv)

	viper.SetDefault(httpPortViperKey, httpPortDefault)
	_ = viper.BindEnv(httpPortViperKey, httpPortEnv)

	viper.SetDefault(metricsHostnameViperKey, metricsHostnameDefault)
	_ = viper.BindEnv(metricsHostnameViperKey, metricsHostnameEnv)

	viper.SetDefault(metricsPortViperKey, metricsPortDefault)
	_ = viper.BindEnv(metricsPortViperKey, metricsPortEnv)
}

const (
	hostnameFlag     = "rest-hostname"
	hostnameViperKey = "rest.hostname"
	hostnameDefault  = ""
	hostnameEnv      = "REST_HOSTNAME"
)

// Hostname register a flag for HTTP server address
func hostname(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Hostname to expose REST services
Environment variable: %q`, hostnameEnv)
	f.String(hostnameFlag, hostnameDefault, desc)
	_ = viper.BindPFlag(hostnameViperKey, f.Lookup(hostnameFlag))
}

const (
	httpPortFlag     = "rest-port"
	httpPortViperKey = "rest.port"
	httpPortDefault  = 8081
	httpPortEnv      = "REST_PORT"
)

// Port register a flag for HTTp server port
func port(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Port to expose REST services
Environment variable: %q`, httpPortEnv)
	f.Uint(httpPortFlag, httpPortDefault, desc)
	_ = viper.BindPFlag(httpPortViperKey, f.Lookup(httpPortFlag))
}

const (
	metricsHostnameFlag     = "metrics-hostname"
	metricsHostnameViperKey = "metrics.hostname"
	metricsHostnameDefault  = ""
	metricsHostnameEnv      = "METRICS_HOSTNAME"
)

// metricsHostname register a flag for metrics server hostname
func metricsHostname(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Hostname to expose metrics services
Environment variable: %q`, metricsHostnameEnv)
	f.String(metricsHostnameFlag, metricsHostnameDefault, desc)
	_ = viper.BindPFlag(metricsHostnameViperKey, f.Lookup(metricsHostnameFlag))
}

const (
	metricsPortFlag     = "metrics-port"
	metricsPortViperKey = "metrics.port"
	metricsPortDefault  = 8082
	metricsPortEnv      = "METRICS_PORT"
)

// MetricsPort register a flag for metrics server port
func metricsPort(f *pflag.FlagSet) {
	desc := fmt.Sprintf(`Port to expose metrics services
Environment variable: %q`, metricsPortEnv)
	f.Uint(metricsPortFlag, metricsPortDefault, desc)
	_ = viper.BindPFlag(metricsPortViperKey, f.Lookup(metricsPortFlag))
}

func Flags(f *pflag.FlagSet) {
	hostname(f)
	port(f)
	MetricFlags(f)
}

func MetricFlags(f *pflag.FlagSet) {
	metricsHostname(f)
	metricsPort(f)
}

func url(hostname string, port uint) string {
	return fmt.Sprintf("%s:%d", hostname, port)
}

func NewEPsConfig(vipr *viper.Viper) traefikstatic.EntryPoints {
	httpEp := &traefikstatic.EntryPoint{
		Address: url(vipr.GetString(hostnameViperKey), vipr.GetUint(httpPortViperKey)),
	}
	httpEp.SetDefaults()

	metricsEp := &traefikstatic.EntryPoint{
		Address: url(vipr.GetString(metricsHostnameViperKey), vipr.GetUint(metricsPortViperKey)),
	}
	metricsEp.SetDefaults()

	return traefikstatic.EntryPoints{
		DefaultHTTPEntryPoint:    httpEp,
		DefaultMetricsEntryPoint: metricsEp,
	}
}

func NewConfig(vipr *viper.Viper) *traefikstatic.Configuration {
	return &traefikstatic.Configuration{
		EntryPoints: NewEPsConfig(vipr),
		Metrics: &traefiktypes.Metrics{
			Prometheus: &traefiktypes.Prometheus{
				EntryPoint:           DefaultMetricsEntryPoint,
				Buckets:              []float64{0.1, 0.3, 1.2, 5},
				AddEntryPointsLabels: true,
				AddServicesLabels:    true,
			},
		},
		API: &traefikstatic.API{},
		ServersTransport: &traefikstatic.ServersTransport{
			MaxIdleConnsPerHost: 200,
			InsecureSkipVerify:  true,
		},
		Log: &traefiktypes.TraefikLog{
			Level:  vipr.GetString(logger.LogLevelViperKey),
			Format: viperToTraefikLogFormat(vipr.GetString(logger.LogFormatViperKey)),
		},
		AccessLog: &traefiktypes.AccessLog{
			Filters: &traefiktypes.AccessLogFilters{
				StatusCodes: []string{"100-199", "400-428", "430-599"},
			},
			Format: viperToTraefikLogFormat(vipr.GetString(logger.LogFormatViperKey)),
		},
	}
}

func DefaultConfig() *traefikstatic.Configuration {
	return NewConfig(viper.New())
}

func viperToTraefikLogFormat(format string) string {
	switch format {
	case "text":
		return "common"
	case "json":
		return "json"
	default:
		return "json"
	}
}
