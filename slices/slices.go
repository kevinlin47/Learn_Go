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

	mySlice := []float32{} //Declare and initialize a slice variable
	fmt.Println(mySlice)

	myOtherSlice := make([]float32, 2) //Declare a slice variable using built in function make
	fmt.Println(myOtherSlice)

	mySlice = append(mySlice, 3.14, 3.199) //Append new element to end of slice, append function always returns a new slice
	fmt.Println(mySlice)

	myOtherSlice[0] = 13.37 //Reassign elemet at an index to a new element
	myOtherSlice[1] = 319.95
	fmt.Println(myOtherSlice)

	copy(mySlice, myOtherSlice) //Copy one slice elements into another slice, using built in copy function

	fmt.Println(mySlice)

	/********* Examples using the slice operator *********/
	newSlice := make([]int, 20)

	for i := 0; i < len(newSlice); i++ {
		newSlice[i] = i * 2
	}

	fmt.Println(newSlice)

	newSlice2 := newSlice[0:10] //Create a new slice with elements from newSlice up to index 9

	fmt.Println(newSlice2)

	newSlice3 := newSlice[15:] //Create new slice with elements from end of newSlice up to index 15

	fmt.Println(newSlice3)

	//Example of removing ane element from a slice at a given index
	//We will remove 18 from the slice newSlice
	//The value of 18 is at index 9
	newSlice = append(newSlice[:9], newSlice[10:]...)

	//Explanation for the algorithm above to remove an element at a certain index
	//The append function takes a source slice and adds the given elements to a new slice that is returned
	//So as the source slice we use a slice containing all the elements up to the index of the element that we want to remove
	//Using the slice operator, remember the value on the right side of the ":" is inclusive.
	//Then we add onto that slice the remaning elements after the idnex of the element that we wanted is removed.
	//So using the slice operator we want elements indexOfElementRemoved + 1 to the end.
	//The ... is used for variadic parameters, in simple terms it will unpack all the elements from that slice.
	//It is shorthand for writing append(newSlice[:9], newSlice[10], newSlice[11], newSlice[12], and so on)

	fmt.Println(newSlice)
}
