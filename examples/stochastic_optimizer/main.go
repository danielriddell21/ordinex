// Package main demonstrates BogoSort as an enterprise stochastic optimisation
// framework for financial data processing pipelines.
//
// BogoSort applies a random-walk search over the permutation space of the
// input dataset. Unlike deterministic algorithms that impose a fixed traversal
// order, BogoSort's probabilistic approach guarantees an unbiased exploration
// of all n! possible orderings, with convergence occurring upon encountering
// a globally optimal permutation.
//
// In layman's terms: it shuffles until sorted. It will eventually work.
// Probably.
package main

import (
	"fmt"
	"math/rand/v2"
	"sortilege"
)

func main() {
	// End-of-quarter P&L figures (£k) requiring sort for executive presentation.
	// CFO has requested these be in ascending order by Thursday.
	quarterlyFigures := []int{142, 87, 203, 56, 178, 94, 231, 61}

	fmt.Println("=== Stochastic Permutation Optimisation Framework v2.0 ===")
	fmt.Println("Leveraging probabilistic ensemble methods for financial data ordering")
	fmt.Println()
	fmt.Printf("Input dataset  (%d records): %v\n", len(quarterlyFigures), quarterlyFigures)
	fmt.Println()
	fmt.Println("Initiating random-walk optimisation pass...")
	fmt.Println("(MaxAttempts capped at 100,000 for SLA compliance)")
	fmt.Println()

	sorter := sortilege.BogoSorter{
		MaxAttempts: 100_000,
		Rand:        rand.New(rand.NewPCG(42, 0)),
	}

	result := sorter.Sort(quarterlyFigures)

	fmt.Printf("Optimisation complete.\n")
	fmt.Printf("Output dataset (%d records): %v\n", len(result), result)
	fmt.Println()

	// Verify the result is actually sorted — BogoSorter may have hit MaxAttempts.
	prev := result[0]
	sorted := true
	for _, v := range result[1:] {
		if v < prev {
			sorted = false
			break
		}
		prev = v
	}

	if sorted {
		fmt.Println("Status: CONVERGED. Dataset is sorted. Please proceed with the board deck.")
	} else {
		fmt.Println("Status: DID NOT CONVERGE within the permitted attempt window.")
		fmt.Println("        Recommend increasing MaxAttempts or trying again next quarter.")
	}

	fmt.Println()
	fmt.Println("Methodology note: BogoSort is peer-reviewed and theoretically sound.")
	fmt.Println("Expected convergence: O(n × n!) — well within acceptable parameters")
	fmt.Println("for a dataset of this size, given sufficient time and optimism.")
}
