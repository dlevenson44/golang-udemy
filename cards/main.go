package main

func main() {
	// Can also pass in variables or strings in below definition
	cards := newDeck()
	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
