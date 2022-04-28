package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

// the ByteCounter correctly accumulates the length of the result
// *ByteCounter also implements io.Writer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// Type Assertion
func main() {
	// 1. The type assertion is for interface variable.

	// 2.Type Assertion has the following two cases:
	// If the asserted type T is a concrete type, then the type assertion checks whether x’s dynamic type is identical to T.
	// If this check succeeds, the result of the type assertion is x’s dynamic value, whose type is of course T.
	// case 1: var.(concrete type)
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Printf("%T %v\n", c, c)

	// If the asserted type T is an interface type, then the type assertion checks whether x’s dynamic type satisfies/implements T.
	// In other words, a type assertion to an interface type changes the type of the expression,
	// making a different (and usually larger) set of methods accessible, but it preserves the dynamic type and value components inside the interface value.
	// case 2: var.(interface)
	_ = w.(io.ReadWriter) // success: *os.File has both Read and Write
	//w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method

	// 3. No matter what type was asserted, if the operand is a nil interface value, the type assertion fails.
	var w2 io.ReadWriter
	fmt.Printf("%T %v\n", w2, w2) // why is the type nil?
	//rw2 := w2.(io.Writer)         // panic
	//fmt.Println(rw2)

	// 4. often we use two variables to receive the results to avoid panic
	var w3 io.Writer = os.Stdout
	f, ok := w3.(*os.File) // success:  ok, f == os.Stdout
	if !ok {
		fmt.Println("f is not a *os.File")
	} else {
		fmt.Println("f is ", f)
	}
	b, ok := w3.(*bytes.Buffer) // failure: !ok, b == nil
	if !ok {
		fmt.Println("b is not a *bytes.Buffer")
	} else {
		fmt.Println("b is ", b)
	}

	// 5. compact version:
	if rwc, ok := w3.(io.ReadWriteCloser); ok {
		fmt.Println("rwc is ", rwc)
	} else {
		fmt.Println("rwc is not a io.ReadWriteCloser")
	}

	// 6. another compact version and shadowing the original variable:
	fmt.Printf("%T %v\n", w3, w3)
	if w3, ok := w3.(io.ReadWriteCloser); ok {
		//w3.Close()
		fmt.Printf("%T %v\n", w3, w3)
	} else {
		fmt.Println("w3 is not a io.ReadWriteCloser")
	}
	fmt.Printf("%T %v\n", w3, w3)
}
