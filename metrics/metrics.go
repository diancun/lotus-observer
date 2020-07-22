package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// TasksInQueue current number of tasks in the queue
	TasksInQueue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tasks_in_queue",
			Help: "Current number of tasks in the queue",
		},
		[]string{"miner", "sector_id", "task_type", "worker"},
	)

	// TasksDurationHistogram Tasks duration distribution
	TasksDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tasks_duration_seconds",
			Help:    "Tasks duration distribution",
			Buckets: []float64{60, 600, 1200, 1800, 3600, 7200},
		},
		[]string{"miner", "sector_id", "task_type", "worker", "code"},
	)
)
