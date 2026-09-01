package scheduler

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type Scheduler struct {
	queue *PriorityQueue

	mu   sync.Mutex
	cond *sync.Cond

	wg sync.WaitGroup

	stopping bool
}

func New() *Scheduler {
	s := &Scheduler{
		queue: NewPriorityQueue(),
	}

	s.cond = sync.NewCond(&s.mu)

	return s
}

func (s *Scheduler) Submit(j job.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.stopping {
		return
	}

	j.CreatedAt = time.Now()
	j.Status = job.StatusQueued

	heap.Push(s.queue, j)

	s.cond.Signal()
}

func (s *Scheduler) Start(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		s.wg.Add(1)

		go func(workerID int) {
			defer s.wg.Done()

			for {
				s.mu.Lock()

				for s.queue.Len() == 0 && !s.stopping {
					s.cond.Wait()
				}

				if s.queue.Len() == 0 && s.stopping {
					s.mu.Unlock()
					return
				}

				j := heap.Pop(s.queue).(job.Job)

				s.mu.Unlock()

				j.StartedAt = time.Now()
				j.Status = job.StatusRunning

				fmt.Printf(
					"Worker %d started job %s (priority=%d, duration=%v)\n",
					workerID,
					j.ID,
					j.Priority,
					j.Duration,
				)

				time.Sleep(j.Duration)

				j.CompletedAt = time.Now()
				j.Status = job.StatusCompleted

				queueWait := j.StartedAt.Sub(j.CreatedAt)
				executionTime := j.CompletedAt.Sub(j.StartedAt)
				totalLatency := j.CompletedAt.Sub(j.CreatedAt)

				fmt.Printf(
					"Worker %d completed job %s | queue_wait=%v execution=%v latency=%v\n",
					workerID,
					j.ID,
					queueWait,
					executionTime,
					totalLatency,
				)
			}
		}(i)
	}
}

func (s *Scheduler) Shutdown() {
	s.mu.Lock()

	s.stopping = true

	s.cond.Broadcast()

	s.mu.Unlock()
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}
