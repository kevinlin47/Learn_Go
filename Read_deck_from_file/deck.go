package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type deck []string

func main() {
	cards := newDeck()

	fmt.Println(reflect.TypeOf(cards))

	cards.print()

	fmt.Println("\n", len(cards))

	bigString := cards.toString()

	fmt.Println(bigString)

	if cards.saveToFile("test_file") == nil {
		fmt.Println("Deck saved successfully")
	} else {
		fmt.Println("Error in saving deck")
	}
}

func newDeck() deck {

	newDeck := []string{}

	cardSuits := [4]string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for i := 0; i < len(cardSuits); i++ {
		for j := 0; j < len(cardValues); j++ {
			newDeck = append(newDeck, cardValues[j]+" of "+cardSuits[i])
		}
	}
	return newDeck
}

func (d deck) print() {
	for i, c := range d {
		if i%13 == 0 {
			fmt.Println()
		}
		fmt.Printf("%s ", c)
	}
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
