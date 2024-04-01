package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	// Strings package
	greeting := "Hello, beautiful!"

	fmt.Println(strings.Contains(greeting, "Hello"))

	newStr := strings.ReplaceAll(greeting, "Hello", "Hi")

	// .ToUpper(params)
	// .Index(params)
	// .Split(params)

	// Sort package
	ages := []int{3, 1, 31, 26, 89, 32, 4, 1, 25}
	sort.Ints(ages)

	index := sort.SearchInts(age, 25)

	// sort.Strings(params)
	// sort.SearchStrings(params)
}