package scheduler

import (
	"testing"

	"github.com/Samuelmasih6/fairsched/internal/job"
)

func TestTenantQueueRoundRobin(t *testing.T) {
	tq := NewTenantQueue()

	// Tenant A has three jobs; Tenant B has two.
	for i := 1; i <= 3; i++ {
		tq.Push(job.Job{
			ID:       string(rune('0' + i)),
			TenantID: "A",
			Priority: 1,
		})
	}

	for i := 1; i <= 2; i++ {
		tq.Push(job.Job{
			ID:       string(rune('0' + i)),
			TenantID: "B",
			Priority: 1,
		})
	}

	expectedTenants := []string{"A", "B", "A", "B", "A"}

	for i, expected := range expectedTenants {
		j, err := tq.PopNext(0)
		if err != nil {
			t.Fatalf("PopNext() returned error: %v", err)
		}

		if j.TenantID != expected {
			t.Fatalf(
				"job %d: got tenant %s, want %s",
				i, j.TenantID, expected,
			)
		}
	}

	if tq.Len() != 0 {
		t.Fatalf("queue should be empty, got %d jobs", tq.Len())
	}
}

func TestTenantQueueEmpty(t *testing.T) {
	tq := NewTenantQueue()

	_, err := tq.PopNext(0)
	if err == nil {
		t.Fatal("expected error when popping from an empty queue")
	}
}

func TestTenantQueuePriorityWithinTenant(t *testing.T) {
	tq := NewTenantQueue()

	tq.Push(job.Job{
		ID:       "low",
		TenantID: "A",
		Priority: 1,
	})

	tq.Push(job.Job{
		ID:       "high",
		TenantID: "A",
		Priority: 10,
	})

	j, err := tq.PopNext(0)
	if err != nil {
		t.Fatalf("PopNext() returned error: %v", err)
	}

	if j.ID != "high" {
		t.Fatalf("got job %s, want high-priority job", j.ID)
	}
}
