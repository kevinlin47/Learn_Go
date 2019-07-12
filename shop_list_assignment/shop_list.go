package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

//Create custome type shopList
//Basically a string slice, but will be used for receiver functions
type shopList []string

func main() {

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
