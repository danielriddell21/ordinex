package ordinex

// GnomeSorter implements Gnome Sort, also known as Stupid Sort. It moves each
// element to its correct position through a series of adjacent swaps, much like
// a garden gnome rearranging flower pots.
//
// The zero value is ready to use.
//
// Time: O(n²). Space: O(1).
type GnomeSorter struct{}

// Name returns the algorithm's name, "Gnome Sort".
func (GnomeSorter) Name() string { return "Gnome Sort" }

// Sort returns a sorted copy of input using Gnome Sort. The input is not modified.
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
