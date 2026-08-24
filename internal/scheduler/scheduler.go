package scheduler

import (
	"fmt"
	"sync"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type Scheduler struct {
	queue []job.Job

	mu sync.Mutex
}

func New() *Scheduler {
	return &Scheduler{
		queue: make([]job.Job, 0),
	}
}

func (s *Scheduler) Submit(j job.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	j.Status = job.StatusQueued
	s.queue = append(s.queue, j)
}

func (s *Scheduler) Next() (job.Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.queue) == 0 {
		return job.Job{}, false
	}

	j := s.queue[0]
	s.queue = s.queue[1:]

	j.Status = job.StatusRunning

	return j, true
}

func Execute(j job.Job) {
	fmt.Printf("Executing job %s: %s\n", j.ID, j.Payload)
}
