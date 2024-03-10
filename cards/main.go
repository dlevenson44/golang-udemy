package main

func main() {
	// Can also pass in variables or strings in below definition
	cards := deck{newCard(), newCard()}
	// creates new slice and reassigns variable to that
	cards = append(cards, "Six of Spades")

	cards.print()
}

// func newCard() (card string) {
func newCard() string {
	return "Five of Diamonds"
}
