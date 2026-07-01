package ordinex

import (
	"cmp"
	"slices"
)

// BubbleSorter implements Bubble Sort. It repeatedly swaps adjacent elements
// that are out of order, stopping early once a full pass makes no swaps.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type BubbleSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Bubble Sort".
func (BubbleSorter[T]) Name() string { return "Bubble Sort" }

// Sort returns a sorted copy of input using Bubble Sort. The input is not modified.
func (BubbleSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
