package main

import "fmt"

func main() {
	// Can also pass in variables or strings in below definition
	cards := []string{newCard(), newCard()}
	// creates new slice and reassigns variable to that
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(card, i)
	}

	fmt.Println(cards)
}

// func newCard() (card string) {
func newCard() string {
	return "Five of Diamonds"
}
