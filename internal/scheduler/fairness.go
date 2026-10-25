package scheduler

// JainFairnessIndex measures how evenly completed jobs
// are distributed among tenants. A value of 1 means
// perfectly equal distribution.
func JainFairnessIndex(completedByTenant map[string]int) float64 {
	if len(completedByTenant) == 0 {
		return 1.0
	}

	var total float64
	var sumSquares float64

	for _, count := range completedByTenant {
		if count < 0 {
			return 0
		}

		x := float64(count)
		total += x
		sumSquares += x * x
	}

	if sumSquares == 0 {
		return 1.0
	}

	n := float64(len(completedByTenant))
	return (total * total) / (n * sumSquares)
}

func CompletedJobsByTenant(metrics []JobMetrics) map[string]int {
	counts := make(map[string]int)

	for _, metric := range metrics {
		counts[metric.TenantID]++
	}

	return counts
}
