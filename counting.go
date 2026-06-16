package ordinex

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
	min, max := input[0], input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	count := make([]int, max-min+1)
	for _, v := range input {
		count[v-min]++
	}
	result := make([]int, 0, len(input))
	for i, c := range count {
		for range c {
			result = append(result, i+min)
		}
	}
	return result
}
