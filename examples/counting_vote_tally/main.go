// Command counting_vote_tally tallies a survey using Counting Sort.
//
// Counting Sort is optimal when values fall within a small, known integer
// range — here, satisfaction ratings from 1 to 5. The algorithm runs in
// O(n + k) time where k is the range, making it faster than comparison-based
// sorts for this kind of bounded data.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
)

func main() {
	// Post-sprint satisfaction survey responses (1 = very dissatisfied, 5 = very satisfied).
	responses := []int{
		4, 5, 3, 5, 4, 2, 5, 4, 3, 5,
		4, 5, 1, 4, 5, 3, 4, 5, 4, 3,
		5, 4, 2, 5, 4, 5, 3, 4, 5, 4,
	}

	fmt.Println("=== Sprint Retrospective — Satisfaction Survey ===")
	fmt.Printf("Responses collected: %d\n\n", len(responses))

	sorted := ordinex.CountingSorter{}.Sort(responses)

	// Count occurrences.
	tally := make(map[int]int)
	for _, r := range responses {
		tally[r]++
	}

	labels := map[int]string{
		1: "Very dissatisfied",
		2: "Dissatisfied",
		3: "Neutral",
		4: "Satisfied",
		5: "Very satisfied",
	}

	fmt.Printf("%-3s  %-18s  %s\n", "★", "Label", "Count")
	fmt.Println("---  ------------------  -----")
	for i := 5; i >= 1; i-- {
		bar := ""
		for j := 0; j < tally[i]; j++ {
			bar += "█"
		}
		fmt.Printf("%-3d  %-18s  %2d  %s\n", i, labels[i], tally[i], bar)
	}

	total := 0
	for _, r := range sorted {
		total += r
	}
	mean := float64(total) / float64(len(sorted))
	fmt.Printf("\nMean satisfaction score: %.2f / 5.00\n", mean)
}
