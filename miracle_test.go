package ordinex_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/danielriddell21/ordinex/v2"
)

func TestMiracleSort(t *testing.T) {
	t.Run("already sorted returns immediately", func(t *testing.T) {
		s := ordinex.MiracleSorter[int]{} // MaxChecks=0 is safe when input is sorted
		input := []int{1, 2, 3, 4, 5}
		original := copyForTest(input)
		got := s.Sort(input)
		if !reflect.DeepEqual(got, input) {
			t.Errorf("MiracleSort(%v) = %v, want %v", input, got, input)
		}
		if !reflect.DeepEqual(input, original) {
			t.Error("MiracleSort mutated input")
		}
	})
	t.Run("empty", func(t *testing.T) {
		s := ordinex.MiracleSorter[int]{}
		got := s.Sort([]int{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
	t.Run("single element", func(t *testing.T) {
		s := ordinex.MiracleSorter[int]{}
		got := s.Sort([]int{42})
		if !reflect.DeepEqual(got, []int{42}) {
			t.Errorf("expected [42], got %v", got)
		}
	})
	t.Run("unsorted with MaxChecks exits without hanging", func(t *testing.T) {
		// With MaxChecks=1, Sort exits after one check without blocking.
		s := ordinex.MiracleSorter[int]{MaxChecks: 1}
		input := []int{5, 3, 1}
		got := s.Sort(input)
		// Result is returned (unsorted) — miracle didn't happen.
		if len(got) != len(input) {
			t.Errorf("expected same length as input, got %v", got)
		}
	})
}

func BenchmarkMiracleSort(b *testing.B) {
	// Only benchmark with already-sorted input to avoid infinite loop.
	s := ordinex.MiracleSorter[int]{}
	data := randomSlice(1000)
	// Pre-sort the data so MiracleSort returns immediately.
	sortedData := ordinex.MergeSorter[int]{}.Sort(data)
	b.Run("n=1000_presorted", func(b *testing.B) {
		for b.Loop() {
			s.Sort(sortedData)
		}
	})
}

// MiracleSorter returns immediately when the input is already sorted.
func ExampleMiracleSorter() {
	s := ordinex.MiracleSorter[int]{}
	fmt.Println(s.Sort([]int{1, 2, 3, 4, 5}))
	// Output:
	// [1 2 3 4 5]
}
