// Command bucket_exam_scores grades a class of exam results using Bucket Sort.
//
// Bucket Sort excels when input values are uniformly distributed over a known
// range — exactly the case for exam scores (0–100). Once sorted, computing
// grade boundaries and percentiles is straightforward.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex/v2"
)

func main() {
	scores := []int{
		73, 88, 45, 92, 61, 78, 55, 83, 70, 96,
		49, 67, 84, 58, 91, 76, 62, 87, 53, 79,
	}

	fmt.Println("=== Exam Results — Year 2 Computer Science ===")
	fmt.Println()
	fmt.Printf("Raw scores: %v\n\n", scores)

	sorted := ordinex.BucketSorter{}.Sort(scores)

	fmt.Printf("Sorted:     %v\n\n", sorted)

	grade := func(s int) string {
		switch {
		case s >= 90:
			return "A"
		case s >= 75:
			return "B"
		case s >= 60:
			return "C"
		case s >= 45:
			return "D"
		default:
			return "F"
		}
	}

	counts := map[string]int{}
	for _, s := range sorted {
		counts[grade(s)]++
	}

	fmt.Println("Grade distribution:")
	for _, g := range []string{"A", "B", "C", "D", "F"} {
		fmt.Printf("  %s: %d student(s)\n", g, counts[g])
	}

	median := sorted[len(sorted)/2]
	fmt.Printf("\nMedian score: %d (%s)\n", median, grade(median))
}
