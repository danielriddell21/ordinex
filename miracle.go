package ordinex

// MiracleSorter implements Miracle Sort.
// Checks if the slice is sorted. If not, it waits for a cosmic ray to flip a bit
// in memory and sorts the array. Repeats until a miracle happens.
//
// Time: O(∞)  Space: O(1)
type MiracleSorter struct {
	MaxChecks int // 0 = no limit
}

func (MiracleSorter) Name() string { return "Miracle Sort" }

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
