package ordinex_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/danielriddell21/ordinex/v2"
)

func TestStalinSort(t *testing.T) {
	s := ordinex.StalinSorter[int]{}

	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"all equal", []int{3, 3, 3}, []int{3, 3, 3}},
		{"removes non-conforming elements", []int{1, 3, 2, 4, 1, 5}, []int{1, 3, 4, 5}},
		{"keeps only first when descending", []int{5, 4, 3, 2, 1}, []int{5}},
		{"duplicates at max retained", []int{1, 3, 3, 2, 4}, []int{1, 3, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := copyForTest(tc.input)
			got := s.Sort(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("StalinSort(%v) = %v, want %v", tc.input, got, tc.want)
			}
			if !reflect.DeepEqual(tc.input, original) {
				t.Error("StalinSort mutated input")
			}
		})
	}

	t.Run("result is always sorted", func(t *testing.T) {
		input := randomSlice(50)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("StalinSort(%v) = %v, result is not sorted", input, got)
		}
	})
}

func BenchmarkStalinSort(b *testing.B) {
	s := ordinex.StalinSorter[int]{}
	data := randomSlice(1000)
	b.Run("n=1000", func(b *testing.B) {
		for b.Loop() {
			s.Sort(data)
		}
	})
}

// StalinSorter removes any element smaller than the running maximum.
// The returned slice is sorted but may be shorter than the input.
func ExampleStalinSorter() {
	s := ordinex.StalinSorter[int]{}
	fmt.Println(s.Sort([]int{3, 1, 4, 1, 5, 9, 2, 6}))
	// Output:
	// [3 4 5 9]
}
