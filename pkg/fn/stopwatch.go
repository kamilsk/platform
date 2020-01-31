package fn

import "time"

// Stopwatch calculates the fn execution time.
//
//  var result interface{}
//
//  duration := Stopwatch(func() { result = do.some("heavy") })
//
// Deprecated: use go.octolab.org/fn.Stopwatch instead.
func Stopwatch(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
