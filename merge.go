package ordinex

// MergeSorter implements Merge Sort. It recursively divides the slice in half,
// sorts each half, then merges the two sorted halves.
//
// The zero value is ready to use.
//
// Time: O(n log n). Space: O(n).
type MergeSorter struct{}

// Name returns the algorithm's name, "Merge Sort".
func (MergeSorter) Name() string { return "Merge Sort" }

// Sort returns a sorted copy of input using Merge Sort. The input is not modified.
func (MergeSorter) Sort(input []int) []int {
	arr := copySlice(input)
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return mergeSorted(left, right)
}

func mergeSorted(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
