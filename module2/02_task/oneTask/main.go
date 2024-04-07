package main

import "fmt"

func main() {
	var a *int
	b := 112
	a = &b
	fmt.Println(*a)
	*a = 101
	fmt.Println(b)
}
