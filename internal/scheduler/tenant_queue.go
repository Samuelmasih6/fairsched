package scheduler

import (
	"fmt"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

type TenantQueue struct {
	queues      map[string]*PriorityQueue
	tenantOrder []string
	nextTenant  int
}

func NewTenantQueue() *TenantQueue {
	return &TenantQueue{
		queues: make(map[string]*PriorityQueue),
	}
}

// Push adds a job to its tenant's queue.
func (tq *TenantQueue) Push(j job.Job) {
	tenantID := j.TenantID

	if _, exists := tq.queues[tenantID]; !exists {
		tq.queues[tenantID] = NewPriorityQueue()
		tq.tenantOrder = append(tq.tenantOrder, tenantID)
	}

	tq.queues[tenantID].Push(j)
}

// PopNext selects a tenant in round-robin order,
// then selects that tenant's highest effective-priority job.
func (tq *TenantQueue) PopNext(agingFactor float64) (job.Job, error) {
	if len(tq.tenantOrder) == 0 {
		return job.Job{}, fmt.Errorf("no jobs available")
	}

	tenantCount := len(tq.tenantOrder)

	for checked := 0; checked < tenantCount; checked++ {
		if tq.nextTenant >= len(tq.tenantOrder) {
			tq.nextTenant = 0
		}

		tenantID := tq.tenantOrder[tq.nextTenant]
		tq.nextTenant = (tq.nextTenant + 1) % len(tq.tenantOrder)

		queue := tq.queues[tenantID]

		if queue.Len() == 0 {
			continue
		}

		j := queue.PopHighestPriority(agingFactor)

		if queue.Len() == 0 {
			tq.removeTenant(tenantID)
		}

		return j, nil
	}

	return job.Job{}, fmt.Errorf("no jobs available")
}

func (tq *TenantQueue) Len() int {
	total := 0

	for _, queue := range tq.queues {
		total += queue.Len()
	}

	return total
}

func (tq *TenantQueue) removeTenant(tenantID string) {
	delete(tq.queues, tenantID)

	for i, id := range tq.tenantOrder {
		if id == tenantID {
			tq.tenantOrder = append(
				tq.tenantOrder[:i],
				tq.tenantOrder[i+1:]...,
			)

			if len(tq.tenantOrder) == 0 {
				tq.nextTenant = 0
			} else if i < tq.nextTenant {
				tq.nextTenant--
			}

			if len(tq.tenantOrder) > 0 {
				tq.nextTenant %= len(tq.tenantOrder)
			}

			return
		}
	}
}
