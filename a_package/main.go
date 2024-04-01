package main

import (
	"fmt"
)

var score = 99.5

// You can use the global variables
// and the functions from greetings.go
// on this file since they are on the same package.
//
// But first run the other file as well before you
// run this main one as such: go run main.go greetings.go

func main() {
	sayHello("Karl")

	for _, value := range points {
		fmt.Println(value)
	}
}