package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)

	// *bytes.Buffer must satisfy io.Writer
	// The declaration below asserts at compile time that a value of type *bytes.Buffer satisfies io.Writer:
	var w io.Writer = new(bytes.Buffer)
	fmt.Printf("%T %v\n", w, w)
	var a bytes.Buffer
	fmt.Printf("%T %v\n", a, a)
	//var _ io.Writer = a // compile error
	var _ io.Writer = (*bytes.Buffer)(nil)
}
