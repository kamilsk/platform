package fn

// Repeat repeats the action the required number of times.
//
//  func FillByValue(slice []int, value int) {
//  	Repeat(func () { slice = append(slice, value) }, cap(slice) - len(slice))
//  }
//
// Deprecated: use go.octolab.org/fn.Repeat instead.
func Repeat(action func(), times int) {
	for i := 0; i < times; i++ {
		action()
	}
}
