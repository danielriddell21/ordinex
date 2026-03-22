package sortilege

// SelectionSorter implements Selection Sort.
// Repeatedly finds the minimum element in the unsorted portion and moves it to the front.
// Time: O(n²)  Space: O(1)
type SelectionSorter struct{}

func (SelectionSorter) Name() string { return "Selection Sort" }

func (SelectionSorter) Sort(input []int) []int {
	arr := copySlice(input)
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	return arr
}
