package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	TaskType_AddPiece         = "Add_Piece"
	TaskType_PreCommit1       = "PreCommit1"
	TaskType_PreCommit1_Queue = "PreCommit1_Queue"
	TaskType_Fetch            = "Fetch"
	TaskType_PreCommit2_Queue = "PreCommit2_Queue"
	TaskType_PreCommit2       = "PreCommit2"
	TaskType_Commit1_Queue    = "Commit1_Queue"
	TaskType_Commit1          = "Commit1"
	TaskType_Commit2_Queue    = "Commit2_Queue"
	TaskType_Commit2          = "Commit2"
	TaskType_Finalize_Queue   = "Finalize_Queue"
	TaskType_Finalize         = "Finalize"
	TaskType_ReadUnsealed     = "Read_Unsealed"
	TaskType_Unseal           = "Unseal"
)

var (
	// TasksInQueue current number of tasks in the queue
	TasksInQueue = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tasks_in_queue",
			Help: "Current number of tasks in the queue",
		},
		[]string{"miner", "sector_id", "task_type"},
	)

	// TasksInQueueHistogram tasks in queue distribution
	TasksInQueueHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tasks_in_queue_seconds",
			Help:    "Tasks in queue distribution",
			Buckets: []float64{60, 600, 1200, 1800, 3600},
		},
		[]string{"miner", "sector_id", "task_type"},
	)
)

var (
	// TasksInProgress current number of tasks in progress
	TasksInProgress = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tasks_in_progress",
			Help: "Current number of tasks in progress",
		},
		[]string{"miner", "sector_id", "task_type", "worker"},
	)

	// TasksInProgressHistogram tasks in progress distribution
	TasksInProgressHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tasks_in_progress_seconds",
			Help:    "Tasks in progress distribution",
			Buckets: []float64{60, 600, 1200, 1800, 3600, 7200},
		},
		[]string{"miner", "sector_id", "task_type", "worker", "code"},
	)
)
