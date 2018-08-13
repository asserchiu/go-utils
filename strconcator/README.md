# strconcator

## TL;DR

Use `strings.Builder` if you are using Go 1.10 or later with zero external dependency.

Use `bytes.Buffer` before Go 1.10 with zero external dependency.

Use/Copy/Enhance this `strconcator` if you do a lot of concatenation from list(slice) of string and you are too lazy to use for-range by yourself.

## Bench result

```
$ go1.11beta3 test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/asserchiu/go-utils/strconcator
Benchmark_10_PlusEqual-4                       	 3000000	       462 ns/op	     208 B/op	       9 allocs/op
Benchmark_10_StringConcator_WriteString-4      	10000000	       217 ns/op	     120 B/op	       5 allocs/op
Benchmark_10_StringConcator_WriteStrings-4     	 3000000	       398 ns/op	     120 B/op	       5 allocs/op
Benchmark_10_BytesBuffer-4                     	10000000	       192 ns/op	     144 B/op	       2 allocs/op
Benchmark_10_StringsBuilder-4                  	10000000	       133 ns/op	      56 B/op	       3 allocs/op
Benchmark_100_PlusEqual-4                      	  200000	      8725 ns/op	   16016 B/op	      99 allocs/op
Benchmark_100_StringConcator_WriteString-4     	 2000000	       832 ns/op	    1368 B/op	       9 allocs/op
Benchmark_100_StringConcator_WriteStrings-4    	  500000	      2921 ns/op	    1368 B/op	       9 allocs/op
Benchmark_100_BytesBuffer-4                    	 1000000	      1708 ns/op	    1440 B/op	       5 allocs/op
Benchmark_100_StringsBuilder-4                 	 2000000	       750 ns/op	    1016 B/op	       7 allocs/op
Benchmark_1000_PlusEqual-4                     	    3000	    477547 ns/op	 1602960 B/op	     999 allocs/op
Benchmark_1000_StringConcator_WriteString-4    	  200000	      6246 ns/op	   13336 B/op	      13 allocs/op
Benchmark_1000_StringConcator_WriteStrings-4   	   50000	     26109 ns/op	   13592 B/op	      14 allocs/op
Benchmark_1000_BytesBuffer-4                   	  100000	     12410 ns/op	   12512 B/op	       8 allocs/op
Benchmark_1000_StringsBuilder-4                	  200000	      5087 ns/op	   10232 B/op	      11 allocs/op
PASS
ok  	github.com/asserchiu/go-utils/strconcator	28.047s
```
