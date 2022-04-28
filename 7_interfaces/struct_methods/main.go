package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type employee struct {
	id   int
	name string
}

// implements fmt.Stringer
func (e employee) String() string {
	return strconv.Itoa(e.id) + ": " + e.name
}

// implements io.Writer
func (e *employee) Write(p []byte) (n int, err error) {
	newName := string(p)
	if newName != "" {
		return len(newName), nil
	}
	return 0, errors.New("no name specified")
}

func (e *employee) setName(n string) {
	e.name = n
}

func (e employee) print() {
	fmt.Println(strconv.Itoa(e.id) + ": " + e.name)
}

func main() {
	e1 := &employee{1, "aaa"}
	fmt.Printf("%T %v\n", e1, e1)
	var w io.Writer
	w = e1
	w2 := w.(io.ReadWriter)
	fmt.Printf("%T %v\n", w2, w2)
	w3 := w.(*employee)
	fmt.Printf("%T %v\n", w3, w3)

	fmt.Println("==================")

	e2 := employee{2, "bbb"}
	var f fmt.Stringer
	f = e1
	fmt.Printf("%T %v\n", f, f)
	f = e2
	fmt.Printf("%T %v\n", f, f)
	//w = e2  // compile error

	/*fmt.Printf("%T %v\n", e1, e1)
	fmt.Println(e1.getName())
	e1.setName("bbb")
	fmt.Println(e1.getName())

	fmt.Println("==============")


	fmt.Printf("%T %v\n", e2, e2)
	fmt.Println(e2.getName())
	e2.setName("ccc")
	fmt.Println(e2.getName())*/
}
