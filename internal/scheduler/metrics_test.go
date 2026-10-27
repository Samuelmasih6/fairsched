package scheduler

import (
	"testing"
	"time"
)

func TestSummarizeMetricsByTenant(t *testing.T) {
	metrics := []JobMetrics{
		{
			JobID:        "a-1",
			TenantID:     "tenant-a",
			QueueWait:    2 * time.Second,
			TotalLatency: 5 * time.Second,
		},
		{
			JobID:        "a-2",
			TenantID:     "tenant-a",
			QueueWait:    4 * time.Second,
			TotalLatency: 8 * time.Second,
		},
		{
			JobID:        "b-1",
			TenantID:     "tenant-b",
			QueueWait:    1 * time.Second,
			TotalLatency: 3 * time.Second,
		},
	}

	got := SummarizeMetricsByTenant(metrics)

	if len(got) != 2 {
		t.Fatalf("expected 2 tenants, got %d", len(got))
	}

	a := got["tenant-a"]
	if a.JobCount != 2 {
		t.Errorf("tenant-a JobCount = %d, want 2", a.JobCount)
	}
	if a.AverageQueueWait != 3*time.Second {
		t.Errorf("tenant-a AverageQueueWait = %v, want 3s", a.AverageQueueWait)
	}
	if a.MaximumQueueWait != 4*time.Second {
		t.Errorf("tenant-a MaximumQueueWait = %v, want 4s", a.MaximumQueueWait)
	}
	if a.AverageLatency != 6500*time.Millisecond {
		t.Errorf("tenant-a AverageLatency = %v, want 6.5s", a.AverageLatency)
	}
	if a.MaximumLatency != 8*time.Second {
		t.Errorf("tenant-a MaximumLatency = %v, want 8s", a.MaximumLatency)
	}

	b := got["tenant-b"]
	if b.JobCount != 1 {
		t.Errorf("tenant-b JobCount = %d, want 1", b.JobCount)
	}
	if b.AverageQueueWait != time.Second {
		t.Errorf("tenant-b AverageQueueWait = %v, want 1s", b.AverageQueueWait)
	}
	if b.MaximumLatency != 3*time.Second {
		t.Errorf("tenant-b MaximumLatency = %v, want 3s", b.MaximumLatency)
	}
}

func TestSummarizeMetricsByTenantEmpty(t *testing.T) {
	got := SummarizeMetricsByTenant(nil)

	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
