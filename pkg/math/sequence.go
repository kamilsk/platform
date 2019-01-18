package math

// Sequence returns an empty slice with the specified size.
//
//     for range Sequence(5) {
//             // do something five times
//     }
//
func Sequence(size int) []struct{} {
	return make([]struct{}, size)
}
