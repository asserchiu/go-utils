# How to write benchmarks in Go

2018-08-28 Golang Taipei Gathering #34 - Lightning Talk

* How to run benchmarks
* How to write benchmarks in Go.
* Evaluating performance improvement.

## How to run benchmarks

Run tests matching `NONE` and benchmarks matching `.` (everything)

``` bash
$ go test -run=NONE -bench=.
```

PLUS memory allocation statistics for benchmarks

``` bash
$ go test -run=NONE -bench=. -benchmem
```

AND running it multiple times

``` bash
$ go test -run=NONE -bench=. -benchmem -count=3
```

## How to write benchmarks in Go

### Basic benchmark

``` go
func Benchmark_GTG34_Basic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s = GTG34(s, "Gopher")
	}
}
```

### Avoid compiler optimization

Always store the result to a package level variable so the compiler cannot eliminate the Benchmark itself.

``` go
var GTG34Result string

func Benchmark_GTG34_AvoidCompilerOptimization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s = GTG34(s, "Gopher")
		GTG34Result = s
	}
}
```

### Do some loops

Some operations has different cost under different memory usage

``` go
var GTG34Result string

func Benchmark_GTG34_Loops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < 10; idx++ {
			s = GTG34(s, "Gopher")
		}
		GTG34Result = s
	}
}
```

### Do some loops -- Don't make it wrong

``` go
var GTG34Result string

func Benchmark_GTG34_LoopsWrong(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < 10; idx++ {
			s = GTG34(s, "Gopher")
		}
	}
	GTG34Result = s
}
```

### We need more cases

Extract benchmark logic and make iteration count as parameter

``` go
var GTG34Result string

func Benchmark_GTG34_MoreCase_10(b *testing.B) {
	BenchCase_GTG34(b, 10)
}

func Benchmark_GTG34_MoreCase_100(b *testing.B) {
	BenchCase_GTG34(b, 100)
}

func Benchmark_GTG34_MoreCase_1000(b *testing.B) {
	BenchCase_GTG34(b, 1000)
}

func BenchCase_GTG34(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < benchCount; idx++ {
			s = GTG34(s, "Gopher")
		}
		GTG34Result = s
	}
}
```

## Evaluating performance improvement

### `benchcmp` -- naïve direct comparison

``` bash
$ go get golang.org/x/tools/cmd/benchcmp
$ benchcmp
usage: benchcmp old.txt new.txt

  -best
		compare best times from old and new
  -changed
		show only benchmarks that have changed
  -mag
		sort benchmarks by magnitude of change

Each input file should be from:
	go test -run=NONE -bench=. > [old,new].txt

Benchcmp compares old and new for each benchmark.

If -test.benchmem=true is added to the "go test" command
benchcmp will also compare memory allocations.
```

### `benchstat` -- statistical comparison

``` bash
$ go get golang.org/x/perf/cmd/benchstat
$ benchstat 
usage: benchstat [options] old.txt [new.txt] [more.txt ...]
options:
  -alpha α
		consider change significant if p < α (default 0.05)
  -delta-test test
		significance test to apply to delta: utest, ttest, or none (default "utest")
  -geomean
		print the geometric mean of each file
  -html
		print results as an HTML table
  -sort order
		sort by order: [-]delta, [-]name, none (default "none")
  -split labels
		split benchmarks by labels (default "pkg,goos,goarch")
```

### Prepare input data (Run benchmarks)

