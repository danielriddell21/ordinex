package ordinex_test

import (
	"math/rand/v2"
	"reflect"
	"testing"

	"github.com/danielriddell21/ordinex"
)

func TestBogoSort(t *testing.T) {
	r := rand.New(rand.NewPCG(12345, 0))
	s := ordinex.BogoSorter[int]{MaxAttempts: 10000, Rand: r}

	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{7}, []int{7}},
		{"two elements sorted", []int{1, 2}, []int{1, 2}},
		{"two elements unsorted", []int{2, 1}, []int{1, 2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := copyForTest(tc.input)
			got := s.Sort(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("BogoSort(%v) = %v, want %v", tc.input, got, tc.want)
			}
			if !reflect.DeepEqual(tc.input, original) {
				t.Error("BogoSort mutated input")
			}
		})
	}
}

func BenchmarkBogoSort(b *testing.B) {
	r := rand.New(rand.NewPCG(42, 0))
	s := ordinex.BogoSorter[int]{MaxAttempts: 100000, Rand: r}
	data := []int{3, 1, 2}
	b.Run("n=3", func(b *testing.B) {
		for b.Loop() {
			s.Sort(data)
		}
	})
}
