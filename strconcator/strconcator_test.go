package strconcator

import (
	"bytes"
	"strings"
	"testing"
)

// TODO: Update these tests

//func Test_New(t *testing.T) {
//	type args struct {
//		ss []string
//	}
//	tests := []struct {
//		name   string
//		args   args
//		wantSc *StringConcator
//	}{
//		{`New()`, args{}, &StringConcator{}},
//		{`New(nil)`, args{nil}, &StringConcator{}},
//		{`New([])`, args{[]string{}}, &StringConcator{}},
//		{`New(["abc123"])`, args{[]string{"abc123"}}, &StringConcator{[]byte("abc123")}},
//		{`New(["a", "b", "c", "1", "2", "3"])`, args{[]string{"a", "b", "c", "1", "2", "3"}}, &StringConcator{[]byte("abc123")}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotSc := New(tt.args.ss...); !reflect.DeepEqual(gotSc, tt.wantSc) {
//				t.Errorf("New() = %v, want %v", gotSc, tt.wantSc)
//			}
//		})
//	}
//}
//
//func Test_StringConcator_WriteString(t *testing.T) {
//	type fields struct {
//		raw []byte
//	}
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantN   int
//		wantErr bool
//		wantSc  *StringConcator
//	}{
//		{`WriteString("")`, fields{}, args{""}, 0, false, &StringConcator{}},
//		{`WriteString("abc")`, fields{}, args{"abc"}, 3, false, &StringConcator{[]byte("abc")}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &StringConcator{
//				raw: tt.fields.raw,
//			}
//			gotN, err := r.WriteString(tt.args.s)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("StringConcator.WriteString() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if gotN != tt.wantN {
//				t.Errorf("StringConcator.WriteString() = %v, want %v", gotN, tt.wantN)
//			}
//		})
//	}
//}
//
//func Test_StringConcator_WriteStrings(t *testing.T) {
//	type fields struct {
//		raw []byte
//	}
//	type args struct {
//		ss []string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		wantSc *StringConcator
//	}{
//		{`WriteStrings()`, fields{}, args{}, &StringConcator{}},
//		{`WriteStrings(nil)`, fields{}, args{nil}, &StringConcator{}},
//		{`WriteStrings([])`, fields{}, args{[]string{}}, &StringConcator{}},
//		{`WriteStrings(["abc123"])`, fields{}, args{[]string{"abc123"}}, &StringConcator{[]byte("abc123")}},
//		{`WriteStrings(["a", "b", "c", "1", "2", "3"])`, fields{}, args{[]string{"a", "b", "c", "1", "2", "3"}}, &StringConcator{[]byte("abc123")}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &StringConcator{
//				raw: tt.fields.raw,
//			}
//			r.WriteStrings(tt.args.ss...)
//			if !reflect.DeepEqual(r, tt.wantSc) {
//				t.Errorf("StringConcator.WriteStrings() = %v, want %v", r, tt.wantSc)
//			}
//		})
//	}
//}
//
//func Test_StringConcator_String(t *testing.T) {
//	type fields struct {
//		raw []byte
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		{"empty", fields{}, ""},
//		{"abc123", fields{[]byte("abc123")}, "abc123"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &StringConcator{
//				raw: tt.fields.raw,
//			}
//			if got := r.String(); got != tt.want {
//				t.Errorf("StringConcator.String() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

// go test -bench . -benchmem

// always store the result to a package level variable
// so the compiler cannot eliminate the Benchmark itself.
var result string

func Benchmark_10_PlusEqual(b *testing.B) {
	BenchCase_PlusEqual(b, 10)
}
func Benchmark_10_StringConcator_WriteString(b *testing.B) {
	BenchCase_StringConcator_WriteString(b, 10)
}
func Benchmark_10_StringConcator_WriteStrings(b *testing.B) {
	BenchCase_StringConcator_WriteStrings(b, 10)
}
func Benchmark_10_BytesBuffer(b *testing.B) {
	BenchCase_BytesBuffer(b, 10)
}
func Benchmark_10_StringsBuilder(b *testing.B) {
	BenchCase_StringsBuilder(b, 10)
}

func Benchmark_100_PlusEqual(b *testing.B) {
	BenchCase_PlusEqual(b, 100)
}
func Benchmark_100_StringConcator_WriteString(b *testing.B) {
	BenchCase_StringConcator_WriteString(b, 100)
}
func Benchmark_100_StringConcator_WriteStrings(b *testing.B) {
	BenchCase_StringConcator_WriteStrings(b, 100)
}
func Benchmark_100_BytesBuffer(b *testing.B) {
	BenchCase_BytesBuffer(b, 100)
}
func Benchmark_100_StringsBuilder(b *testing.B) {
	BenchCase_StringsBuilder(b, 100)
}

func Benchmark_1000_PlusEqual(b *testing.B) {
	BenchCase_PlusEqual(b, 1000)
}
func Benchmark_1000_StringConcator_WriteString(b *testing.B) {
	BenchCase_StringConcator_WriteString(b, 1000)
}
func Benchmark_1000_StringConcator_WriteStrings(b *testing.B) {
	BenchCase_StringConcator_WriteStrings(b, 1000)
}
func Benchmark_1000_BytesBuffer(b *testing.B) {
	BenchCase_BytesBuffer(b, 1000)
}
func Benchmark_1000_StringsBuilder(b *testing.B) {
	BenchCase_StringsBuilder(b, 1000)
}

func BenchCase_PlusEqual(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		var s string
		for idx := 0; idx < benchCount; idx++ {
			s += "xxx"
		}
		result = s
	}
}

func BenchCase_StringConcator_WriteString(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		sc := New()
		for idx := 0; idx < benchCount; idx++ {
			sc.WriteString("xxx")
		}
		result = sc.String()
	}
}

func BenchCase_StringConcator_WriteStrings(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		sc := New()
		for idx := 0; idx < benchCount; idx++ {
			sc.WriteStrings("x", "x", "x")
		}
		result = sc.String()
	}
}

func BenchCase_BytesBuffer(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for idx := 0; idx < benchCount; idx++ {
			buf.WriteString("xxx")
		}
		result = buf.String()
	}
}

func BenchCase_StringsBuilder(b *testing.B, benchCount int) {
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for idx := 0; idx < benchCount; idx++ {
			b.WriteString("xxx")
		}
		result = b.String()
	}
}
