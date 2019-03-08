package genome_test

import (
	"testing"

	. "github.com/kamilsk/platform/internal/sandbox/genome"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		src  []T
	}{
		{
			"nil",
			nil,
		},
		{
			"empty",
			[]T{},
		},
		{
			"not empty",
			[]T{1, 2, 3},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.src, Copy(tc.src))
		})
	}
}

// BenchmarkCopy/presented-4         	20000000	        52.1 ns/op	      80 B/op	       1 allocs/op
// BenchmarkCopy/append_to_nil-4     	20000000	        58.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkCopy/append_to_src-4     	20000000	        60.2 ns/op	      80 B/op	       1 allocs/op
func BenchmarkCopy(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T) []T
	}{
		{"presented", Copy},
		{"append to nil", func(src []T) []T { return append([]T(nil), src...) }},
		{"append to src", func(src []T) []T { return append(src[:0:0], src...) }},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name, func(b *testing.B) {
			origin := []T{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(origin)
			}
		})
	}
}

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		origin   []T
		from, to int
		expected []T
	}{
		{
			"left",
			[]T{1, 2, 3, 4, 5},
			0, 2,
			[]T{3, 4, 5},
		},
		{
			"right",
			[]T{1, 2, 3, 4, 5},
			3, 5,
			[]T{1, 2, 3},
		},
		{
			"center",
			[]T{1, 2, 3, 4, 5},
			1, 4,
			[]T{1, 5},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Cut(tc.origin, tc.from, tc.to))
		})
	}
}

// BenchmarkCut/presented-4         	30000000	        47.7 ns/op	      80 B/op	       1 allocs/op
// BenchmarkCut/gc_safe-4           	30000000	        49.8 ns/op	      80 B/op	       1 allocs/op
func BenchmarkCut(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T, int, int) []T
	}{
		{"presented", Cut},
		{"gc safe", func(src []T, from, to int) []T {
			copy(src[from:], src[to:])
			for k, n := len(src)-to+from, len(src); k < n; k++ {
				src[k] = 0 // nil for pointers
			}
			return src[:len(src)-to+from]
		}},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 4, 6)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		src      []T
		position int
		expected []T
	}{
		{
			"first",
			[]T{1, 2, 3},
			0,
			[]T{2, 3},
		},
		{
			"center",
			[]T{1, 2, 3},
			1,
			[]T{1, 3},
		},
		{
			"last",
			[]T{1, 2, 3},
			2,
			[]T{1, 2},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Delete(tc.src, tc.position))
		})
	}
}

// BenchmarkDelete/presented,_first-4         	20000000	        61.3 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/presented,_center-4        	30000000	        47.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/presented,_last-4          	30000000	        48.5 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_first-4       	20000000	        51.8 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_center-4      	30000000	        47.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_last-4        	30000000	        46.7 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_first-4          	30000000	        43.3 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_center-4         	30000000	        46.0 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_last-4           	30000000	        44.9 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc_safe,_first-4           	20000000	        73.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc_safe,_center-4          	20000000	        65.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc_safe,_last-4            	30000000	        52.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gc_safe,_first-4  	30000000	        43.8 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gc_safe,_center-4 	30000000	        44.1 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gc_safe,_last-4   	30000000	        43.9 ns/op	      80 B/op	       1 allocs/op
func BenchmarkDelete(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T, int) []T
	}{
		{"presented", Delete},
		{"alternative", func(src []T, i int) []T { return src[:i+copy(src[i:], src[i+1:])] }},
		{"unstable", func(src []T, i int) []T {
			last := len(src) - 1
			src[i] = src[last]
			return src[:last]
		}},
		{"gc safe", func(src []T, i int) []T {
			copy(src[i:], src[i+1:])
			last := len(src) - 1
			src[last] = 0 // nil for pointers
			return src[:last]
		}},
		{"unstable gc safe", func(src []T, i int) []T {
			last := len(src) - 1
			src[i] = src[last]
			src[last] = 0 // nil for pointers
			return src
		}},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name+", first", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 0)
			}
		})
		b.Run(bm.name+", center", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 5)
			}
		})
		b.Run(bm.name+", last", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 9)
			}
		})
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		name     string
		src      []T
		at, size int
		expected []T
	}{
		{
			"at the beginning",
			[]T{1, 2, 3},
			0, 2,
			[]T{0, 0, 1, 2, 3},
		},
		{
			"at the center",
			[]T{1, 2, 3},
			2, 2,
			[]T{1, 2, 0, 0, 3},
		},
		{
			"at the end",
			[]T{1, 2, 3},
			3, 2,
			[]T{1, 2, 3, 0, 0},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Expand(tc.src, tc.at, tc.size))
		})
	}
}
