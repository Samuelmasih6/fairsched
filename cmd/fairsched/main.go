package main

import (
	"fmt"

	"github.com/Samuelmasih6/fairsched/internal/job"
	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	s := scheduler.New()

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
	}

	for _, j := range jobs {
		s.Submit(j)
	}

	for {
		j, ok := s.Next()

		if !ok {
			break
		}

		scheduler.Execute(j)

		fmt.Printf(
			"Job %s status: %s\n",
			j.ID,
			j.Status,
		)
	}
}
