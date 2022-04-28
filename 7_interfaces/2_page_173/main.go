package main

import "fmt"

type ByteCounter int

// the ByteCounter correctly accumulates the length of the result
// *ByteCounter also implements io.Writer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	// Since *ByteCounter satisfies the io.Writer contract, we can pass it to Fprintf
	var c ByteCounter
	fmt.Println(c) // "0"
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
