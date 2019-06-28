package main

import (
	"errors"
	"fmt"
)

type myCustomerError string

func main() {
	result, err := addTwoPositiveNumbers(3, 19)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	result, err = addTwoPositiveNumbers(-1, 10)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	newErr := errors.New("Failed to find repository")

	if newErr != nil {
		fmt.Println(newErr)
	}
}

//Function returns the addition of two positive numbers
func addTwoPositiveNumbers(numOne int32, numTwo int32) (int32, error) {
	if numOne < 0 || numTwo < 0 {
		var err myCustomerError = "Error: Invalid arguments"
		return 0, err
	}

	return numOne + numTwo, nil
}

func (err myCustomerError) Error() string {
	return string(err)
}
