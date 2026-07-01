package ordinex

import (
	"cmp"
	"slices"
)

// HeapSorter implements Heap Sort. It builds a max-heap, then repeatedly
// extracts the root to produce a sorted slice.
//
// The zero value is ready to use.
//
// Time: O(n log n). Space: O(1).
type HeapSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Heap Sort".
func (HeapSorter[T]) Name() string { return "Heap Sort" }

// Sort returns a sorted copy of input using Heap Sort. The input is not modified.
func (HeapSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	return arr
}

func heapify[T cmp.Ordered](arr []T, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
