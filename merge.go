package sortilege

// MergeSorter implements Merge Sort.
// Recursively divides the slice in half, sorts each half, then merges them.
// Time: O(n log n)  Space: O(n)
type MergeSorter struct{}

func (MergeSorter) Name() string { return "Merge Sort" }

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
