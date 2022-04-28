package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	c := a[:]
	d := a[1:2]
	fmt.Printf("%p %p %p %p %p\n", &a, &a[0], b, c, d)
}
