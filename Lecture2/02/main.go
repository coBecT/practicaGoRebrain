package main

import "fmt"

func main() {
	var A *int
	B := 2134
	A = &B
	fmt.Println(*A)
	*A = 2268
	fmt.Println(B)
}
