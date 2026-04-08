package ordinex

import (
	"math/rand/v2"
	"time"
)

// BogoSorter implements Bogo Sort (also known as Permutation Sort or Stupid Sort).
// Repeatedly shuffles the slice at random until it happens to be sorted.
// This is highly inefficient and is included for educational/humour purposes only.
//
// MaxAttempts caps the number of shuffle attempts (0 = unlimited, dangerous).
// Rand is the random source; if nil, one is seeded from the current time.
// Time: O(n × n!)  Space: O(1)
type BogoSorter struct {
	MaxAttempts int
	Rand        *rand.Rand
}

func (b BogoSorter) Name() string { return "Bogo Sort" }

func (b BogoSorter) Sort(input []int) []int {
	arr := copySlice(input)
	r := b.Rand
	if r == nil {
		r = rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	}
	for attempt := 0; !isSorted(arr); attempt++ {
		if b.MaxAttempts > 0 && attempt >= b.MaxAttempts {
			break
		}
		bogoShuffle(arr, r)
	}
	return arr
}

func bogoShuffle(arr []int, r *rand.Rand) {
	for i := len(arr) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
