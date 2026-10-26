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
	TenantJobCounts  map[string]int
	FairnessIndex    float64
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
	tenantCounts := CompletedJobsByTenant(metrics)
	fairnessIndex := JainFairnessIndex(tenantCounts)

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
		TenantJobCounts:  tenantCounts,
		FairnessIndex:    fairnessIndex,
	}
}

func RunTenantFairnessBenchmark() BenchmarkResult {
	s := New(0)

	s.Start(3)

	// Give both tenants equal demand.
	for i := 1; i <= 20; i++ {
		if err := s.Submit(job.Job{
			ID:       fmt.Sprintf("tenant-a-%02d", i),
			TenantID: "tenant-a",
			Priority: 5,
			Duration: 200 * time.Millisecond,
		}); err != nil {
			panic(err)
		}

		if err := s.Submit(job.Job{
			ID:       fmt.Sprintf("tenant-b-%02d", i),
			TenantID: "tenant-b",
			Priority: 5,
			Duration: 200 * time.Millisecond,
		}); err != nil {
			panic(err)
		}
	}

	s.Shutdown()
	s.Wait()

	metrics := s.Metrics()
	summary := SummarizeMetrics(metrics)
	tenantCounts := CompletedJobsByTenant(metrics)

	return BenchmarkResult{
		AgingFactor:      0,
		JobsCompleted:    summary.JobCount,
		AverageQueueWait: summary.AverageQueueWait,
		MaximumQueueWait: summary.MaximumQueueWait,
		AverageLatency:   summary.AverageLatency,
		MaximumLatency:   summary.MaximumLatency,
		TenantJobCounts:  tenantCounts,
		FairnessIndex:    JainFairnessIndex(tenantCounts),
	}
}