``` bash
$ git stash
...
$ go test -run=NONE -bench=GTG34 -benchmem | tee GTG34_bench_before_x1.txt
goos: darwin
goarch: amd64
pkg: github.com/asserchiu/go-utils/strconcator
Benchmark_GTG34_Basic-4                       	10000000	       112 ns/op	      48 B/op	       2 allocs/op
Benchmark_GTG34_AvoidCompilerOptimization-4   	10000000	       114 ns/op	      48 B/op	       2 allocs/op
Benchmark_GTG34_Loops-4                       	  500000	      2251 ns/op	    2992 B/op	      29 allocs/op
Benchmark_GTG34_LoopsWrong-4                  	   10000	   4803541 ns/op	27121295 B/op	      30 allocs/op
Benchmark_GTG34_MoreCase_10-4                 	  500000	      2343 ns/op	    2992 B/op	      29 allocs/op
Benchmark_GTG34_MoreCase_100-4                	   20000	     82677 ns/op	  287040 B/op	     299 allocs/op
Benchmark_GTG34_MoreCase_1000-4               	     200	   6688279 ns/op	28606401 B/op	    2999 allocs/op
PASS
ok  	github.com/asserchiu/go-utils/strconcator	57.569s

$ git stash pop
...
$ go test -run=NONE -bench=GTG34 -benchmem | tee GTG34_bench_after_x1.txt
goos: darwin
goarch: amd64
pkg: github.com/asserchiu/go-utils/strconcator
Benchmark_GTG34_Basic-4                       	10000000	       132 ns/op	      56 B/op	       3 allocs/op
Benchmark_GTG34_AvoidCompilerOptimization-4   	10000000	       137 ns/op	      56 B/op	       3 allocs/op
Benchmark_GTG34_Loops-4                       	  500000	      2653 ns/op	    2696 B/op	      21 allocs/op
Benchmark_GTG34_LoopsWrong-4                  	   10000	   2067687 ns/op	 9065086 B/op	      10 allocs/op
Benchmark_GTG34_MoreCase_10-4                 	 1000000	      1885 ns/op	    2696 B/op	      21 allocs/op
Benchmark_GTG34_MoreCase_100-4                	   30000	     45322 ns/op	  123560 B/op	     135 allocs/op
Benchmark_GTG34_MoreCase_1000-4               	     500	   2873632 ns/op	 9764145 B/op	    1057 allocs/op
PASS
ok  	github.com/asserchiu/go-utils/strconcator	32.143s
```

### Try `benchcmp`

Compare performance between two bench results.

``` bash
$ benchcmp GTG34_bench_before_x1.txt GTG34_bench_after_x1.txt
benchmark                                       old ns/op     new ns/op     delta
Benchmark_GTG34_Basic-4                         112           132           +17.86%
Benchmark_GTG34_AvoidCompilerOptimization-4     114           137           +20.18%
Benchmark_GTG34_Loops-4                         2251          2653          +17.86% <--- !
Benchmark_GTG34_LoopsWrong-4                    4803541       2067687       -56.95%
Benchmark_GTG34_MoreCase_10-4                   2343          1885          -19.55%
Benchmark_GTG34_MoreCase_100-4                  82677         45322         -45.18%
Benchmark_GTG34_MoreCase_1000-4                 6688279       2873632       -57.03%

benchmark                                       old allocs     new allocs     delta
Benchmark_GTG34_Basic-4                         2              3              +50.00%
Benchmark_GTG34_AvoidCompilerOptimization-4     2              3              +50.00%
Benchmark_GTG34_Loops-4                         29             21             -27.59%
Benchmark_GTG34_LoopsWrong-4                    30             10             -66.67%
Benchmark_GTG34_MoreCase_10-4                   29             21             -27.59%
Benchmark_GTG34_MoreCase_100-4                  299            135            -54.85%
Benchmark_GTG34_MoreCase_1000-4                 2999           1057           -64.75%

benchmark                                       old bytes     new bytes     delta
Benchmark_GTG34_Basic-4                         48            56            +16.67%
Benchmark_GTG34_AvoidCompilerOptimization-4     48            56            +16.67%
Benchmark_GTG34_Loops-4                         2992          2696          -9.89%
Benchmark_GTG34_LoopsWrong-4                    27121295      9065086       -66.58%
Benchmark_GTG34_MoreCase_10-4                   2992          2696          -9.89%
Benchmark_GTG34_MoreCase_100-4                  287040        123560        -56.95%
Benchmark_GTG34_MoreCase_1000-4                 28606401      9764145       -65.87%
```

### Try `benchstat`

Compare performance on single runs may lead to wrong result. (see `benchcmp` Benchmark_GTG34_Loops-4)

