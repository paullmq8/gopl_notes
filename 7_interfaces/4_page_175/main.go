package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// an *os.File satisfies io.Reader,Writer,Closer,and ReadWriter
// an *bytes.Buffer satisfies io.Reader, Writer, and ReadWriter, but does not satisfy Closer

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w) //*os.File
	w.Write([]byte("aaa\n"))

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) //*bytes.Buffer
	w.Write([]byte("bbb"))
	//fmt.Println(string((w.(*bytes.Buffer)).Bytes()))

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	w = rwc
	//rwc = w // compile error: io.Writer lacks Close method

}
