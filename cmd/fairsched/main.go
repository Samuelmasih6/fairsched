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

	// Low-priority job
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

	// Many high-priority jobs
	for i := 1; i <= 20; i++ {
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
