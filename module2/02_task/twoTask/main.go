package main

import (
	"fmt"
	"math"
)

func main() {
	var R *float64
	*R = float64(35 / math.Pi * 2)
	S := math.Pi * *R * *R
	fmt.Print(S)
}
