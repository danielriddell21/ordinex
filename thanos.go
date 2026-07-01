package ordinex

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"time"
)

// ThanosSorter implements Thanos Sort. It checks whether the slice is sorted
// and, if not, randomly eliminates half the elements and checks again, repeating
// until the survivors happen to be sorted.
//
// Time: O(n). Space: O(n).
type ThanosSorter[T cmp.Ordered] struct {
	// Rand is the random source used to choose which elements survive. If nil, a
	// source seeded from the current time is used.
	Rand *rand.Rand
}

// Name returns the algorithm's name, "Thanos Sort".
func (t ThanosSorter[T]) Name() string { return "Thanos Sort" }

// Sort returns a sorted subsequence of input. The input is not modified. The
// result is always sorted but may be shorter than input, since elements are
// discarded until the remainder is in order.
func (t ThanosSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	r := t.Rand
	if r == nil {
		r = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	}
	for !slices.IsSorted(arr) {
		arr = thanosSnap(arr, r)
	}
	return arr
}

func thanosSnap[T any](arr []T, r *rand.Rand) []T {
	n := len(arr)
	keep := (n + 1) / 2
	tmp := slices.Clone(arr)
	for i := 0; i < keep; i++ {
		j := i + r.IntN(n-i)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp[:keep]
}
