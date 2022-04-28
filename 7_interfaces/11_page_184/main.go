package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Employee struct {
	id   int
	name string
}

func (e *Employee) String() string {
	return strconv.Itoa(e.id) + ": " + e.name
}

// Comparison of interfaces
func main() {
	var x os.File
	var y os.File
	fmt.Printf("%T %v\n", y, y)
	fmt.Println(x == y)
	var z io.Writer
	fmt.Println(z)
	//fmt.Println(x == z) // compile error: mismatched types
	var e1 interface{}
	e1 = Employee{1, "e1"}
	fmt.Println(e1 == e1)

	// However, if two interface values are compared and have the same dynamic type,
	// but that type is not comparable (a slice, for instance), then the comparison fails with a panic:
	var m interface{} = []int{1, 2, 3}
	fmt.Println(m)
	fmt.Println(m == m) // panic: runtime error: comparing uncomparable type []int

	var n interface{} = make(map[string]int, 0)
	fmt.Println(n)
	//fmt.Println(n == n) // panic: runtime error: comparing uncomparable type map[string]int

	//var o interface{} = func() {}
	//fmt.Println(o == o) // panic: runtime error: comparing uncomparable type func()

	// Summary of interface comparison:
	// 1. dynamic types match
	// 2. dynamic types are comparable
	// 3. dynamic values are comparable
	// So interface values could be used as the keys of map or the operand of a switch statement.
	// when comparing interface values or aggregate types that contain interface values, we must be aware of the potential for a panic.
	// A similar risk exists when using interfaces as map keys or switch operands.
	// Only compare interface values if you are certain that they contain dynamic values of comparable types.

}
