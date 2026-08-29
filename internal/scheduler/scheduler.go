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

				for s.queue.Len() == 0 {
					s.cond.Wait()
				}

				j := heap.Pop(s.queue).(job.Job)

				s.mu.Unlock()

				j.Status = job.StatusRunning

				fmt.Printf(
					"Worker %d started job %s (priority=%d)\n",
					workerID,
					j.ID,
					j.Priority,
				)

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
	s.wg.Wait()
}
