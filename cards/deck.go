package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	// bs = byteslice, err = error
	// both returned from os.ReadFile
	// err object or nil, if nil then no error
	bs, err := os.ReadFile(filename)
	if err != nil {
		// option 1: could log error and return call to new deck
		// option 2: could log error and quit the program
		// we do option 2 here
		fmt.Println("Error in newDeckFromFile:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn((len(d) - 1))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
