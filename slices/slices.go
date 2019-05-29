package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5} //Create an array of of type int that holds 5 elements.

	for i, d := range array {
		fmt.Printf("array[%d] = %d \n", i, d)
	}

	fmt.Println(len(array)) // We can use the built in len function to get the length of an array.

	anotherArray := [10]string{} //Create an array that holds 10 string elements
	// We do not need to have initial values, if no initial values are given,
	//Go will default all the entries to the elements respective zero value.

	fmt.Println(anotherArray, len(anotherArray))

	//You can also declare and initialize an array and have Go count the number of elements for you
	myArray := [...]string{"Go", "is", "Awesome"}

	fmt.Println(myArray)
	fmt.Println(len(myArray))
}
