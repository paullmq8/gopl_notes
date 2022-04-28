package main

import "fmt"

// Duck interface definition:
// duck typing: If it quacks like a duck and swims like a duck, then it is a duck.
type Duck interface {
	Quack() string
	Swim()
}
type RealDuck struct{}

func (rd *RealDuck) Quack() string { return "Quack Quack Quack..." }
func (rd *RealDuck) Swim()         { fmt.Println("Swimming in the pool") }

type ToyDuck struct{}

func (rd *ToyDuck) Quack() string { return "Quik Quik Quik..." }
func (rd *ToyDuck) Swim()         { fmt.Println("Swimming in bath") }

// DuckSwim shows the purpose of interface: polymorphism
func DuckSwim(duck Duck) {
	duck.Swim()
}

// ByteCounter synonym
type ByteCounter int

func (bc *ByteCounter) String() string {
	return fmt.Sprintf("ByteCounter: %d", *bc)
}

func main() {
	var a int
	a = 3
	bc := ByteCounter(a)
	fmt.Println(bc.String())
	var w fmt.Stringer
	w = &bc
	fmt.Println(w)
}
