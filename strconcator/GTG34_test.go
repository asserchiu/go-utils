package strconcator

import (
	"testing"
)

// Basic benchmark

func Benchmark_GTG34_Basic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s = GTG34(s, "Gopher")
	}
}

// Avoid compiler optimization

var GTG34Result string

func Benchmark_GTG34_AvoidCompilerOptimization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		s = GTG34(s, "Gopher")
		GTG34Result = s
	}
}

// Do some loops

func Benchmark_GTG34_Loops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < 10; idx++ {
			s = GTG34(s, "Gopher")
		}
		GTG34Result = s
	}
}

// Do some loops -- Don't make it wrong

func Benchmark_GTG34_LoopsWrong(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < 10; idx++ {
			s = GTG34(s, "Gopher")
		}
	}
	GTG34Result = s
}

// We need more cases

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
