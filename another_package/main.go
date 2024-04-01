package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func getInput() {
	
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("New bill name: ")

	// Reads everything up until we click enter (hence: \n)
	name, _ = reader.readString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Print("Bill created successfully")
}

func main() {
	
	mybill = newBill("Mary's bill")

	mybill.updateTip(10)
	mybill.addItem("Apple Pie", 5.94)

	fmt.Println(mybill)
	fmt.Println(mybill.format())
}