package ordinex

// PancakeSorter implements Pancake Sort. It repeatedly finds the maximum element
// and uses prefix flips to move it into place, much like sorting a stack of
// pancakes with a spatula.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type PancakeSorter struct{}

// Name returns the algorithm's name, "Pancake Sort".
func (PancakeSorter) Name() string { return "Pancake Sort" }

// Sort returns a sorted copy of input using Pancake Sort. The input is not modified.
func (PancakeSorter) Sort(input []int) []int {
	arr := copySlice(input)
	for size := len(arr); size > 1; size-- {
		// Find index of the maximum element in arr[0..size-1]
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
			pancakeFlip(arr, maxIdx)
		}
		pancakeFlip(arr, size-1)
	}
	return arr
}

// pancakeFlip reverses arr[0..k].
func pancakeFlip(arr []int, k int) {
	for start := 0; start < k; start, k = start+1, k-1 {
		arr[start], arr[k] = arr[k], arr[start]
	}
}
