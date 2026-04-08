package ordinex

// GnomeSorter implements Gnome Sort (also known as Stupid Sort).
// Moves each element to its correct position by a series of adjacent swaps,
// similar to a garden gnome rearranging flower pots.
// Time: O(n²)  Space: O(1)
type GnomeSorter struct{}

func (GnomeSorter) Name() string { return "Gnome Sort" }

func (GnomeSorter) Sort(input []int) []int {
	arr := copySlice(input)
	n := len(arr)
	pos := 0
	for pos < n {
		if pos == 0 || arr[pos] >= arr[pos-1] {
			pos++
		} else {
			arr[pos], arr[pos-1] = arr[pos-1], arr[pos]
			pos--
		}
	}
	return arr
}
