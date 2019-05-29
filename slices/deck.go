package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards)
	fmt.Println(len(cards))
}

func newDeck() []string {

	newDeck := make([]string, 52)

	cardSuits := [5]string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for i := 0; i < len(cardSuits); i++ {
		for j := 0; j < len(cardValues); j++ {
			newDeck = append(newDeck, cardValues[j]+" of "+cardSuits[i])
		}
	}
	return newDeck
}
