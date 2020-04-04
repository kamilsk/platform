package runtime

import "runtime"

// Caller returns information about the current caller.
//
//  func StoreToDatabase(data Payload) error {
//  	timer := stats.NewTiming()
//  	defer timer.Send(Caller().Name)
//
//  	// do something heavy
//  }
//
// Deprecated: use go.octolab.org/runtime.Caller instead.
func Caller() CallerInfo {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return CallerInfo{f.Name(), file, line}
	/*
		Alternative solution is

		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		return CallerInfo{f.Name(), file, line}

		But it has performance issue

		BenchmarkCaller/direct_caller-4         	 3000000	       390 ns/op	       0 B/op	       0 allocs/op
		BenchmarkCaller/chain_caller-4          	 3000000	       396 ns/op	       0 B/op	       0 allocs/op
		BenchmarkCaller/lambda_caller-4         	 3000000	       524 ns/op	       0 B/op	       0 allocs/op

		BenchmarkCaller/direct_caller-4         	 2000000	       732 ns/op	     184 B/op	       2 allocs/op
		BenchmarkCaller/chain_caller-4          	 2000000	       702 ns/op	     184 B/op	       2 allocs/op
		BenchmarkCaller/lambda_caller-4         	 2000000	      1019 ns/op	     248 B/op	       3 allocs/op
	*/
}

// CallerInfo holds information about a caller.
//
// Deprecated: use go.octolab.org/runtime.CallerInfo instead.
type CallerInfo struct {
	Name string
	File string
	Line int
}
