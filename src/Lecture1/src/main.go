package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("Hello Polzavatel !")
	fmt.Print("Seishas: ")
	t := time.Date(time.Now().Year(), time.March, time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Local().Second(), time.UTC)
	fmt.Println(t)
	file, err := os.Open("C:/Users/dimad/Desktop/golangPracticum/stepik_GO-BASIC/1_3_1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(file)
}
