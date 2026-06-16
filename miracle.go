package ordinex

// MiracleSorter implements Miracle Sort. It checks whether the slice is sorted
// and, if not, waits for a cosmic ray to flip a bit in memory into the right
// place, then checks again. It repeats until a miracle happens.
//
// Time: O(∞). Space: O(1).
type MiracleSorter struct {
	// MaxChecks caps the number of sorted-ness checks before giving up. A value
	// of 0 means no limit, so Sort will not return until the input is already
	// sorted.
	MaxChecks int
}

// Name returns the algorithm's name, "Miracle Sort".
func (MiracleSorter) Name() string { return "Miracle Sort" }

// Sort returns a copy of input once it is observed to be sorted. The input is
// not modified. Because no miracle is ever performed, Sort returns the input
// unchanged only when it is already sorted or MaxChecks is reached.
func (m MiracleSorter) Sort(input []int) []int {
	arr := copySlice(input)
	for checks := 0; !isSorted(arr); checks++ {
		if m.MaxChecks > 0 && checks >= m.MaxChecks {
			break
		}
		// Wait for a miracle (cosmic ray flipping a bit in memory).
		// In practice this never happens, so we just check again.
	}
	return arr
}
