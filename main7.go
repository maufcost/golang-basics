package main

import (
	"fmt"
)

// regular vars are pass-by-value
// arrays are pass-by-reference

func updateName(n * string) {
	*x = "New Name"
}

func main() {
	
	name := "Mauricio"

	p := &name

	updateName(p)

	// value: *p
}