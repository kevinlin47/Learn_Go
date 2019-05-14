package main

import "fmt"

func main() {
	myCard := getNewCard()

	fmt.Println(myCard)

	r := 5.0
	areaOfCircle := calcAreaOfCircle(r)

	fmt.Println(areaOfCircle)

	f, c, k := getTempReading()

	fmt.Println(f, c, k)
}

func getNewCard() string {
	return "Ace of Spades"
}

func calcAreaOfCircle(r float64) float64 {
	return 3.14 * (r * r)
}

func getTempReading() (float32, float32, float32) {
	return 32, 0, 273.15
}
