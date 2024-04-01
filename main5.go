package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreetingTo(n string) {
	fmt.Printf("Hello, %v! \n", n);
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v);
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	age := 21

	if age < 30 {
		sayGreetingTo("Peralta")
	} else if age < 40 {
		fmt.Println("Ok...")
	} else {
		fmt.Println("Jurassic Park")
	}

	cycleNames([]string{"Marcus", "Joseph", "Carly"})
	firstLetter, lastLetter = getInitials("Mauricio Costa")
}