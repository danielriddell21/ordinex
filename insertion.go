package ordinex

// InsertionSorter implements Insertion Sort. It takes each element in turn and
// inserts it into its correct position within the already-sorted prefix.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type InsertionSorter struct{}

// Name returns the algorithm's name, "Insertion Sort".
func (InsertionSorter) Name() string { return "Insertion Sort" }

// Sort returns a sorted copy of input using Insertion Sort. The input is not modified.
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
