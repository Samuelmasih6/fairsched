package main

import (
	"fmt"
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New(1.0)

	s.Start(3)

	// Occupy all workers.
	for i := 1; i <= 3; i++ {
		j := job.Job{
			ID:       fmt.Sprintf("initial-high-%02d", i),
			TenantID: "tenant-b",
			Priority: 10,
			Payload:  "Initial high priority job",
			Duration: 3 * time.Second,
		}

		if err := s.Submit(j); err != nil {
			fmt.Println("failed to submit job:", err)
		}
	}

	// Give the workers time to start.
	time.Sleep(100 * time.Millisecond)

	// Low-priority job enters while all workers are busy.
	lowPriorityJob := job.Job{
		ID:       "low-priority",
		TenantID: "tenant-a",
		Priority: 1,
		Payload:  "Low priority job",
		Duration: 1 * time.Second,
	}

	if err := s.Submit(lowPriorityJob); err != nil {
		fmt.Println("failed to submit job:", err)
	}

	// Continuously submit high-priority jobs.
	for i := 1; i <= 40; i++ {
		j := job.Job{
			ID:       fmt.Sprintf("high-priority-%02d", i),
			TenantID: "tenant-b",
			Priority: 10,
			Payload:  "High priority job",
			Duration: 1 * time.Second,
		}

		if err := s.Submit(j); err != nil {
			fmt.Println("failed to submit job:", err)
		}

		// New high-priority jobs arrive over time.
		time.Sleep(300 * time.Millisecond)
	}

	s.Shutdown()
	s.Wait()

	metrics := s.Metrics()
	summary := scheduler.SummarizeMetrics(metrics)

	fmt.Println()
	fmt.Println("=== Scheduling Summary ===")
	fmt.Printf("Jobs completed: %d\n", summary.JobCount)
	fmt.Printf("Average queue wait: %v\n", summary.AverageQueueWait)
	fmt.Printf("Maximum queue wait: %v\n", summary.MaximumQueueWait)
	fmt.Printf("Average latency: %v\n", summary.AverageLatency)
	fmt.Printf("Maximum latency: %v\n", summary.MaximumLatency)
}
