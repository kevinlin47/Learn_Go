package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberOne = 3            //integer variable
	var numberTwo float32 = 3.19 //float variable

	// Can't add numberOne with numberTwo because of a mismatch in variable type
	// We can cast numberOne to a float

	numberOneAsFloat := float32(numberOne)

	fmt.Println(numberOneAsFloat + numberTwo)

	numberOne = 5
	numberTwo = 2

	fmt.Println(5 / 2)
	fmt.Println(float32(numberOne) / float32(numberTwo))

	//convert float32 to float64
	var myFloat32 float32 = 19.95
	fmt.Println(myFloat32)

	myFloat64 := float64(myFloat32)

	fmt.Println(myFloat64)

	//How to change a number represented as a string to an integer and vice versa?
	myInt := 7
	myString := "5"

	//Can't add these because they are not the same type
	stringAsInt, err := strconv.Atoi(myString)

	if err == nil {
		fmt.Println(myInt + stringAsInt)
	}

	//Convert integer to string
	intAsString := strconv.Itoa(myInt)

	fmt.Println(intAsString + myString) //Should print concatenation of "7" and "5", so string "75"
}
