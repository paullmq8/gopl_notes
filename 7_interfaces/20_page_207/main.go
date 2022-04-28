package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)
	// Output:
	// &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}

	fmt.Println(os.IsNotExist(err))

	/*func IsNotExist(err error) bool {
		if pe, ok := err.(*PathError); ok {
			err = pe.Err
		}
		return err == syscall.ENOENT || err == ErrNotExist
	}*/

	/*
		type switch statement, cases are considered in order:
		The position of the default case relative to the others is immaterial. No fallthrough is allowed.
		switch x.(type) {
		     case nil:       // ...
		     case int, uint: // ...
			 case bool:
			 case string:
			 default:
		}

		switch x := x.(type) {
			case nil:       // ...
			case int, uint: // ...
			case bool:
			case string:
			default:
		}

	*/

}
