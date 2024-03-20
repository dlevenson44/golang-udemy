package main

func main() {
	// Can also pass in variables or strings in below definition
	cards := newDeck()
	cards.saveToFile("my_cards")
}
