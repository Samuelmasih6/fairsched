package scheduler

import (
	"fmt"
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type BenchmarkResult struct {
	AgingFactor      float64
	JobsCompleted    int
	AverageQueueWait time.Duration
	MaximumQueueWait time.Duration
	LowPriorityWait  time.Duration
	AverageLatency   time.Duration
	MaximumLatency   time.Duration
}

func RunBenchmark(agingFactor float64) BenchmarkResult {
	s := New(agingFactor)

	s.Start(3)

	// Occupy all workers first.
	for i := 1; i <= 3; i++ {
		s.Submit(job.Job{
			ID:       fmt.Sprintf("initial-%02d", i),
			TenantID: "tenant-b",
			Priority: 10,
			Duration: 3 * time.Second,
		})
	}

	time.Sleep(100 * time.Millisecond)

	// The low-priority job enters while workers are busy.
	s.Submit(job.Job{
		ID:       "low-priority",
		TenantID: "tenant-a",
		Priority: 1,
		Duration: 1 * time.Second,
	})

	// High-priority jobs continue arriving.
	for i := 1; i <= 40; i++ {
		s.Submit(job.Job{
			ID:       fmt.Sprintf("high-%02d", i),
			TenantID: "tenant-b",
			Priority: 10,
			Duration: 1 * time.Second,
		})

		time.Sleep(300 * time.Millisecond)
	}

	s.Shutdown()
	s.Wait()

	metrics := s.Metrics()
	summary := SummarizeMetrics(metrics)

	var lowPriorityWait time.Duration

	for _, m := range metrics {
		if m.JobID == "low-priority" {
			lowPriorityWait = m.QueueWait
			break
		}
	}

	return BenchmarkResult{
		AgingFactor:      agingFactor,
		JobsCompleted:    summary.JobCount,
		AverageQueueWait: summary.AverageQueueWait,
		MaximumQueueWait: summary.MaximumQueueWait,
		LowPriorityWait:  lowPriorityWait,
		AverageLatency:   summary.AverageLatency,
		MaximumLatency:   summary.MaximumLatency,
	}
}
