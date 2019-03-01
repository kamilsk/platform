package fn

// PrintSilent accepts a result of fmt.*Print* function family
// and allow you to ignore it.
//
//  PrintSilent(fmt.Println("ignore the result"))
//
func PrintSilent(int, error) {}
