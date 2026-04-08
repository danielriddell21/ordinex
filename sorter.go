// Package ordinex provides a collection of sorting algorithm implementations.
//
// All implementations satisfy the Sorter interface. Most return a slice of the
// same length as the input. ThanosSorter and StalinSorter are exceptions — they
// may return a shorter slice as elements are eliminated during sorting.
package ordinex

// Sorter is implemented by every sorting algorithm in this package.
type Sorter interface {
	// Sort returns a sorted copy of input. The original slice is never modified.
	// Most implementations preserve all elements; ThanosSorter and StalinSorter
	// may return fewer.
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
