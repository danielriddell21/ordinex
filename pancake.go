package ordinex

import (
	"cmp"
	"slices"
)

// PancakeSorter implements Pancake Sort. It repeatedly finds the maximum element
// and uses prefix flips to move it into place, much like sorting a stack of
// pancakes with a spatula.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type PancakeSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Pancake Sort".
func (PancakeSorter[T]) Name() string { return "Pancake Sort" }

// Sort returns a sorted copy of input using Pancake Sort. The input is not modified.
func (PancakeSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	for size := len(arr); size > 1; size-- {
		maxIdx := 0
		for i := 1; i < size; i++ {
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}
		if maxIdx == size-1 {
			continue
		}
		if maxIdx != 0 {
			slices.Reverse(arr[:maxIdx+1])
		}
		slices.Reverse(arr[:size])
	}
	return arr
}
