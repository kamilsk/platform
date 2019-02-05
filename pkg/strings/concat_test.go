// +build go1.10

package strings_test

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"

	. "github.com/kamilsk/platform/pkg/strings"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		expected string
	}{
		{name: "nothing to pass"},
		{"simple usage", []string{"127.0.0.1", ":", "80"}, "127.0.0.1:80"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Concat(tc.strings...))
		})
	}
}

// BenchmarkStringConcatenation/bytes-12         	10000000	       119 ns/op	     176 B/op	       3 allocs/op
// BenchmarkStringConcatenation/runes-12         	10000000	       231 ns/op	     112 B/op	       3 allocs/op
// BenchmarkStringConcatenation/concat-12        	50000000	        39.2 ns/op	      16 B/op	       1 allocs/op
// BenchmarkStringConcatenation/join-12          	30000000	        51.2 ns/op	      16 B/op	       1 allocs/op
func BenchmarkStringConcatenation(b *testing.B) {
	b.Run("bytes", func(b *testing.B) {
		var result string
		concat := func(strings ...string) string {
			buf := bytes.NewBuffer(nil)
			for _, str := range strings {
				buf.WriteString(str)
			}
			return buf.String()
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = concat("127.0.0.1", ":", "80")
		}
		assert.Equal(b, "127.0.0.1:80", result)
	})
	b.Run("runes", func(b *testing.B) {
		var result string
		concat := func(strings ...string) string {
			var ln int
			for _, str := range strings {
				ln += utf8.RuneCountInString(str)
			}
			pos, rr := 0, make([]rune, ln)
			for _, str := range strings {
				ln = utf8.RuneCountInString(str)
				copy(rr[pos:pos+ln], []rune(str))
				pos += ln
			}
			return string(rr)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = concat("127.0.0.1", ":", "80")
		}
		assert.Equal(b, "127.0.0.1:80", result)
	})

	b.Run("concat", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Concat("127.0.0.1", ":", "80")
		}
	})
	b.Run("join", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			strings.Join([]string{"127.0.0.1", ":", "80"}, "")
		}
	})
}
