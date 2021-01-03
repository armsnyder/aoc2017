# aoc2017

My [Advent of Code 2017](https://adventofcode.com/2017) solutions.
Days 1-18.1 are missing because I solved them years ago using Python. 

Usage:

```
$ go run . [-d day] [-2]
```

Requires a `session.txt` file containing a session token, for pulling puzzle input and submitting answers.
(Inputs and answers are cached.)

## Benchmarks

I thought it would be fun to share performance [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
for each of my puzzle solutions, since I write benchmarks anyway to help guide my optimizations.
I don't always optimize for the best possible time if I think it impacts code readability.
Benchmarks use the real puzzle input, which is preloaded in memory.

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2017
Benchmark/Day_18/Part_1-16      162001521                7.84 ns/op            0 B/op          0 allocs/op
Benchmark/Day_18/Part_2-16          1123           1043256 ns/op          174266 B/op       7120 allocs/op
Benchmark/Day_19/Part_1-16          8401            145337 ns/op           54120 B/op        214 allocs/op
Benchmark/Day_19/Part_2-16         10000            119705 ns/op           54088 B/op        212 allocs/op
Benchmark/Day_20/Part_1-16          7132            189876 ns/op          110248 B/op       2002 allocs/op
Benchmark/Day_20/Part_2-16            25          46477632 ns/op          390696 B/op       4699 allocs/op
Benchmark/Day_21/Part_1-16          9307            132477 ns/op          100228 B/op        788 allocs/op
Benchmark/Day_21/Part_2-16            28          43247706 ns/op        23317658 B/op      12034 allocs/op
Benchmark/Day_22/Part_1-16          1342            914320 ns/op           74952 B/op         98 allocs/op
Benchmark/Day_22/Part_2-16             1        1268284417 ns/op         8818600 B/op       6429 allocs/op
Benchmark/Day_23/Part_1-16          2706            390187 ns/op            9944 B/op        186 allocs/op
Benchmark/Day_23/Part_2-16            26          43202228 ns/op            9944 B/op        186 allocs/op
Benchmark/Day_24/Part_1-16             8         144083012 ns/op            9682 B/op        133 allocs/op
Benchmark/Day_24/Part_2-16             8         135273390 ns/op            9490 B/op        132 allocs/op
Benchmark/Day_25/Part_1-16             3         335663523 ns/op          185794 B/op       5674 allocs/op
PASS
ok      github.com/armsnyder/aoc2017    22.419s
```
