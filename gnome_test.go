package ordinex_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/danielriddell21/ordinex"
)

func TestGnomeSort(t *testing.T) {
	s := ordinex.GnomeSorter[int]{}
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{42}, []int{42}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"all duplicates", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{"with duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5}, []int{1, 1, 2, 3, 4, 5, 5, 6, 9}},
		{"with negatives", []int{-3, 5, -1, 0, 2}, []int{-3, -1, 0, 2, 5}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := copyForTest(tc.input)
			got := s.Sort(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Sort(%v) = %v, want %v", tc.input, got, tc.want)
			}
			if !reflect.DeepEqual(tc.input, original) {
				t.Errorf("Sort mutated input: got %v, want %v", tc.input, original)
			}
		})
	}
	t.Run("random n=50", func(t *testing.T) {
		data := randomSlice(50)
		original := copyForTest(data)
		got := s.Sort(data)
		if !isSortedTest(got) {
			t.Errorf("Sort(%v) = %v, not sorted", data, got)
		}
		if !reflect.DeepEqual(data, original) {
			t.Error("Sort mutated input")
		}
	})
}

func ExampleGnomeSorter() {
	s := ordinex.GnomeSorter[int]{}
	fmt.Println(s.Sort([]int{5, 3, 1, 4, 2}))
	// Output:
	// [1 2 3 4 5]
}

func BenchmarkGnomeSort(b *testing.B) {
	s := ordinex.GnomeSorter[int]{}
	for _, size := range []int{100, 1000, 10000} {
		data := randomSlice(size)
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			for b.Loop() {
				s.Sort(data)
			}
		})
	}
}
