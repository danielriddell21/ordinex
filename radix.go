package ordinex

import "slices"

// RadixSorter implements Radix Sort, least-significant-digit first in base 10.
// It sorts numbers by processing their digits from least to most significant.
// Negative integers are handled by sorting the negatives and non-negatives
// separately and recombining them.
//
// The zero value is ready to use.
//
// Time: O(d*(n+k)). Space: O(n+k), where d = digits and k = 10.
type RadixSorter struct{}

// Name returns the algorithm's name, "Radix Sort".
func (RadixSorter) Name() string { return "Radix Sort" }

// Sort returns a sorted copy of input using Radix Sort. The input is not modified.
func (RadixSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	var neg, pos []int
	for _, v := range input {
		if v < 0 {
			neg = append(neg, -v)
		} else {
			pos = append(pos, v)
		}
	}
	pos = radixLSD(pos)
	neg = radixLSD(neg)
	slices.Reverse(neg)
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
	hi := slices.Max(arr)
	for exp := 1; hi/exp > 0; exp *= 10 {
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
