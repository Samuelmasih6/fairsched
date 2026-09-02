package main

import (
	"fmt"

	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New()

	s.Start(3)

	jobs := []job.Job{
		{
			ID:       "job-001",
			TenantID: "tenant-a",
			Priority: 1,
			Payload:  "Low priority job",
			Duration: 5,
		},
		{
			ID:       "job-002",
			TenantID: "tenant-a",
			Priority: 10,
			Payload:  "Highest priority job",
			Duration: 1,
		},
		{
			ID:       "job-003",
			TenantID: "tenant-b",
			Priority: 5,
			Payload:  "Medium priority job",
			Duration: 3,
		},
		{
			ID:       "job-004",
			TenantID: "tenant-b",
			Priority: 3,
			Payload:  "Low-medium priority job",
			Duration: 2,
		},
		{
			ID:       "job-005",
			TenantID: "tenant-a",
			Priority: 8,
			Payload:  "High priority job",
			Duration: 7,
		},
	}

	for _, j := range jobs {
		if err := s.Submit(j); err != nil {
			fmt.Println("failed to submit job:", err)
		}
	}

	s.Shutdown()
	s.Wait()
}
