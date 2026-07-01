package ordinex

import (
	"cmp"
	"slices"
)

// CocktailShakerSorter implements Cocktail Shaker Sort, a bidirectional variant
// of Bubble Sort. Each pass alternates direction, shrinking the unsorted region
// from both ends.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type CocktailShakerSorter[T cmp.Ordered] struct{}

// Name returns the algorithm's name, "Cocktail Shaker Sort".
func (CocktailShakerSorter[T]) Name() string { return "Cocktail Shaker Sort" }

// Sort returns a sorted copy of input using Cocktail Shaker Sort. The input is
// not modified.
func (CocktailShakerSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	left, right := 0, len(arr)-1
	for left < right {
		swapped := false
		// Forward pass: bubble max to right
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		right--
		swapped = false
		// Backward pass: bubble min to left
		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		left++
	}
	return arr
}
