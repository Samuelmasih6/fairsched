package scheduler

import "time"

type MetricsSummary struct {
	JobCount         int
	AverageQueueWait time.Duration
	MaximumQueueWait time.Duration
	AverageLatency   time.Duration
	MaximumLatency   time.Duration
}

type TenantMetrics struct {
	JobCount         int
	AverageQueueWait time.Duration
	MaximumQueueWait time.Duration
	AverageLatency   time.Duration
	MaximumLatency   time.Duration
}

func SummarizeMetrics(metrics []JobMetrics) MetricsSummary {
	if len(metrics) == 0 {
		return MetricsSummary{}
	}

	var totalQueueWait time.Duration
	var totalLatency time.Duration
	var maximumQueueWait time.Duration
	var maximumLatency time.Duration

	for _, m := range metrics {
		totalQueueWait += m.QueueWait
		totalLatency += m.TotalLatency

		if m.QueueWait > maximumQueueWait {
			maximumQueueWait = m.QueueWait
		}

		if m.TotalLatency > maximumLatency {
			maximumLatency = m.TotalLatency
		}
	}

	count := time.Duration(len(metrics))

	return MetricsSummary{
		JobCount:         len(metrics),
		AverageQueueWait: totalQueueWait / count,
		MaximumQueueWait: maximumQueueWait,
		AverageLatency:   totalLatency / count,
		MaximumLatency:   maximumLatency,
	}
}

func SummarizeMetricsByTenant(metrics []JobMetrics) map[string]TenantMetrics {
	totals := make(map[string]TenantMetrics)
	queueWaitTotals := make(map[string]time.Duration)
	latencyTotals := make(map[string]time.Duration)

	for _, m := range metrics {
		current := totals[m.TenantID]
		current.JobCount++

		if m.QueueWait > current.MaximumQueueWait {
			current.MaximumQueueWait = m.QueueWait
		}

		if m.TotalLatency > current.MaximumLatency {
			current.MaximumLatency = m.TotalLatency
		}

		totals[m.TenantID] = current
		queueWaitTotals[m.TenantID] += m.QueueWait
		latencyTotals[m.TenantID] += m.TotalLatency
	}

	for tenantID, summary := range totals {
		count := time.Duration(summary.JobCount)
		summary.AverageQueueWait = queueWaitTotals[tenantID] / count
		summary.AverageLatency = latencyTotals[tenantID] / count
		totals[tenantID] = summary
	}

	return totals
}
