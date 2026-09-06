package scheduler

import (
	"time"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type PriorityQueue []job.Job

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(j job.Job) {
	*pq = append(*pq, j)
}

func (pq *PriorityQueue) PopHighestPriority(agingFactor float64) job.Job {
	now := time.Now()

	bestIndex := 0
	bestScore := effectivePriority((*pq)[0], now, agingFactor)

	for i := 1; i < len(*pq); i++ {
		score := effectivePriority((*pq)[i], now, agingFactor)

		if score > bestScore {
			bestScore = score
			bestIndex = i
		}
	}

	j := (*pq)[bestIndex]

	*pq = append((*pq)[:bestIndex], (*pq)[bestIndex+1:]...)

	return j
}

func effectivePriority(j job.Job, now time.Time, agingFactor float64) float64 {
	waitingSeconds := now.Sub(j.CreatedAt).Seconds()

	return float64(j.Priority) + agingFactor*waitingSeconds
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}
