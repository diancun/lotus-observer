package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Workers current number of workers
	Workers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "workers",
			Help: "current number of workers",
		},
	)
)
