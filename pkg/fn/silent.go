package fn

// PrintSilent accepts a result of fmt.Fprint* function family
// and allow you to ignore it.
//
//  PrintSilent(fmt.Fprintln(writer, "ignore the result"))
//
func PrintSilent(int, error) {}
