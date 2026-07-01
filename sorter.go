// Package ordinex provides a collection of sorting algorithm implementations,
// each satisfying the [Sorter] interface so that they can be used
// interchangeably.
//
// Comparison-based sorters are generic over any [cmp.Ordered] type. The
// integer-specific sorters — [BucketSorter], [CountingSorter] and [RadixSorter]
// — rely on integer arithmetic, and the novelty sorters [SleepSorter] and
// [VibeSorter] operate on integers too; all five implement Sorter[int].
//
// Most implementations return a slice of the same length as the input and never
// modify the input. [ThanosSorter] and [StalinSorter] are exceptions: they may
// return a shorter slice, because elements are eliminated during sorting.
//
// All Sorters are safe for concurrent use by multiple goroutines, with the
// exception of [VibeSorter], whose Sort performs network I/O.
package ordinex

import "cmp"

// Sorter is implemented by every sorting algorithm in this package.
// Implementations are interchangeable, so a caller can select an algorithm at
// runtime. T is the element type: comparison-based sorters accept any
// [cmp.Ordered] type, while the integer-specific sorters fix T to int.
type Sorter[T cmp.Ordered] interface {
	// Sort returns a sorted copy of input in non-decreasing order. The input
	// slice is never modified. Most implementations preserve every element;
	// [ThanosSorter] and [StalinSorter] may return fewer.
	Sort(input []T) []T

	// Name returns the human-readable name of the algorithm.
	Name() string
}
