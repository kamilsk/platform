package errors

// DoSilent accepts a result of
// * fmt.Fprint* function family
// * io.Copy* and io.Read* function family
// * io.Writer interface
// and allow you to ignore it.
//
//  DoSilent(fmt.Fprintln(writer, "ignore the result"))
//
func DoSilent(interface{}, error) {}
