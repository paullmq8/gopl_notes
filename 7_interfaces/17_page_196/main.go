package main

import (
	"errors"
	"fmt"
	"io"
	"syscall"
)

// error interface
func main() {
	a := errors.New("error msg")
	fmt.Println(a)

	// func New(text string) error { return &errorString{text} }
	fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"

	b := fmt.Errorf("not found %d", 404)
	fmt.Println(b)

	// We would not want a distinguished error such as io.EOF to compare equal to one that merely happened to have the same message.
	fmt.Println(io.EOF == errors.New("EOF")) // false

	/*var errors = [...]string{
		1:   "operation not permitted",   // EPERM
		2:   "no such file or directory", // ENOENT
		3:   "no such process",           // ESRCH
		// ...
	}*/
	// type: syscall.Errno, value: 2
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"
}
