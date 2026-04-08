# Contributing

Pull requests are welcome.

## Prerequisites

Go 1.25 or later.

## Running tests

```sh
go test ./...
```

## Running benchmarks

```sh
go test -bench=. -benchmem ./...
```

## Adding a new sorter

1. Create `<name>.go` implementing the `Sorter` interface:

```go
type MySorter struct{}

func (MySorter) Name() string { return "My Sort" }

func (MySorter) Sort(input []int) []int {
    out := copySlice(input) // always copy — never mutate input
    // ... sort out ...
    return out
}
```

2. Create `<name>_test.go` in package `ordinex_test`. Use the shared helpers from `sorter_test.go` (`copyForTest`, `isSortedTest`, `randomSlice`). Include at minimum: empty input, single element, already sorted, unsorted, and a random-slice test. Add a benchmark.

3. Add an example in `examples/<scenario>/main.go`. Keep it runnable with `go run`. See existing examples for tone and structure.

4. Add a row to the algorithms table in `README.md` and a row in `examples/EXAMPLES.md`.
