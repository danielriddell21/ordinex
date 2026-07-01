package ordinex

import "slices"

// CountingSorter implements Counting Sort. It counts the frequency of each value
// and reconstructs the sorted slice from those counts. Negative integers are
// supported via a min-value offset.
//
// The zero value is ready to use.
//
// Time: O(n+k). Space: O(k), where k = max-min+1.
type CountingSorter struct{}

// Name returns the algorithm's name, "Counting Sort".
func (CountingSorter) Name() string { return "Counting Sort" }

// Sort returns a sorted copy of input using Counting Sort. The input is not
// modified. Memory use is proportional to the range of values, so inputs with a
// large spread between the smallest and largest values are costly.
func (CountingSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	lo, hi := slices.Min(input), slices.Max(input)
	count := make([]int, hi-lo+1)
	for _, v := range input {
		count[v-lo]++
	}
	result := make([]int, 0, len(input))
	for i, c := range count {
		for range c {
			result = append(result, i+lo)
		}
	}
	return result
}
