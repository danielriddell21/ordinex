package ordinex

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"time"
)

// BogoSorter implements Bogo Sort, also known as Permutation Sort or Stupid
// Sort. It repeatedly shuffles the slice at random until it happens to be
// sorted. It is wildly inefficient and is included for educational and
// entertainment purposes only.
//
// Time: O(n × n!). Space: O(1).
type BogoSorter[T cmp.Ordered] struct {
	// MaxAttempts caps the number of shuffle attempts. A value of 0 means
	// unlimited, which on all but the smallest inputs may never terminate.
	MaxAttempts int

	// Rand is the random source used to shuffle. If nil, a source seeded from
	// the current time is used.
	Rand *rand.Rand
}

// Name returns the algorithm's name, "Bogo Sort".
func (b BogoSorter[T]) Name() string { return "Bogo Sort" }

// Sort returns a sorted copy of input using Bogo Sort. The input is not
// modified. If MaxAttempts is reached before the slice becomes sorted, the
// partially shuffled result is returned as is.
func (b BogoSorter[T]) Sort(input []T) []T {
	arr := slices.Clone(input)
	r := b.Rand
	if r == nil {
		r = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	}
	for attempt := 0; !slices.IsSorted(arr); attempt++ {
		if b.MaxAttempts > 0 && attempt >= b.MaxAttempts {
			break
		}
		bogoShuffle(arr, r)
	}
	return arr
}

func bogoShuffle[T any](arr []T, r *rand.Rand) {
	for i := len(arr) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
