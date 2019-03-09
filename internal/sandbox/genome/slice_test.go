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
	gcSafe := func(src []T, from, to int) []T {
		copy(src[from:], src[to:])
		for k, n := len(src)-to+from, len(src); k < n; k++ {
			src[k] = 0 // nil for pointers
		}
		return src[:len(src)-to+from]
	}

	tests := []struct {
		name     string
		origin   []T
		from, to int
		expected []T
	}{
		{
			"head",
			[]T{1, 2, 3, 4, 5},
			0, 2,
			[]T{3, 4, 5},
		},
		{
			"center",
			[]T{1, 2, 3, 4, 5},
			1, 4,
			[]T{1, 5},
		},
		{
			"tail",
			[]T{1, 2, 3, 4, 5},
			3, 5,
			[]T{1, 2, 3},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Cut(Copy(tc.origin), tc.from, tc.to))

			assert.Equal(t, tc.expected, gcSafe(Copy(tc.origin), tc.from, tc.to))
		})
	}
}

// BenchmarkCut/presented-4         	30000000	        47.7 ns/op	      80 B/op	       1 allocs/op
// BenchmarkCut/gc-safe-4           	30000000	        49.8 ns/op	      80 B/op	       1 allocs/op
func BenchmarkCut(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T, int, int) []T
	}{
		{"presented", Cut},
		{"gc-safe", func(src []T, from, to int) []T {
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
	alternative := func(src []T, i int) []T { return src[:i+copy(src[i:], src[i+1:])] }
	gcSafe := func(src []T, i int) []T {
		copy(src[i:], src[i+1:])
		last := len(src) - 1
		src[last] = 0 // nil for pointers
		return src[:last]
	}

	tests := []struct {
		name     string
		src      []T
		position int
		expected []T
	}{
		{
			"head",
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
			"tail",
			[]T{1, 2, 3},
			2,
			[]T{1, 2},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Delete(Copy(tc.src), tc.position))

			assert.Equal(t, tc.expected, alternative(Copy(tc.src), tc.position))
			assert.Equal(t, tc.expected, gcSafe(Copy(tc.src), tc.position))
		})
	}
}

// BenchmarkDelete/presented,_head-4         	20000000	        52.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/presented,_center-4       	30000000	        48.7 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/presented,_tail-4         	30000000	        48.3 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_head-4       	30000000	        52.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_center-4     	20000000	        50.5 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/alternative,_tail-4       	30000000	        45.6 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_head-4          	30000000	        43.3 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_center-4        	30000000	        42.8 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable,_tail-4          	30000000	        43.1 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc-safe,_head-4           	20000000	        52.1 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc-safe,_center-4         	30000000	        48.5 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/gc-safe,_tail-4           	30000000	        47.3 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gcs,_head-4      	30000000	        45.4 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gcs,_center-4    	30000000	        43.4 ns/op	      80 B/op	       1 allocs/op
// BenchmarkDelete/unstable_gcs,_tail-4      	30000000	        45.4 ns/op	      80 B/op	       1 allocs/op
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
		{"gc-safe", func(src []T, i int) []T {
			copy(src[i:], src[i+1:])
			last := len(src) - 1
			src[last] = 0 // nil for pointers
			return src[:last]
		}},
		{"unstable gcs", func(src []T, i int) []T {
			last := len(src) - 1
			src[i] = src[last]
			src[last] = 0 // nil for pointers
			return src
		}},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name+", head", func(b *testing.B) {
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
		b.Run(bm.name+", tail", func(b *testing.B) {
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
			"head",
			[]T{1, 2, 3},
			0, 1,
			[]T{0, 1, 2, 3},
		},
		{
			"center",
			[]T{1, 2, 3},
			2, 1,
			[]T{1, 2, 0, 3},
		},
		{
			"tail",
			[]T{1, 2, 3},
			3, 1,
			[]T{1, 2, 3, 0},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Expand(tc.src, tc.at, tc.size))
		})
	}
}

func TestExtend(t *testing.T) {
	alternative := func(src []T, size int) []T {
		extended := make([]T, len(src)+size)
		copy(extended, src)
		return extended
	}

	tests := []struct {
		name     string
		src      []T
		size     int
		expected []T
	}{
		{"normal case", []T{3, 2, 1}, 1, []T{3, 2, 1, 0}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Extend(tc.src, tc.size))

			assert.Equal(t, tc.expected, alternative(tc.src, tc.size))
		})
	}
}

