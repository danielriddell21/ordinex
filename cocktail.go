package ordinex

// CocktailShakerSorter implements Cocktail Shaker Sort (bidirectional Bubble Sort).
// Each pass alternates direction, shrinking the unsorted region from both ends.
// Time: O(n²)  Space: O(1)
type CocktailShakerSorter struct{}

func (CocktailShakerSorter) Name() string { return "Cocktail Shaker Sort" }

func (CocktailShakerSorter) Sort(input []int) []int {
	arr := copySlice(input)
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
