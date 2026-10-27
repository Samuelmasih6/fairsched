package main

import (
	"fmt"

	"github.com/Samuelmasih6/fairsched/internal/scheduler"
)

func main() {
	fmt.Println("Running priority-only benchmark...")
	priorityResult := scheduler.RunBenchmark(0)

	fmt.Println()
	fmt.Println("Running priority + aging benchmark...")
	agingResult := scheduler.RunBenchmark(1.0)

	fmt.Println()
	fmt.Println("=== Benchmark Results ===")

	fmt.Println("\nPriority Only")
	printResult(priorityResult)

	fmt.Println("\nPriority + Aging")
	printResult(agingResult)

	fmt.Println("\n=== Comparison ===")
	fmt.Printf("Low-priority queue wait: %v → %v\n",
		priorityResult.LowPriorityWait, agingResult.LowPriorityWait)
	fmt.Printf("Average queue wait: %v → %v\n",
		priorityResult.AverageQueueWait, agingResult.AverageQueueWait)

	fmt.Println("\n=== Tenant Fairness Benchmark ===")
	fairnessResult := scheduler.RunTenantFairnessBenchmark()
	fmt.Printf("Jobs completed: %d\n", fairnessResult.JobsCompleted)
	fmt.Printf("Tenant job counts: %v\n", fairnessResult.TenantJobCounts)
	fmt.Printf("Jain fairness index: %.4f\n", fairnessResult.FairnessIndex)
	fmt.Printf("Average queue wait: %v\n", fairnessResult.AverageQueueWait)
	fmt.Printf("Maximum queue wait: %v\n", fairnessResult.MaximumQueueWait)

	fmt.Println()
	fmt.Println("=== Unequal Tenant Demand Benchmark ===")

	unequalResult := scheduler.RunUnequalTenantBenchmark()

	fmt.Printf("Jobs completed: %d\n", unequalResult.JobsCompleted)
	fmt.Printf("Tenant job counts: %v\n", unequalResult.TenantJobCounts)
	fmt.Printf("Jain fairness index: %.4f\n", unequalResult.FairnessIndex)
	fmt.Printf("Average queue wait: %v\n", unequalResult.AverageQueueWait)
	fmt.Printf("Maximum queue wait: %v\n", unequalResult.MaximumQueueWait)
}

func printResult(result scheduler.BenchmarkResult) {
	fmt.Printf("Aging factor: %v\n", result.AgingFactor)
	fmt.Printf("Jobs completed: %d\n", result.JobsCompleted)
	fmt.Printf("Low-priority wait: %v\n", result.LowPriorityWait)
	fmt.Printf("Average queue wait: %v\n", result.AverageQueueWait)
	fmt.Printf("Maximum queue wait: %v\n", result.MaximumQueueWait)
	fmt.Printf("Average latency: %v\n", result.AverageLatency)
	fmt.Printf("Maximum latency: %v\n", result.MaximumLatency)
	fmt.Printf("Tenant job counts: %v\n", result.TenantJobCounts)
	fmt.Printf("Jain fairness index: %.4f\n", result.FairnessIndex)
}
