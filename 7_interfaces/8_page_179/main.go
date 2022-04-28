package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "Sleep period")

// used flag.Value interface
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	//var a flag.Value
	//var _ fmt.Stringer = a
	//fmt.Printf("%T %[1]v\n", a)
}
