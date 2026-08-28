package scheduler

import (
	"fmt"
	"sync"
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type Scheduler struct {
	queue chan job.Job
	wg    sync.WaitGroup
}

func New(queueSize int) *Scheduler {
	return &Scheduler{
		queue: make(chan job.Job, queueSize),
	}
}

func (s *Scheduler) Submit(j job.Job) {
	j.Status = job.StatusQueued
	s.queue <- j
}

func (s *Scheduler) Start(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		s.wg.Add(1)

		go func(workerID int) {
			defer s.wg.Done()

			for j := range s.queue {
				j.Status = job.StatusRunning

				fmt.Printf(
					"Worker %d started job %s\n",
					workerID,
					j.ID,
				)

				// Simulate job execution.
				time.Sleep(2 * time.Second)

				j.Status = job.StatusCompleted

				fmt.Printf(
					"Worker %d completed job %s\n",
					workerID,
					j.ID,
				)
			}
		}(i)
	}
}

func (s *Scheduler) Close() {
	close(s.queue)
	s.wg.Wait()
}
