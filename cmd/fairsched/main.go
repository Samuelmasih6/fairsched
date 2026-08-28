package main

import (
	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New(10)

	jobs := []job.Job{
		{
			ID:       "job-001",
			TenantID: "tenant-a",
			Priority: 1,
			Payload:  "Process file A",
		},
		{
			ID:       "job-002",
			TenantID: "tenant-a",
			Priority: 10,
			Payload:  "Process file B",
		},
		{
			ID:       "job-003",
			TenantID: "tenant-b",
			Priority: 5,
			Payload:  "Process file C",
		},
		{
			ID:       "job-004",
			TenantID: "tenant-b",
			Priority: 3,
			Payload:  "Process file D",
		},
		{
			ID:       "job-005",
			TenantID: "tenant-a",
			Priority: 8,
			Payload:  "Process file E",
		},
	}

	// Start 3 workers.
	s.Start(3)

	// Submit jobs.
	for _, j := range jobs {
		s.Submit(j)
	}

	// Close the queue and wait for workers.
	s.Close()
}
