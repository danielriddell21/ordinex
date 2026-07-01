package ordinex_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/danielriddell21/ordinex"
)

func TestThanosSort(t *testing.T) {
	r := rand.New(rand.NewPCG(99, 0))
	s := ordinex.ThanosSorter[int]{Rand: r}

	t.Run("empty", func(t *testing.T) {
		got := s.Sort([]int{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
	t.Run("single", func(t *testing.T) {
		got := s.Sort([]int{5})
		if len(got) != 1 || got[0] != 5 {
			t.Errorf("expected [5], got %v", got)
		}
	})
	t.Run("already sorted preserves length", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("ThanosSort(%v) = %v, not sorted", input, got)
		}
		if len(got) != len(input) {
			t.Errorf("already-sorted input: expected length %d, got %d", len(input), len(got))
		}
		if !sliceEqual(input, original) {
			t.Error("ThanosSort mutated input")
		}
	})
	t.Run("result is sorted", func(t *testing.T) {
		// Thanos may return fewer elements; we only check that what remains is sorted.
		input := []int{5, 3, 1, 4, 2}
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("ThanosSort(%v) = %v, not sorted", input, got)
		}
		if !sliceEqual(input, original) {
			t.Error("ThanosSort mutated input")
		}
	})
}

func ExampleThanosSorter() {
	// A fixed random source makes the elimination deterministic. The result is
	// always sorted but may be shorter than the input.
	s := ordinex.ThanosSorter[int]{Rand: rand.New(rand.NewPCG(1, 0))}
	result := s.Sort([]int{5, 3, 1, 4, 2})
	fmt.Println(result)
	// Output:
	// [1 4]
}

func BenchmarkThanosSort(b *testing.B) {
	r := rand.New(rand.NewPCG(42, 0))
	s := ordinex.ThanosSorter[int]{Rand: r}
	data := []int{3, 1, 2, 4, 5}
	b.Run("n=5", func(b *testing.B) {
		for b.Loop() {
			s.Sort(data)
		}
	})
}
