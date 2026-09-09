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

	fmt.Println()
	fmt.Println("Priority Only")
	printResult(priorityResult)

	fmt.Println()
	fmt.Println("Priority + Aging")
	printResult(agingResult)

	fmt.Println()
	fmt.Println("=== Comparison ===")

	fmt.Printf(
		"Low-priority queue wait: %v → %v\n",
		priorityResult.LowPriorityWait,
		agingResult.LowPriorityWait,
	)

	fmt.Printf(
		"Average queue wait: %v → %v\n",
		priorityResult.AverageQueueWait,
		agingResult.AverageQueueWait,
	)
}

func printResult(result scheduler.BenchmarkResult) {
	fmt.Printf("Aging factor: %v\n", result.AgingFactor)
	fmt.Printf("Jobs completed: %d\n", result.JobsCompleted)
	fmt.Printf("Low-priority wait: %v\n", result.LowPriorityWait)
	fmt.Printf("Average queue wait: %v\n", result.AverageQueueWait)
	fmt.Printf("Maximum queue wait: %v\n", result.MaximumQueueWait)
	fmt.Printf("Average latency: %v\n", result.AverageLatency)
	fmt.Printf("Maximum latency: %v\n", result.MaximumLatency)
}
