# iTree
[![Go Reference](https://pkg.go.dev/badge/dmathieu/itree.svg)](https://pkg.go.dev/dmathieu/itree)
[![CircleCI](https://circleci.com/gh/dmathieu/itree.svg?style=svg)](https://app.circleci.com/pipelines/github/dmathieu/itree)

An Interval Tree Implementation written in Go

## Usage

```go
intervals := []itree.Interval{
  itree.Interval{Start: 1, End: 3},
  itree.Interval{Start: 2, End: 5},
}

tree, err := itree.NewTree(intervals)
tree.Contains(context.Background(), 3)
```

## Benchmarks

You can run the internal benchmarks using  `go test -bench=. ./...`
Here is the output from a 2019 MacBook Pro

```
BenchmarkBuildTree
BenchmarkBuildTree-12              10000           2758013 ns/op
BenchmarkTreeContains
BenchmarkTreeContains-12         2977766               638 ns/op
```
