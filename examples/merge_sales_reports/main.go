// Package main combines two pre-sorted regional sales reports using Merge Sort.
//
// Merge Sort's divide-and-conquer structure maps naturally onto scenarios where
// data arrives in sorted partitions — such as independently sorted reports from
// separate regions or time windows — and needs to be combined into a single
// sorted sequence efficiently.
package main

import (
	"fmt"
	"github.com/danielriddell21/sortilege"
)

func main() {
	// Q4 order values (£) — already sorted within each region's report.
	northRegion := []int{120, 245, 310, 450, 780, 920}
	southRegion := []int{85, 195, 275, 510, 640, 870, 1050}

	combined := append(northRegion, southRegion...)

	fmt.Println("=== Q4 Sales Report — Combined Regions ===")
	fmt.Println()
	fmt.Printf("North region (%d orders): %v\n", len(northRegion), northRegion)
	fmt.Printf("South region (%d orders): %v\n", len(southRegion), southRegion)
	fmt.Println()

	sorted := sortilege.MergeSorter{}.Sort(combined)

	fmt.Printf("Combined     (%d orders): %v\n\n", len(sorted), sorted)

	total := 0
	for _, v := range sorted {
		total += v
	}
	mean := total / len(sorted)
	p90 := sorted[int(float64(len(sorted))*0.9)]

	fmt.Printf("Total revenue:  £%d\n", total)
	fmt.Printf("Mean order:     £%d\n", mean)
	fmt.Printf("Median order:   £%d\n", sorted[len(sorted)/2])
	fmt.Printf("p90 order:      £%d\n", p90)
	fmt.Printf("Largest order:  £%d\n", sorted[len(sorted)-1])
}
