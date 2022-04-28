package main

import (
	"fmt"
	"io"
	"os"
)

// A nil interface value, which contains no value at all, is not the same
// as an interface value containing a pointer that happens to be nil.
const debug = true

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ... do something ...
	fmt.Printf("%T %[1]v\n", out)
	if out != nil {
		out.Write([]byte("done!\n"))
		fmt.Println("Written to buffer")
	}
}

func main() {
	var buf *os.File // io.Writer, *os.File, *bytes.Buffer
	fmt.Printf("%T %[1]v\n", buf)
	if buf == nil {
		fmt.Println("buf is nil")
	}
	if debug {
		buf = new(os.File) // new(os.File) new(bytes.Buffer) // enable collection of output
	}
	fmt.Printf("%T %[1]v\n", buf)
	f(buf) // will panic if debug is set to false
}
