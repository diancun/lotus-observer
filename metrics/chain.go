package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Workers current number of workers
	ChainHeight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "chain_height",
			Help: "current blockchain height of ",
		},
	)
)
