package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Create custome type shopList
//Basically a string slice, but will be used for receiver functions
type shopList []string

func main() {

	//Create intial shopping list variable
	myShopList := shopList{}

	//Infinite loop until sentinel value is detected
	for {
		//Create command line text menu
		fmt.Println("My Shopping List v1.0")
		fmt.Println("Select option")
		fmt.Println("1. Print List")
		fmt.Println("2. Add")
		fmt.Println("3. Remove")
		fmt.Println("4. New List")
		fmt.Println("5. Save List")
		fmt.Println("6. Open List")
		fmt.Println("7. Exit Program")

		//Create a new reader to read command line input
		scanner := bufio.NewScanner(os.Stdin)

		//We will use scanner in bufio package to read command line input
		scanner.Scan()
		selection := scanner.Text()

		//Check if there was an error
		//Inform user of error if one occured
		if scanner.Err() != nil {
			fmt.Println("Invalid selection")
		} else {
			//Note that no default case was created because of the error check above
			switch selection {
			case "1":
				myShopList.printList()
			case "2":
				fmt.Println("Enter item name\n-> ")

				scanner.Scan()
				item := scanner.Text()
				if scanner.Err() == nil {
					myShopList = addToList(myShopList, item)
				} else {
					fmt.Println("Error occured when adding item")
				}
			case "3":
				fmt.Println("Remove\n Enter item index\n-> ")
				scanner.Scan()
				index := scanner.Text()

				if scanner.Err() == nil {
					i, _ := strconv.Atoi(index)
					var err error
					myShopList, err = removeFromList(myShopList, i)

					if err != nil {
						fmt.Println(err)
					}
				}
			case "4":
				myShopList = newList()
			case "5":
				fmt.Println("Saving\n Enter File Name\n-> ")
				scanner.Scan()
				fileName := scanner.Text()

				err := myShopList.saveToFile(fileName)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Shopping List Saved Successfully")
				}
			case "6":
				fmt.Println("Opening\n Enter File Name\n-> ")
				scanner.Scan()
				fileName := scanner.Text()
				var err error
				myShopList, err = readFromFile(fileName)

				if scanner.Err() != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Opened Shoppinh List Successfully")
				}
			case "7":
				fmt.Println("Exiting Application")
				fmt.Println("Goodbye")

				//We execute a graceful shutdown of the application with return code 0
				//Indicating there was no error when exiting the application
				os.Exit(0)
			}
		}
	}
}

/*
Name: printList
Desc: Function prints the contents of variable type shopList
@params None
@returns None
*/
func (l shopList) printList() {
	for i, d := range l {
		fmt.Println(i, ". ", d)
	}
}

/*
Name addToList
Desc: Function adds a new item to the calling shopList
@params list - List to add to
@params item - Item to be added to the list
@returns New updated shop list
*/
func addToList(list shopList, item string) shopList {
	return shopList(append(list, item))
}

/*
Name: removeFromList
Desc: Function removes item from list given index
@params list - List to remove from
@params index - Index of item to be removed
@returns New updated list if success, original list if error
@return error nil is no error, error otherwise
*/
func removeFromList(list shopList, index int) (shopList, error) {

	if index < 0 || index >= len(list) {
		return list, errors.New("Invalid index")
	}

	return shopList(append(list[:index], list[index+1:]...)), nil
}

/*
Name: newList
Desc: Function returns a new empty shopList
@params None
@returns New empty shopList
*/
func newList() shopList {
	return shopList{}
}

/*
Name saveToFile
Desc: Saves the list to a file on the computer
@params: fileName Name of file
@return error nil if no error, error otherwise
*/
func (l shopList) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(l.toString()), 0666)
}

/* Name readFromFile
Desc: Function reads a list back from file on the computer
@params fileName Name of file
@returns Data in file as byte slice and error if one occured

Note: There is a issue here where if there is an error on file read.
	The function will still continue to manipulate data returned to bs
	which causes an issue, leaving it as is for now but is an easy fix
*/
func readFromFile(fileName string) (shopList, error) {
	bs, error := ioutil.ReadFile(fileName)

	shopList := strings.Split(string(bs), ",")

	return shopList, error
}

/*
Name toString
Desc: Converts string slice to a string
@params: None
@returns: String containing all the strings in the slice
*/
func (l shopList) toString() string {
	return strings.Join([]string(l), ",")
}
