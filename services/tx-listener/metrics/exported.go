package metrics

import (
	"fmt"

	pkgmetrics "github.com/consensys/orchestrate/pkg/toolkit/app/metrics"
	kitmetrics "github.com/go-kit/kit/metrics"
)

//go:generate mockgen -source=exported.go -destination=mock/mock.go -package=mock

var ModuleName = fmt.Sprintf("%s_%s", pkgmetrics.Namespace, Subsystem)

type ListenerMetrics interface {
	BlockCounter() kitmetrics.Counter
	pkgmetrics.Prometheus
}
