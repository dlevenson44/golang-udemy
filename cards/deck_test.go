package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		// %d is saying we're passing an extra value, then looks at len(d)
		t.Errorf("Expected deck length of 52 but instead got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		// %v is the same as above but used for strings
		t.Errorf("Expected first card of Ace of Spades but instead got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs but instead got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52 but instead got %d", len(d))
	}

	os.Remove("_decktesting")
}
