package main

import (
	"fmt"
)

type IntSet struct {
}

func (*IntSet) String() string {
	return "string"
}

func (IntSet) Aaa() string {
	return "aaa"
}

func main() {

	//var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	var a = (&IntSet{}).String()
	var b = IntSet{}.Aaa()
	var c = (&IntSet{}).Aaa()
	fmt.Println("1: ", a, b, c)

	var s IntSet
	fmt.Printf("%v\n", s)
	// It is legal to call a *T method on an argument of type T as long as the argument is a variable
	// This is mere syntactic sugar
	var a2 = s.String() // OK, s is a variable and &s has a String method.
	var b2 = (&s).String()
	var c2 = (&s).Aaa()
	var d2 = s.Aaa()
	fmt.Println("2: ", a2, b2, c2, d2)

	var s2 = &s
	var a3 = s2.String()
	var b3 = (*s2).String()
	var c3 = s2.Aaa()
	var d3 = (*s2).Aaa()
	fmt.Println("3: ", a3, b3, c3, d3)

	// But a value of type T doesn't possess all the methods that a *T pointer does
	// as a result it might satisfy fewer interfaces.
	var _ fmt.Stringer = &s // OK
	//var _ fmt.Stringer = s  // compile error: IntSet lacks String method
}
