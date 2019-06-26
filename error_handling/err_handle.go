package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Failed to find repository")

	if err != nil {
		fmt.Println(err)
	}
}
