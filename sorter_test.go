package ordinex_test

import (
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/danielriddell21/ordinex"
)

func copyForTest(s []int) []int {
	return slices.Clone(s)
}

func isSortedTest(s []int) bool {
	return slices.IsSorted(s)
}

func randomSlice(n int) []int {
	s := make([]int, n)
	r := rand.New(rand.NewPCG(42, 0))
	for i := range s {
		s[i] = r.IntN(10000) - 5000
	}
	return s
}

// Every comparison-based sorter satisfies Sorter[int], so algorithms are
// interchangeable.
func Example() {
	sorters := []ordinex.Sorter[int]{
		ordinex.QuickSorter[int]{},
		ordinex.MergeSorter[int]{},
		ordinex.HeapSorter[int]{},
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
