package main

import (
	"fmt"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// Make new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

// Receiver function to format the bill
// It is associated with bill objects
// --> That's how we associate this function: (b bill)
func (b * bill) format() string {
	// Formatted string
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Printf("%v cost $%0.2f", k, v)
		total += v
	}

	fs += fmt.Printf("Tip: %v", b.tip)
	fs += fmt.Printf("Total: %v", total + tip)

	return fs
}

// This is a copy of the bill (the 'b')
// So, we need to pass a pointer to the original bill.
//
// Another advantage to using pointers with receiving
// functions: to not create multiple copies of the same
// bill.
func (b * bill) updateTip(tip float64) {
	(*b).tip = tip
	// structs automatically dereference, so we could've done:
	// b.tip = tip
}

func (b * bill) addItem(name string, price float64) {
	b.items[name] = price;
}