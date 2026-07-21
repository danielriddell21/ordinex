package ordinex_test

import (
	"slices"
	"testing"

	"github.com/danielriddell21/ordinex/v2"
)

// The comparison-based sorters are generic, so they sort any cmp.Ordered type,
// not just int.
func TestGenericStringSort(t *testing.T) {
	input := []string{"pear", "apple", "orange", "banana", "apple"}
	want := []string{"apple", "apple", "banana", "orange", "pear"}
	sorters := []ordinex.Sorter[string]{
		ordinex.QuickSorter[string]{},
		ordinex.MergeSorter[string]{},
		ordinex.HeapSorter[string]{},
		ordinex.InsertionSorter[string]{},
	}
	for _, s := range sorters {
		if got := s.Sort(input); !slices.Equal(got, want) {
			t.Errorf("%s.Sort(%v) = %v, want %v", s.Name(), input, got, want)
		}
	}
}

func TestGenericFloatSort(t *testing.T) {
	input := []float64{3.5, 1.2, -4.1, 0.0, 2.2}
	want := []float64{-4.1, 0.0, 1.2, 2.2, 3.5}
	if got := (ordinex.MergeSorter[float64]{}).Sort(input); !slices.Equal(got, want) {
		t.Errorf("Sort(%v) = %v, want %v", input, got, want)
	}
}
