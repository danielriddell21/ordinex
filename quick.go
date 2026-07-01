package ordinex

import (
	"cmp"
	"slices"
)

// QuickSorter implements Quick Sort using the Lomuto partition scheme. It
// selects the last element as the pivot and partitions around it.
//
// The zero value is ready to use.
//
// Time: O(n log n) average, O(n²) worst. Space: O(log n).
type QuickSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Quick Sort".
func (QuickSorter[T]) Name() string { return "Quick Sort" }

// Sort returns a sorted copy of input using Quick Sort. The input is not modified.
func (QuickSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort[T cmp.Ordered](arr []T, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition[T cmp.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
