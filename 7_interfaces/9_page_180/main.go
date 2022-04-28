package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

// It's easy to define new flag notations for our own data types.
// We need only define a type that satisfies the flag.Value interface

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

// -temp 50C
// -help
func main() {
	var cf celsiusFlag
	cf.Set("100°C")
	fmt.Println(cf.String())
	flag.Parse()
	fmt.Println(*temp)
}
