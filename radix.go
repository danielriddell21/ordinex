package sortilege

// RadixSorter implements Radix Sort (LSD, base-10).
// Sorts numbers by processing digits from least to most significant.
// Handles negative integers by sorting negatives and non-negatives separately.
// Time: O(d*(n+k))  Space: O(n+k)  where d = digits, k = 10
type RadixSorter struct{}

func (RadixSorter) Name() string { return "Radix Sort" }

func (RadixSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	// Split into non-negatives and negatives
	var neg, pos []int
	for _, v := range input {
		if v < 0 {
			neg = append(neg, -v) // store absolute values
		} else {
			pos = append(pos, v)
		}
	}
	pos = radixLSD(pos)
	neg = radixLSD(neg)
	// Reverse negatives: largest absolute value is most negative
	for i, j := 0, len(neg)-1; i < j; i, j = i+1, j-1 {
		neg[i], neg[j] = neg[j], neg[i]
	}
	result := make([]int, 0, len(input))
	for _, v := range neg {
		result = append(result, -v)
	}
	result = append(result, pos...)
	return result
}

func radixLSD(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	for exp := 1; max/exp > 0; exp *= 10 {
		arr = countingByDigit(arr, exp)
	}
	return arr
}

func countingByDigit(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)
	for _, v := range arr {
		digit := (v / exp) % 10
		count[digit]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		count[digit]--
		output[count[digit]] = arr[i]
	}
	return output
}
