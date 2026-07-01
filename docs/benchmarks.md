# Benchmarks

Measured with `go test -bench=. -benchmem -benchtime=3s ./...` on Apple M1 Pro.

Input is a fixed-seed random slice of integers in the range `[-5000, 9999]`.

## Standard sorters

| Algorithm | n=100 | n=1,000 | n=10,000 |
|---|---|---|---|
| `QuickSorter` | 986 ns | 14.8 µs | 497 µs |
| `ShellSorter` | 1.18 µs | 47.7 µs | 921 µs |
| `InsertionSorter` | 1.93 µs | 163 µs | 15.4 ms |
| `HeapSorter` | 2.08 µs | 45.8 µs | 888 µs |
| `RadixSorter` | 3.06 µs | 26.3 µs | 273 µs |
| `BucketSorter` | 3.79 µs | 40.5 µs | 456 µs |
| `MergeSort` | 4.26 µs | 63.2 µs | 996 µs |
| `CocktailShakerSorter` | 4.99 µs | 641 µs | 88.0 ms |
| `SelectionSorter` | 5.71 µs | 476 µs | 42.0 ms |
| `BubbleSorter` | 6.35 µs | 719 µs | 126 ms |
| `GnomeSorter` | 7.34 µs | 687 µs | 66.4 ms |
| `PancakeSorter` | 7.81 µs | 712 µs | 64.1 ms |
| `CountingSorter` | 14.9 µs | 20.0 µs | 74.5 µs |

`CountingSorter` allocates a range-sized buffer on each call; its apparent slowness at n=100 is dominated by the allocation, not the sort itself.

## Unconventional sorters

These algorithms have constraints that prevent standard n=100/1,000/10,000 benchmarking.

| Algorithm | Input | Time | Notes |
|---|---|---|---|
| `StalinSort` | n=1,000 | 990 ns | O(n) single pass; returns fewer elements |
| `ThanosSort` | n=5 | 112 ns | Non-deterministic; benchmarked small to avoid long runs |
| `MiracleSort` | n=1,000 (pre-sorted) | 1.48 µs | Only benchmarkable when input is already sorted |
| `BogoSort` | n=3 | 174 ns | Expected O(n·n!) — not suitable for large inputs |
| `SleepSort` | n=10 | 10.2 ms | Wall-clock bound; time scales with max element value |
| `VibeSorter` | — | varies | Requires `OPENAI_API_KEY`; latency dominated by API round-trip |
