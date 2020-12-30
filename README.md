# iTree
[![Go Reference](https://pkg.go.dev/badge/dmathieu/itree.svg)](https://pkg.go.dev/dmathieu/itree)
[![CircleCI](https://circleci.com/gh/dmathieu/itree.svg?style=svg)](https://app.circleci.com/pipelines/github/dmathieu/itree)

An Interval Tree Implementation written in Go

## Usage

## With int64 values

```go
intervals := []itree.Interval{
  itree.Interval{Start: 1, End: 3},
  itree.Interval{Start: 2, End: 5},
}

tree, err := itree.NewTree(intervals)
tree.Contains(context.Background(), 3)
```

## With times

```go
now := time.Now()
intervals := []time.Time{
  // The first value is the beginning of the interval, the second is the end
  []time.Time{now.Add(-2 * time.Minute), now.Add(-1 * time.Minute)},
  []time.Time{now.Add(-2 * time.Hour), now.Add(-1 * time.Hour)},
}

tree, err := itree.NewTimesTree(intervals)
tree.Contains(now)
```

## With IP Networks

```go
_, ipnet, _ := net.ParseCIDR("127.0.0.1/32")

intervals := []*net.IPNet{ipnet}
tree, err := itree.NewIPNetTree(intervals)

tree.Contains(net.ParseIP("8.8.8.8"))
```

## Benchmarks

You can run the internal benchmarks using  `go test -bench=. ./...`
Here is the output from a 2019 MacBook Pro

```
BenchmarkBuildTree
BenchmarkBuildTree-12              10000            990813 ns/op
BenchmarkTreeContains
BenchmarkTreeContains-12         2967738               469 ns/op
```
