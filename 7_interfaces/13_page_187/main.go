package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int { return len(p) }

func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// To sort any sequence, we need to define a type that implements these three methods,
// then apply sort.Sort to an instance of that type.
func main() {
	var names []string
	names = append(names, "c")
	names = append(names, "b")
	names = append(names, "a")
	fmt.Println(names, len(names), cap(names))
	fmt.Printf("%p\n", names)
	// The conversion yields a slice value with the same length, capacity,
	// and underlying array as names but with a type that has the three methods required for sorting.
	var stringSlice = StringSlice(names)
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice))
	fmt.Printf("%p\n", stringSlice)
	sort.Sort(stringSlice)
	fmt.Println(stringSlice)
	fmt.Println(names) // caution: original slice names is also changed instead of returning a new slice
	//fmt.Println("1000" < "001")
}
