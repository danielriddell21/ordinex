package ordinex_test

import (
	"fmt"
	"math/rand/v2"

	"github.com/danielriddell21/ordinex"
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

// Every sorter satisfies the Sorter interface, so algorithms are interchangeable.
func Example() {
	sorters := []ordinex.Sorter{
		ordinex.QuickSorter{},
		ordinex.MergeSorter{},
		ordinex.HeapSorter{},
	}
	input := []int{5, 3, 1, 4, 2}
	for _, s := range sorters {
		fmt.Printf("%s: %v\n", s.Name(), s.Sort(input))
	}
	// Output:
	// Quick Sort: [1 2 3 4 5]
	// Merge Sort: [1 2 3 4 5]
	// Heap Sort: [1 2 3 4 5]
}
