package ordinex

// StalinSorter implements Stalin Sort. It scans the slice once and keeps only
// elements that are greater than or equal to the running maximum. Any element
// smaller than the running maximum is removed from the dataset. Permanently. No
// appeal process.
//
// The zero value is ready to use.
//
// Time: O(n). Space: O(n).
type StalinSorter struct{}

// Name returns the algorithm's name, "Stalin Sort".
func (StalinSorter) Name() string { return "Stalin Sort" }

// Sort returns the elements of input that are in non-decreasing order, dropping
// any element smaller than the running maximum. The result is always sorted but
// may be shorter than input, which is not modified.
func (StalinSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	out := []int{input[0]}
	max := input[0]
	for _, v := range input[1:] {
		if v >= max {
			out = append(out, v)
			max = v
		}
	}
	return out
}
