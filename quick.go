package sortilege

// QuickSorter implements Quick Sort using the Lomuto partition scheme.
// Selects the last element as pivot and partitions around it.
// Time: O(n log n) average, O(n²) worst  Space: O(log n)
type QuickSorter struct{}

func (QuickSorter) Name() string { return "Quick Sort" }

func (QuickSorter) Sort(input []int) []int {
	arr := copySlice(input)
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
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
