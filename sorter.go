// Package ordinex provides a collection of sorting algorithm implementations,
// each satisfying the [Sorter] interface so that they can be used
// interchangeably.
//
// Most implementations return a slice of the same length as the input and never
// modify the input. [ThanosSorter] and [StalinSorter] are exceptions: they may
// return a shorter slice, because elements are eliminated during sorting.
//
// All Sorters are safe for concurrent use by multiple goroutines.
package ordinex

// Sorter is implemented by every sorting algorithm in this package.
// Implementations are interchangeable, so a caller can select an algorithm at
// runtime.
type Sorter interface {
	// Sort returns a sorted copy of input in non-decreasing order. The input
	// slice is never modified. Most implementations preserve every element;
	// [ThanosSorter] and [StalinSorter] may return fewer.
	Sort(input []int) []int

	// Name returns the human-readable name of the algorithm.
	Name() string
}

// copySlice returns a copy of s.
func copySlice(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	return out
}

// isSorted reports whether s is sorted in non-decreasing order.
func isSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
