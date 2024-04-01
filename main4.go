package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println("value of x: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of y is: ", i)
	}

	names := []string{"mark", "moris", "majay"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// This is like a forEach (or enumerate())
	for index, value := range names {
		fmt.Println("The value at index %v is %v\n", index, value)
	}
}