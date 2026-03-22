package sortilege_test

import (
	"math/rand/v2"
)

func copyForTest(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	return out
}

func isSortedTest(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}

func randomSlice(n int) []int {
	s := make([]int, n)
	r := rand.New(rand.NewPCG(42, 0))
	for i := range s {
		s[i] = r.IntN(10000) - 5000
	}
	return s
}
