# How to write benchmarks in Go

2018-08-28 Golang Taipei Gathering #34 - Lightning Talk

1. How to write benchmarks in Go.
2. Evaluating performance improvement.

## How to write benchmarks in Go

``` bash
$ go test -bench=.
$ go test -bench=. -benchmem
```

### Basic benchmark

``` go
func Benchmark_PlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PlusEqual("GTG34")
	}
}
```

``` go
func Benchmark_PlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s += "GTG34"
	}
}
```

### Avoid compiler optimization
Always store the result to a package level variable so the compiler cannot eliminate the Benchmark itself.

``` go
var result string

func Benchmark_PlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s += "GTG34"
		result = s
	}
}
```

### Do some loops
Some operations has different cost when using different memory usage

``` go
var result string

func Benchmark_PlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < 10; idx++ {
			s += "GTG34"
		}
		result = s
	}
}
```

### Do some loops -- Don't make it wrong

``` go
var result string

func Benchmark_PlusEqual(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < 10; idx++ {
			s += "GTG34"
		}
	}
	result = s
}
```

### We need more cases
Extract benchmark logic and make iteration count as parameter

``` go
var result string

func Benchmark_PlusEqual_10(b *testing.B) {
	BenchCase_PlusEqual(b, 10)
}

func Benchmark_PlusEqual_100(b *testing.B) {
	BenchCase_PlusEqual(b, 100)
}

func Benchmark_PlusEqual_1000(b *testing.B) {
	BenchCase_PlusEqual(b, 1000)
}

func BenchCase_PlusEqual(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < benchCount; idx++ {
			s += "GTG34"
		}
		result = s
	}
}
```

## Evaluating performance improvement

### Prepare input data (Run Benchmarks)
``` bash
$ go test -run=NONE -bench=. -benchmem | tee new.txt
$ go test -run=NONE -bench=. -benchmem -count=10 | tee new10.txt
$ git stash
$ go test -run=NONE -bench=. -benchmem | tee old.txt
$ go test -run=NONE -bench=. -benchmem -count=10 | tee old10.txt
$ git stash pop
```

### Benchmarking has a timeout
``` bash
$ go test -run=NONE -bench=. -benchmem -count=100 | tee old100.txt
...
*** Test killed with quit: ran too long (10m0s).
exit status 2
FAIL	github.com/asserchiu/go-utils/strconcator	600.006s
```

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

### Run it!
``` bash
$ benchcmp old.txt new.txt
$ benchstat old.txt new.txt
$ benchcmp old10.txt new10.txt
$ benchstat old10.txt new10.txt
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
