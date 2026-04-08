package ordinex

// BucketSorter implements Bucket Sort.
// Distributes elements into buckets based on value range, sorts each bucket
// with insertion sort, then concatenates the results.
// Time: O(n+k)  Space: O(n+k)
type BucketSorter struct{}

func (BucketSorter) Name() string { return "Bucket Sort" }

func (BucketSorter) Sort(input []int) []int {
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
	if min == max {
		return copySlice(input)
	}
	numBuckets := len(input)
	buckets := make([][]int, numBuckets)
	rangeSize := float64(max-min+1) / float64(numBuckets)
	for _, v := range input {
		idx := int(float64(v-min) / rangeSize)
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
