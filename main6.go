package main

import (
	"fmt"
)

// Maps
// keys must be of the same type
// values must be of the same type

func main() {
	
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"coffee": 3.59,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, v);
	}

	phonebook := map[int]string{
		4077284444: "mario",
		4108291800: "sarah",
		4128921112: "carlson",
	}

	fmt.Println(phonebook[4077284444])

	phonebook[4077284444] = "Jess"

	fmt.Println(phonebook)
}