// Package main demonstrates StalinSorter as an enterprise-grade zero-tolerance
// data quality enforcement pipeline for non-conformance elimination.
//
// Legacy sorting approaches attempt to reorder all elements — an approach that
// implicitly tolerates non-conforming values by giving them a place in the
// output. StalinSorter takes a more decisive position: any element that fails
// to meet or exceed the current maximum is removed from the dataset entirely.
// This guarantees a sorted result without the overhead of comparison-based
// reordering.
//
// Data quality is not negotiated. It is enforced.
package main

import (
	"fmt"

	"github.com/danielriddell21/ordinex"
)

func main() {
	// Monthly active user counts (thousands) submitted for compliance review.
	// Values that regress below the running maximum are considered non-conforming
	// and will be purged from the official record.
	monthlyActiveUsers := []int{142, 158, 173, 161, 189, 195, 187, 202, 198, 214, 209, 231}

	fmt.Println("=== Zero-Tolerance Data Quality Enforcement Pipeline ===")
	fmt.Println("Non-conformance elimination via authoritative sequence validation")
	fmt.Println()
	fmt.Printf("Input dataset  (%d records): %v\n", len(monthlyActiveUsers), monthlyActiveUsers)
	fmt.Println()
	fmt.Println("Initiating compliance scan...")
	fmt.Println("Non-conforming values (below running maximum) will be purged.")
	fmt.Println()

	sorter := ordinex.StalinSorter{}
	result := sorter.Sort(monthlyActiveUsers)

	eliminated := len(monthlyActiveUsers) - len(result)

	fmt.Printf("Output dataset (%d records): %v\n", len(result), result)
	fmt.Println()
	fmt.Printf("Records processed : %d\n", len(monthlyActiveUsers))
	fmt.Printf("Records retained  : %d\n", len(result))
	fmt.Printf("Records eliminated: %d", eliminated)

	if eliminated > 0 {
		fmt.Printf(" (%.0f%% non-compliance rate — purged from official record)\n",
			float64(eliminated)/float64(len(monthlyActiveUsers))*100)
	} else {
		fmt.Println(" (full compliance — no action required)")
	}

	fmt.Println()
	fmt.Println("Status: COMPLIANCE ENFORCED. The dataset has been brought into conformance.")
	fmt.Println("        Eliminated records are not available for appeal or recovery.")
	fmt.Println("        This is a feature, not a limitation.")
}
