package ordinex_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/danielriddell21/ordinex"
)

func TestVibeSort(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set — skipping live API tests")
	}

	s := ordinex.VibeSorter{}

	t.Run("empty", func(t *testing.T) {
		got := s.Sort([]int{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		got := s.Sort([]int{42})
		if !reflect.DeepEqual(got, []int{42}) {
			t.Errorf("expected [42], got %v", got)
		}
	})

	t.Run("already sorted", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("VibeSort(%v) = %v, want sorted", input, got)
		}
		if !reflect.DeepEqual(input, original) {
			t.Error("VibeSort mutated input")
		}
	})

	t.Run("unsorted", func(t *testing.T) {
		input := []int{5, 3, 1, 4, 2}
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("VibeSort(%v) = %v, want sorted", input, got)
		}
		if !reflect.DeepEqual(input, original) {
			t.Error("VibeSort mutated input")
		}
	})

	t.Run("random n=10", func(t *testing.T) {
		input := randomSlice(10)
		original := copyForTest(input)
		got := s.Sort(input)
		if !isSortedTest(got) {
			t.Errorf("VibeSort(%v) = %v, want sorted", input, got)
		}
		if !reflect.DeepEqual(input, original) {
			t.Error("VibeSort mutated input")
		}
	})
}

func BenchmarkVibeSort(b *testing.B) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		b.Skip("OPENAI_API_KEY not set")
	}
	s := ordinex.VibeSorter{}
	data := randomSlice(10)
	b.Run("n=10", func(b *testing.B) {
		for b.Loop() {
			s.Sort(data)
		}
	})
}
