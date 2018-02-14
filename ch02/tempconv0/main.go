package main

import (
	"fmt"
)

// Celcius is type
type Celcius float64

// Fahrenheit type
type Fahrenheit float64

const (
	// AbsoluteZeroC is absolute 0 in Celcius
	AbsoluteZeroC Celcius = -273.15
	// FreezingC is freezing temperature in Celcius
	FreezingC Celcius = 0
	//BoilingC is boiling temperature in Celcius
	BoilingC Celcius = 100
)

// CToF converts Celcius to Fahrenheit
func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts Fahrenheit to Celcius
func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
}