``` bash
$ git stash
$ go test -run=NONE -bench=GTG34 -benchmem -count=5 | tee GTG34_bench_before_x5.txt

$ git stash pop
$ go test -run=NONE -bench=GTG34 -benchmem -count=5 | tee GTG34_bench_after_x5.txt

$ benchstat GTG34_bench_before_x5.txt GTG34_bench_after_x5.txt
name                                old time/op    new time/op    delta
_GTG34_Basic-4                         114ns ± 1%     131ns ± 1%  +15.54%  (p=0.016 n=5+4)
_GTG34_AvoidCompilerOptimization-4     117ns ± 3%     139ns ±14%  +19.18%  (p=0.008 n=5+5)
_GTG34_Loops-4                        2.48µs ±22%    1.73µs ± 3%  -30.20%  (p=0.008 n=5+5)
_GTG34_LoopsWrong-4                   5.40ms ± 7%    2.04ms ± 6%  -62.15%  (p=0.008 n=5+5)
_GTG34_MoreCase_10-4                  2.46µs ± 7%    1.86µs ±14%  -24.17%  (p=0.008 n=5+5)
_GTG34_MoreCase_100-4                  108µs ±48%      40µs ± 1%  -63.11%  (p=0.016 n=5+4)
_GTG34_MoreCase_1000-4                7.23ms ± 5%    2.76ms ± 7%  -61.78%  (p=0.008 n=5+5)

name                                old alloc/op   new alloc/op   delta
_GTG34_Basic-4                         48.0B ± 0%     56.0B ± 0%  +16.67%  (p=0.008 n=5+5)
_GTG34_AvoidCompilerOptimization-4     48.0B ± 0%     56.0B ± 0%  +16.67%  (p=0.008 n=5+5)
_GTG34_Loops-4                        2.99kB ± 0%    2.70kB ± 0%   -9.89%  (p=0.008 n=5+5)
_GTG34_LoopsWrong-4                   27.1MB ± 0%     9.1MB ± 0%  -66.58%  (p=0.008 n=5+5)
_GTG34_MoreCase_10-4                  2.99kB ± 0%    2.70kB ± 0%   -9.89%  (p=0.008 n=5+5)
_GTG34_MoreCase_100-4                  287kB ± 0%     124kB ± 0%  -56.95%  (p=0.008 n=5+5)
_GTG34_MoreCase_1000-4                28.6MB ± 0%     9.8MB ± 0%  -65.87%  (p=0.008 n=5+5)

name                                old allocs/op  new allocs/op  delta
_GTG34_Basic-4                          2.00 ± 0%      3.00 ± 0%  +50.00%  (p=0.008 n=5+5)
_GTG34_AvoidCompilerOptimization-4      2.00 ± 0%      3.00 ± 0%  +50.00%  (p=0.008 n=5+5)
_GTG34_Loops-4                          29.0 ± 0%      21.0 ± 0%  -27.59%  (p=0.008 n=5+5)
_GTG34_LoopsWrong-4                     30.0 ± 0%      10.0 ± 0%  -66.67%  (p=0.008 n=5+5)
_GTG34_MoreCase_10-4                    29.0 ± 0%      21.0 ± 0%  -27.59%  (p=0.008 n=5+5)
_GTG34_MoreCase_100-4                    299 ± 0%       135 ± 0%  -54.85%  (p=0.008 n=5+5)
_GTG34_MoreCase_1000-4                 3.00k ± 0%     1.06k ± 0%  -64.75%  (p=0.008 n=5+5)
```

### Benchmarking has a timeout

``` bash
$ go test -run=NONE -bench=. -benchmem -count=100
...
*** Test killed with quit: ran too long (10m0s).
exit status 2
FAIL	github.com/asserchiu/go-utils/strconcator	600.006s
```

# Reference

testing - The Go Programming Language
https://golang.org/pkg/testing/

benchcmp - GoDoc
https://godoc.org/golang.org/x/tools/cmd/benchcmp

benchstat - GoDoc
https://godoc.org/golang.org/x/perf/cmd/benchstat

How to write benchmarks in Go | Dave Cheney
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

Optimizations in Go are easy!
https://talks.godoc.org/github.com/mvdan/talks/2018/benchstat.slide
