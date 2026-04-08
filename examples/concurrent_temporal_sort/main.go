// Package main demonstrates SleepSort as a high-throughput concurrent sorting
// solution built on Go's goroutine scheduler.
//
// SleepSort achieves sorted output through time-domain element separation.
// Each element is assigned a dedicated goroutine which sleeps for a duration
// proportional to its magnitude. Smaller values wake earlier, naturally
// producing ascending order without comparison operations or auxiliary memory.
//
// This is the logical conclusion of "let the runtime handle it."
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
	"time"
)

func main() {
	// Sprint velocity data from the last six quarters.
	// Sorting required for roadmap prioritisation dashboard.
	sprintVelocities := []int{34, 12, 45, 7, 28, 51, 19, 3, 40, 22}

	fmt.Println("=== Concurrent Temporal Sort Engine (CTSE) ===")
	fmt.Println("Powered by Go's world-class concurrency primitives")
	fmt.Println()
	fmt.Printf("Input  (%d elements): %v\n", len(sprintVelocities), sprintVelocities)
	fmt.Println()
	fmt.Println("Dispatching goroutines...")
	fmt.Println("Each element will self-report when its temporal window opens.")
	fmt.Println("Please wait. This is by design.")
	fmt.Println()

	start := time.Now()

	sorter := ordinex.SleepSorter{
		// Each unit of value sleeps for 10ms.
		// Increasing this improves accuracy at the cost of your afternoon.
		ScaleFactor: 10 * time.Millisecond,
	}

	sorted := sorter.Sort(sprintVelocities)

	elapsed := time.Since(start)

	fmt.Printf("Output (%d elements): %v\n", len(sorted), sorted)
	fmt.Println()
	fmt.Printf("Sort completed in %v\n", elapsed.Round(time.Millisecond))
	fmt.Println()
	fmt.Println("Performance note: runtime scales with the largest element value,")
	fmt.Println("not the number of elements. Large numbers will make this very slow.")
	fmt.Println("This is a feature, not a limitation. The algorithm is working correctly.")
}
