package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64
	var europa EuropeanVelocity = 120.4 * 3.6
	var america AmericanVelocity = 130 * 2.24
	s := 123.4232523525
	a := math.Round(s)
	fmt.Println(europa, america, a)
}
