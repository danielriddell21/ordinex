package ordinex

// PancakeSorter implements Pancake Sort.
// Repeatedly finds the maximum element and uses flips to move it to its correct position,
// similar to sorting pancakes with a spatula.
// Time: O(n²)  Space: O(1)
type PancakeSorter struct{}

func (PancakeSorter) Name() string { return "Pancake Sort" }

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
