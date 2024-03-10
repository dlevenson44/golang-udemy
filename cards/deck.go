package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string

// parenthesis are receiver, this is a receiver function
// any variable of type "deck" can access this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
