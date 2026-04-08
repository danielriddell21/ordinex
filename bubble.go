package ordinex

// BubbleSorter implements Bubble Sort.
// Repeatedly swaps adjacent elements that are out of order.
// Stops early if no swaps occur in a full pass.
// Time: O(n²)  Space: O(1)
type BubbleSorter struct{}

func (BubbleSorter) Name() string { return "Bubble Sort" }

func (BubbleSorter) Sort(input []int) []int {
	arr := copySlice(input)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
