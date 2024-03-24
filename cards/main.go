package main

func main() {
	// Can also pass in variables or strings in below definition
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
