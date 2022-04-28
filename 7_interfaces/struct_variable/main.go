package main

import "fmt"

type employee struct {
	id   int
	name string
}

func (e employee) getName() string {
	return e.name
}

func (e *employee) setName(n string) {
	e.name = n
}

func main() {
	e1 := &employee{1, "aaa"}
	fmt.Printf("%T %v\n", e1, e1)
	fmt.Println(e1.getName())
	e1.setName("bbb")
	fmt.Println(e1.getName())

	fmt.Println("==============")

	e2 := employee{2, "bbb"}
	fmt.Printf("%T %v\n", e2, e2)
	fmt.Println(e2.getName())
	e2.setName("ccc")
	fmt.Println(e2.getName())
}
