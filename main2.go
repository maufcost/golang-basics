package main

import "fmt"

// Arrays: fixed length
// Slices: variable length

func main() {
	// Fixed length!
	var ages [3]int = [3]int{18, 24, 25}
	ages[0] = 19

	// A shorthand:
	var ages2 = [3]int{17, 23, 24}

	fmt.Println(ages, len(ages))

	// Another way to assign 
	names := [3]string{"Zex", "Jax", "Mox"}

	// Slices - they use arrays under the hood.
	var scores = []int{100, 50, 80}
	scores[1] = 90

	// Returns the new array (the new "slice")
	scores = append(scores, 20)

	// Slice Ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:2]

	fmt.Println(rangeOne)

	append(rangeOne, "Koala")
}