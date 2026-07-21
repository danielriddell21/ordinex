package ordinex_test

import (
	"testing"
	"time"

	"github.com/danielriddell21/ordinex/v2"
)

func TestSleepSort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Sleep Sort test in short mode")
	}
	// A generous per-unit scale keeps adjacent values far enough apart that
	// scheduler jitter cannot wake them out of order on a loaded machine.
	s := ordinex.SleepSorter{ScaleFactor: 10 * time.Millisecond}
	// Only test with small positive inputs; Sleep Sort is non-deterministic with duplicates.
	cases := [][]int{
		{},
		{1},
		{3, 1, 2},
		{5, 2, 8, 1, 9},
	}
	for _, input := range cases {
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("SleepSort(%v) = %v, not sorted", input, got)
		}
		if !sliceEqual(input, original) {
			t.Errorf("SleepSort mutated input: got %v, want %v", input, original)
		}
	}
}

func BenchmarkSleepSort(b *testing.B) {
	s := ordinex.SleepSorter{ScaleFactor: time.Millisecond}
	data := []int{5, 2, 8, 1, 3, 9, 4, 7, 6, 10}
	b.Run("n=10", func(b *testing.B) {
		for b.Loop() {
			s.Sort(data)
		}
	})
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
