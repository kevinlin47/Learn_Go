package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//Make sure deck has 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	//Make sure first card is Ace of HEarts
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %v", d[0])
	}

	//Make sure last is King of Diamonds
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card of King of Diamonds, but got %v", d[len(d)-1])
	}
}
