package ordinex_test

import "github.com/danielriddell21/ordinex/v2"

// Every sorter satisfies Sorter[int]; the comparison-based ones also satisfy
// Sorter for other cmp.Ordered element types.
var (
	_ ordinex.Sorter[int] = ordinex.BubbleSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.InsertionSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.SelectionSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.ShellSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.GnomeSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.CocktailShakerSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.PancakeSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.QuickSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.MergeSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.HeapSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.BogoSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.MiracleSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.ThanosSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.StalinSorter[int]{}
	_ ordinex.Sorter[int] = ordinex.CountingSorter{}
	_ ordinex.Sorter[int] = ordinex.RadixSorter{}
	_ ordinex.Sorter[int] = ordinex.BucketSorter{}
	_ ordinex.Sorter[int] = ordinex.SleepSorter{}
	_ ordinex.Sorter[int] = ordinex.VibeSorter{}

	_ ordinex.Sorter[string] = ordinex.QuickSorter[string]{}
)
