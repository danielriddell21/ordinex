# ordinex

*n.* divination by drawing lots. Also: making your slices behave.

One interface. No sacrifices required.

[![Go Reference](https://pkg.go.dev/badge/github.com/danielriddell21/ordinex.svg)](https://pkg.go.dev/github.com/danielriddell21/ordinex)
[![CI](https://github.com/danielriddell21/ordinex/actions/workflows/ci.yaml/badge.svg)](https://github.com/danielriddell21/ordinex/actions/workflows/ci.yaml)
[![Go 1.26](https://img.shields.io/badge/go-1.26-blue)](https://go.dev)
[![MIT License](https://img.shields.io/badge/licence-MIT-green)](LICENSE)

## Installation

```sh
go get github.com/danielriddell21/ordinex@latest
```

## Quick start

```go
import "github.com/danielriddell21/ordinex"

s := ordinex.MergeSorter{}
sorted := s.Sort([]int{5, 3, 1, 4, 2})
// sorted == [1 2 3 4 5], original is unchanged
fmt.Println(s.Name()) // "Merge Sort"
```

Every sorter satisfies the `Sorter` interface:

```go
type Sorter interface {
    Sort(input []int) []int
    Name() string
}
```

`Sort` returns a sorted **copy** — the original slice is never modified.
`ThanosSorter` and `StalinSorter` may return a shorter slice than the input.

## Algorithms

| Sorter | Time (avg) | Time (worst) | Space | Notes |
|---|---|---|---|---|
| `BubbleSorter` | O(n²) | O(n²) | O(1) | Early exit if no swaps |
| `InsertionSorter` | O(n²) | O(n²) | O(1) | |
| `SelectionSorter` | O(n²) | O(n²) | O(1) | |
| `CocktailShakerSorter` | O(n²) | O(n²) | O(1) | Bidirectional bubble sort |
| `GnomeSorter` | O(n²) | O(n²) | O(1) | |
| `ShellSorter` | O(n²) | O(n²) | O(1) | Gap sequence: n/2 halved each step |
| `QuickSorter` | O(n log n) | O(n²) | O(log n) | Lomuto partition, last-element pivot |
| `MergeSorter` | O(n log n) | O(n log n) | O(n) | |
| `HeapSorter` | O(n log n) | O(n log n) | O(1) | Max-heap |
| `CountingSorter` | O(n+k) | O(n+k) | O(k) | k = value range; supports negatives |
| `RadixSorter` | O(d·(n+k)) | O(d·(n+k)) | O(n+k) | LSD base-10; supports negatives |
| `BucketSorter` | O(n+k) | O(n²) | O(n+k) | Buckets sorted with insertion sort |
| `PancakeSorter` | O(n²) | O(n²) | O(1) | Sorts by reversing prefixes |
| `BogoSorter` | O(n·n!) | O(∞) | O(1) | Shuffle until sorted |
| `SleepSorter` | O(max) | O(max) | O(n) | Goroutine per element |
| `MiracleSorter` | O(∞) | O(∞) | O(1) | Waits for a cosmic ray |
| `ThanosSorter` | O(n) | O(n) | O(n) | Eliminates half until sorted |
| `VibeSorter` | O($) | O(☁) | O(☁) | Sends the array to an LLM. Prays. |
| `StalinSorter` | O(n) | O(n) | O(n) | Removes any element smaller than the running maximum |

## Configurable sorters

Most sorters are empty structs and need no configuration. A few have fields:

### BogoSorter

```go
s := ordinex.BogoSorter{
    MaxAttempts: 1000,                        // 0 = unlimited
    Rand:        rand.New(rand.NewPCG(42, 0)), // nil = seeded from time
}
```

### SleepSorter

```go
s := ordinex.SleepSorter{
    ScaleFactor: 10 * time.Millisecond, // sleep per unit of value; default 1ms
}
```

### MiracleSorter

```go
s := ordinex.MiracleSorter{
    MaxChecks: 1000, // 0 = no limit
}
```

### VibeSorter

```go
s := ordinex.VibeSorter{
    APIKey: "sk-...",     // if empty, uses OPENAI_API_KEY env var
    Model:  "gpt-4o",    // if empty, defaults to "gpt-4o-mini"
}
```

### ThanosSorter

```go
s := ordinex.ThanosSorter{
    Rand: rand.New(rand.NewPCG(42, 0)), // nil = seeded from time
}
```

## Examples

Runnable examples for every algorithm are in the [`examples/`](examples/EXAMPLES.md) directory.

## Documentation
- [Benchmark results](docs/benchmarks.md) — measured on Apple M1 Pro
- [Complexity chart](docs/complexity.md) — visual comparison of all algorithms
