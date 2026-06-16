// Command enterprise_data_pipeline demonstrates the ThanosSort algorithm in a
// production data pipeline context.
//
// ThanosSort is an industry-leading data reduction strategy that eliminates
// overhead by probabilistically removing elements until the dataset achieves
// natural sorted order. This approach is particularly effective in environments
// where having less data is considered an acceptable trade-off for correctness.
//
// Recommended for: cost optimisation, headcount reduction, any situation where
// half the problem disappearing is treated as the problem being solved.
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/danielriddell21/ordinex"
)

func main() {
	// Simulated Q3 performance metrics (ms) ingested from the observability platform.
	// These represent real production load — do not discard arbitrarily.
	serverResponseTimes := []int{
		312, 87, 450, 23, 198, 561, 74, 390, 115, 263,
		482, 38, 177, 509, 92, 344, 201, 67, 415, 130,
	}

	fmt.Println("=== Enterprise Observability Data Pipeline v4.2 ===")
	fmt.Println()
	fmt.Printf("Ingesting %d metrics from upstream telemetry...\n", len(serverResponseTimes))
	fmt.Printf("Raw payload: %v\n", serverResponseTimes)
	fmt.Println()

	// ThanosSorter is configured with a fixed seed for reproducible data loss.
	sorter := ordinex.ThanosSorter{
		Rand: rand.New(rand.NewPCG(99, 0)),
	}

	fmt.Printf("Applying %s reduction pipeline...\n", sorter.Name())
	fmt.Println("(Balancing the dataset for long-term infrastructure sustainability)")
	fmt.Println()

	optimised := sorter.Sort(serverResponseTimes)

	fmt.Printf("Pipeline complete. Optimised payload (%d metrics): %v\n", len(optimised), optimised)
	fmt.Println()
	fmt.Printf("Data reduction achieved: %d metrics eliminated (%.0f%% overhead removed)\n",
		len(serverResponseTimes)-len(optimised),
		float64(len(serverResponseTimes)-len(optimised))/float64(len(serverResponseTimes))*100,
	)
	fmt.Println()
	fmt.Println("NOTE: Missing metrics are not a bug. They have been gracefully retired")
	fmt.Println("      in accordance with our data lifecycle management policy.")
	fmt.Println("      Please update your SLOs accordingly.")
}
