// Package main sorts a stack of pancakes by diameter using Pancake Sort.
//
// Pancake Sort is the natural algorithm for this exact problem: the only
// permitted operation is flipping a prefix of the stack, which corresponds
// to sliding a spatula under the nth pancake and flipping everything above it.
// The result is a stack with the largest pancake at the bottom.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex/v2"
	"strings"
)

func main() {
	// Pancake diameters in cm, from top to bottom of the stack as served.
	stack := []int{14, 20, 9, 17, 12, 22, 7, 18}

	fmt.Println("=== Pancake Stack Organiser ===")
	fmt.Println()

	draw := func(diameters []int, label string) {
		fmt.Printf("%s\n", label)
		for _, d := range diameters {
			padding := strings.Repeat(" ", (22-d)/2)
			fmt.Printf("  %s%s%s\n", padding, strings.Repeat("═", d), padding)
		}
		fmt.Println()
	}

	draw(stack, "Before (as served):")

	sorted := ordinex.PancakeSorter[int]{}.Sort(stack)

	// Reverse for display: largest at bottom.
	display := make([]int, len(sorted))
	for i, v := range sorted {
		display[len(sorted)-1-i] = v
	}

	draw(display, "After (largest at bottom):")

	fmt.Printf("Smallest pancake: %d cm\n", sorted[0])
	fmt.Printf("Largest pancake:  %d cm\n", sorted[len(sorted)-1])
}
