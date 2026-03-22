package sortilege

import (
	"math/rand/v2"
	"time"
)

// ThanosSorter implements Thanos Sort.
// Checks if the slice is sorted. If not, it randomly eliminates half the elements
// and checks again. Repeats until the remaining elements happen to be sorted.
// The returned slice may be shorter than the input.
//
// Rand is the random source; if nil, one is seeded from the current time.
type ThanosSorter struct {
	Rand *rand.Rand
}

func (t ThanosSorter) Name() string { return "Thanos Sort" }

func (t ThanosSorter) Sort(input []int) []int {
	arr := copySlice(input)
	r := t.Rand
	if r == nil {
		r = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	}
	for !isSorted(arr) {
		arr = thanosSnap(arr, r)
	}
	return arr
}

// thanosSnap randomly selects ceil(n/2) elements from arr.
func thanosSnap(arr []int, r *rand.Rand) []int {
	n := len(arr)
	keep := (n + 1) / 2
	// Fisher-Yates partial shuffle: select `keep` elements
	tmp := make([]int, n)
	copy(tmp, arr)
	for i := 0; i < keep; i++ {
		j := i + r.IntN(n-i)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp[:keep]
}
