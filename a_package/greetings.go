package main

import (
	"fmt"
)

// We can use global vars from main.go here as well
// By global, I mean "the package scope" i.e. outside of main.

// You can't use := here (outside functions).
var points = []int{ 20, 90, 100, 45, 70 }

func sayHello(n string) {
	fmt.Println("hello, ", n)
}

func showScore() {
	fmt.Println("You scored: ", points)
}