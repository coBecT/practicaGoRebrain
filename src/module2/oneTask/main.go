package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	str := "104"
	i := 35
	a := strconv.Itoa(i)
	//b := strconv.Atoi(str)
	c, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	d := 10
	d += c
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(c), d)
}
