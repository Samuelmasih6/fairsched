package main

import (
	"fmt"
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New()

	s.Start(3)

	// Occupy all workers first.
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

	// Give workers time to start the initial jobs.
	time.Sleep(100 * time.Millisecond)

	// This job enters while all workers are busy.
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

	// Keep adding high-priority jobs.
	for i := 1; i <= 60; i++ {
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
	}

	s.Shutdown()
	s.Wait()
}
