package runtime_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/platform/pkg/runtime"
)

func TestCaller(t *testing.T) {
	ahead := func(t *testing.T, current GoVersion, target struct {
		version GoVersion
		release string
	}) bool {
		if current.Equal(target.version) {
			return true
		}
		if !unstable(current.Raw) {
			return current.Later(target.version)
		}
		prefix := "devel +61170f85e6 "
		layout := "Mon Jan 2 15:04:05 2006 -0700"
		release, _ := time.Parse(layout, target.release)
		control, _ := time.Parse(layout, current.Raw[len(prefix):])
		t.Log(target.release, "->", release, "<->", control, "<-", current.Raw[len(prefix):])
		return control.After(release)
	}

	tests := []struct {
		name     string
		caller   func() CallerInfo
		expected string
	}{
		{"direct caller", callerA, "github.com/kamilsk/platform/pkg/runtime_test.callerA"},
		{"chain caller", callerB, "github.com/kamilsk/platform/pkg/runtime_test.callerA"},
		{"lambda caller", callerC, func() string {
			if ahead(t, Version(), go112) {
				// https://golang.org/doc/go1.12#runtime
				return "github.com/kamilsk/platform/pkg/runtime_test.callerC"
			}
			return "github.com/kamilsk/platform/pkg/runtime_test.callerC.func1"
		}()},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.caller().Name)
		})
	}
}

// BenchmarkCaller/direct_caller-12         	 5000000	       288 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/chain_caller-12          	 5000000	       298 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/lambda_caller-12         	 5000000	       297 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCaller(b *testing.B) {
	benchmarks := []struct {
		name   string
		caller func() CallerInfo
	}{
		{"direct caller", callerA},
		{"chain caller", callerB},
		{"lambda caller", callerC},
	}
	for _, bm := range benchmarks {
		tc := bm
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = tc.caller()
			}
		})
	}
}

//go:noinline
func callerA() CallerInfo {
	return Caller()
}

func callerB() CallerInfo {
	return callerA()
}

func callerC() CallerInfo {
	return func() CallerInfo {
		return Caller()
	}()
}
