package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	apiExecTimeName            = "api_execution_time_per_block"
	apiExecTimeProjectionLabel = "api"
)

var (
	apiExecTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: apiExecTimeName,
		},
		[]string{
			apiExecTimeProjectionLabel,
		},
	)
)

func RecordApiExecTime(apiName string, timeInMilliseconds int64) {
	apiExecTime.With(
		prometheus.Labels{
			apiExecTimeProjectionLabel: apiName,
		},
	).Observe(float64(timeInMilliseconds))
}
