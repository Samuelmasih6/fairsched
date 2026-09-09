package scheduler

import "time"

type MetricsSummary struct {
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

	return MetricsSummary{
		JobCount:         len(metrics),
		AverageQueueWait: totalQueueWait / time.Duration(len(metrics)),
		MaximumQueueWait: maximumQueueWait,
		AverageLatency:   totalLatency / time.Duration(len(metrics)),
		MaximumLatency:   maximumLatency,
	}
}
