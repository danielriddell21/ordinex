package ordinex_test

import (
	"slices"
	"testing"

	"github.com/danielriddell21/ordinex"
)

// fullSorters lists every sorter that returns a full permutation of its input
// in sorted order, so its result must match slices.Sorted. The novelty sorters
// that drop elements (Stalin, Thanos), never terminate deterministically (Bogo,
// Miracle) or reach out to the network (Vibe) are excluded.
func fullSorters() []ordinex.Sorter[int] {
	return []ordinex.Sorter[int]{
		ordinex.BubbleSorter[int]{},
		ordinex.InsertionSorter[int]{},
		ordinex.SelectionSorter[int]{},
		ordinex.ShellSorter[int]{},
		ordinex.GnomeSorter[int]{},
		ordinex.CocktailShakerSorter[int]{},
		ordinex.PancakeSorter[int]{},
		ordinex.QuickSorter[int]{},
		ordinex.MergeSorter[int]{},
		ordinex.HeapSorter[int]{},
		ordinex.CountingSorter{},
		ordinex.RadixSorter{},
		ordinex.BucketSorter{},
	}
}

func FuzzSorters(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{5, 3, 1, 4, 2})
	f.Add([]byte{9, 9, 9})
	f.Fuzz(func(t *testing.T, data []byte) {
		input := make([]int, len(data))
		for i, b := range data {
			input[i] = int(b)
		}
		want := slices.Sorted(slices.Values(input))
		for _, s := range fullSorters() {
			got := s.Sort(input)
			if !slices.Equal(got, want) {
				t.Errorf("%s.Sort(%v) = %v, want %v", s.Name(), input, got, want)
			}
			if !slices.Equal(input, intsFromBytes(data)) {
				t.Errorf("%s.Sort mutated its input", s.Name())
			}
		}
	})
}

func intsFromBytes(data []byte) []int {
	out := make([]int, len(data))
	for i, b := range data {
		out[i] = int(b)
	}
	return out
}
