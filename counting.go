package sortilege

// CountingSorter implements Counting Sort.
// Counts the frequency of each value and reconstructs the sorted slice.
// Supports negative integers via a min-value offset.
// Time: O(n+k)  Space: O(k)  where k = max-min+1
type CountingSorter struct{}

func (CountingSorter) Name() string { return "Counting Sort" }

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
