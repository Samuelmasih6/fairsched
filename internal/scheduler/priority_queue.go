package scheduler

import (
	"container/heap"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type PriorityQueue []job.Job

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(job.Job))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	j := old[n-1]
	*pq = old[:n-1]

	return j
}

func (pq PriorityQueue) Peek() job.Job {
	return pq[0]
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}

	heap.Init(pq)

	return pq
}
