package main

import "fmt"

// Fprintf accepts an io.Writer interface
func main() {

	// Fprintf may not assume that it is writing to a file or to memory, only that it can call Write.
	// Because fmt.Fprintf assumes nothing about the representation of the value
	// and relies only on the behaviors guaranteed by the io.Writer contract,
	// we can safely pass a value of any concrete type that satisfies io.Writer as the first argument to fmt.Fprintf.

	// func Printf(format string, args ...interface{}) (int, error) {
	//         return Fprintf(os.Stdout, format, args...)
	// }
	fmt.Printf("%s\n", "sss")

	// func Sprintf(format string, args ...interface{}) string {
	//         var buf bytes.Buffer
	//         Fprintf(&buf, format, args...)
	//         return buf.String()
	// }
	var s = fmt.Sprintf("%s\n", "sss")
	fmt.Println(s)
}
