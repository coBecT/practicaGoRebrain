package main

import (
	"fmt"
	"math"
)

func main() {
	var R *float64
	const p = 3.14
	var length float64
	fmt.Scanf("%f", &length)

	rad := length / (2 * p)
	fmt.Println(rad)
	R = &rad
	fmt.Printf("R = %f\n", *R)

	S := 2 * p * math.Pow(*R, 2)
	fmt.Printf("s = %f", S)
}
