package ordinex

import (
	"cmp"
	"slices"
)

// SelectionSorter implements Selection Sort. It repeatedly finds the minimum
// element in the unsorted portion and moves it to the front.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type SelectionSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Selection Sort".
func (SelectionSorter[T]) Name() string { return "Selection Sort" }

// Sort returns a sorted copy of input using Selection Sort. The input is not modified.
func (SelectionSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return arr
}
