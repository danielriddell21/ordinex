package ordinex

// StalinSorter implements Stalin Sort.
// Scans the slice once and keeps only elements that are greater than or equal
// to the current maximum. Any element smaller than the running maximum is
// removed from the dataset. Permanently. No appeal process.
//
// The returned slice is guaranteed to be sorted, but may be shorter than the input.
//
// Time: O(n)  Space: O(n)
type StalinSorter struct{}

func (StalinSorter) Name() string { return "Stalin Sort" }

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
