package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// interface variable has dynamic type and dynamic value
func main() {
	// type: nil, value: nil
	var w io.Writer
	//w.Write([]byte("hello")) // panic: nil pointer dereference

	// type: *os.File, value: holds a pointer to the os.File variable representing the standard output of the process
	w = os.Stdout // assign a value of type *os.File to w. This involves an implicit conversion from a concrete type to an interface type.
	// It's equivalent to the explicit conversion io.Writer(os.Stdout)
	w.Write([]byte("hello\n")) // causes (*os.File).Write([]byte("hello")), which prints hello in console

	// type: *bytes.Buffer, value: a pointer to the newly allocated buffer
	w = new(bytes.Buffer)
	fmt.Printf("%T %#v\n", w, w)
	w.Write([]byte("Hello\n"))
	fmt.Println(string((w.(*bytes.Buffer)).Bytes()))

	// type: nil, value: nil
	w = nil
	fmt.Printf("%T %#v\n", w, w)

}
