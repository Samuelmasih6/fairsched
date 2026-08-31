package main

import (
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New()

	s.Start(1)

	s.Submit(job.Job{
		ID:       "job-001",
		TenantID: "tenant-a",
		Priority: 1,
		Payload:  "Low priority job",
		Duration: 5 * time.Second,
	})

	s.Submit(job.Job{
		ID:       "job-002",
		TenantID: "tenant-a",
		Priority: 10,
		Payload:  "Highest priority job",
		Duration: 1 * time.Second,
	})

	s.Submit(job.Job{
		ID:       "job-003",
		TenantID: "tenant-b",
		Priority: 5,
		Payload:  "Medium priority job",
		Duration: 3 * time.Second,
	})

	s.Submit(job.Job{
		ID:       "job-004",
		TenantID: "tenant-b",
		Priority: 3,
		Payload:  "Low-medium priority job",
		Duration: 2 * time.Second,
	})

	s.Submit(job.Job{
		ID:       "job-005",
		TenantID: "tenant-a",
		Priority: 8,
		Payload:  "High priority job",
		Duration: 7 * time.Second,
	})

	s.Close()
}