// BenchmarkExtend/presented-4         	20000000	        55.7 ns/op	      96 B/op	       1 allocs/op
// BenchmarkExtend/alternative-4       	20000000	        59.1 ns/op	      96 B/op	       1 allocs/op
func BenchmarkExtend(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T, int) []T
	}{
		{"presented", Extend},
		{"alternative", func(src []T, size int) []T {
			extended := make([]T, len(src)+size)
			copy(extended, src)
			return extended
		}},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Extend(make([]T, 1), 10)
			}
		})
	}
}

//
// Insert, InsertVector
//

func InsertV2(dst []T, t T, at int) []T {
	var def T
	dst = append(dst, def)
	copy(dst[at+1:], dst[at:])
	dst[at] = t
	return dst
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		dst  []T
		t    T
		at   int
	}{
		{"head", []T{2, 3}, 1, 0},
		{"center", []T{1, 3}, 2, 1},
		{"tail", []T{1, 2}, 3, 2},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.t, Insert(tc.dst, tc.t, tc.at)[tc.at])
			assert.Equal(t, Insert(tc.dst, tc.t, tc.at), InsertV2(tc.dst, tc.t, tc.at))
		})
	}
}

// BenchmarkInsert/v1,_head-4         	10000000	       143 ns/op	     240 B/op	       2 allocs/op
// BenchmarkInsert/v1,_center-4       	10000000	       122 ns/op	     240 B/op	       2 allocs/op
// BenchmarkInsert/v1,_tail-4         	10000000	       130 ns/op	     240 B/op	       2 allocs/op
// BenchmarkInsert/v2,_head-4         	10000000	       125 ns/op	     240 B/op	       2 allocs/op
// BenchmarkInsert/v2,_center-4       	10000000	       130 ns/op	     240 B/op	       2 allocs/op
// BenchmarkInsert/v2,_tail-4         	10000000	       138 ns/op	     240 B/op	       2 allocs/op
func BenchmarkInsert(b *testing.B) {
	benchmarks := []struct {
		name      string
		algorithm func([]T, T, int) []T
	}{
		{"v1", Insert},
		{"v2", InsertV2},
	}
	for _, bm := range benchmarks {
		bm := bm
		b.Run(bm.name+", head", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 0, 10)
			}
		})
		b.Run(bm.name+", center", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 5, 10)
			}
		})
		b.Run(bm.name+", tail", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bm.algorithm(make([]T, 10), 9, 10)
			}
		})
	}
}

func TestInsertVector(t *testing.T) {
	tests := []struct {
		name string
		dst  []T
		v    []T
		at   int
	}{
		{"head", []T{2, 3}, []T{1}, 0},
		{"center", []T{1, 3}, []T{2}, 1},
		{"tail", []T{1, 2}, []T{3}, 2},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.v, InsertVector(tc.dst, tc.v, tc.at)[tc.at:tc.at+len(tc.v)])
		})
	}
}

//
// Push, Pop
//

func TestPush(t *testing.T) {
	tt := Push([]T{1, 2}, 3)
	assert.Equal(t, []T{1, 2, 3}, tt)
}

func TestPop(t *testing.T) {
	el, tt := Pop([]T{1, 2, 3})
	assert.Equal(t, T(3), el)
	assert.Equal(t, []T{1, 2}, tt)
}

//
// Unshift, Shift
//

func TestUnshift(t *testing.T) {
	tt := Unshift([]T{2, 3}, 1)
	assert.Equal(t, []T{1, 2, 3}, tt)
}

func TestShift(t *testing.T) {
	el, tt := Shift([]T{1, 2, 3})
	assert.Equal(t, T(1), el)
	assert.Equal(t, []T{2, 3}, tt)
}

//
// Filter
//

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		src      []T
		filter   func(int, T) bool
		expected []T
	}{
		{"even key", []T{1, 2, 3, 4}, func(i int, _ T) bool { return i%2 == 0 }, []T{1, 3}},
		{"odd value", []T{1, 2, 3, 4}, func(_ int, t T) bool { return t%2 != 0 }, []T{1, 3}},
		{"combination", []T{1, 2, 3, 4}, func(i int, t T) bool { return i%2 == 0 && t%2 != 0 }, []T{1, 3}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Filter(tc.src, tc.filter))
		})
	}
	t.Run("filter by key", func(t *testing.T) {
		assert.Equal(t,
			Filter([]T{1, 2, 3, 4}, func(i int, _ T) bool { return i%2 == 0 }),
			FilterByKey([]T{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }))
	})
	t.Run("filter by value", func(t *testing.T) {
		assert.Equal(t,
			Filter([]T{1, 2, 3, 4}, func(_ int, t T) bool { return t%2 != 0 }),
			FilterByValue([]T{1, 2, 3, 4}, func(t T) bool { return t%2 != 0 }))
	})
}
