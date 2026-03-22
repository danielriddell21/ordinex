package sortilege

// InsertionSorter implements Insertion Sort.
// Picks each element and inserts it into its correct position in the sorted portion.
// Time: O(n²)  Space: O(1)
type InsertionSorter struct{}

func (InsertionSorter) Name() string { return "Insertion Sort" }

func (InsertionSorter) Sort(input []int) []int {
	arr := copySlice(input)
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
	return arr
}
