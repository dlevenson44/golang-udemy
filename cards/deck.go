package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// not receiver func
// takes d of type deck and hs type int
// returns two values both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver func
// parenthesis are receiver, this is a receiver function
// any variable of type "deck" can access this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Type Conversion:
// takes type deck and converts it to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// anyone can read/write file with 0666 permission
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
