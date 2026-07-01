package ordinex

import "slices"

// BucketSorter implements Bucket Sort. It distributes elements into buckets
// based on their value range, sorts each bucket with insertion sort, then
// concatenates the buckets in order.
//
// The zero value is ready to use.
//
// Time: O(n+k) average, O(n²) worst. Space: O(n+k).
type BucketSorter struct{}

// Name returns the algorithm's name, "Bucket Sort".
func (BucketSorter) Name() string { return "Bucket Sort" }

// Sort returns a sorted copy of input using Bucket Sort. The input is not modified.
func (BucketSorter) Sort(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	lo, hi := slices.Min(input), slices.Max(input)
	if lo == hi {
		return slices.Clone(input)
	}
	numBuckets := len(input)
	buckets := make([][]int, numBuckets)
	rangeSize := float64(hi-lo+1) / float64(numBuckets)
	for _, v := range input {
		idx := int(float64(v-lo) / rangeSize)
		if idx >= numBuckets {
			idx = numBuckets - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}
	result := make([]int, 0, len(input))
	for _, bucket := range buckets {
		insertionSortSlice(bucket)
		result = append(result, bucket...)
	}
	return result
}

func insertionSortSlice(arr []int) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
}
