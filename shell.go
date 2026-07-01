package ordinex

import (
	"cmp"
	"slices"
)

// ShellSorter implements Shell Sort, an optimisation of Insertion Sort that
// exchanges far-apart elements first. It uses a gap sequence starting at n/2 and
// halved on each iteration.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type ShellSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Shell Sort".
func (ShellSorter[T]) Name() string { return "Shell Sort" }

// Sort returns a sorted copy of input using Shell Sort. The input is not modified.
func (ShellSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
	return arr
}
