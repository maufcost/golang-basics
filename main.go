package main 

import "fmt"

func main() {
	fmt.Println("Hello, all!")

	var firstName string = "Mauricio"
	var lastName = "Costa"
	var nameThree

	fmt.Println(firstName, lastName, nameThree)

	// Both below are the same thing.
	nameThree = "Peachtree St."

	// Same declaration in all three assignments below.
	var age1 int = 18
	var age2 = 24
	age3 := 25

	fmt.Println(age2, "My age: ", age3);

	// ints + bits & memory
	// var num int8 = 215 <-- not allowed since int8 can only go up to 127.

	// Other int "sub-types":
	// int
	// uint
	// int8, int16, etc

	// Floats
	var mass float32 = 12.58
	weight := 25.43

	// Printfs
	fmt.Printf("this is the number: %v %v \n", mass, weight)
	fmt.Printf("mass is of type: %T \n", mass)
	fmt.Printf("mass: %0.2f \n", mass)

	// Sprintf
	var str = fmt.Sprintf("This is kinda cool: %v", firstName)
	fmt.Println("The saved string is: ", str)
}