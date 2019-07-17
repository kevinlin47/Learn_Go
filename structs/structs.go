package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int32
}

func main() {
	//This syntax creates a new struct where each field must be in order
	//as defined in the person struct, ie first field is firstName, second field is lastName
	//, and third field is age
	person1 := person{"Bruce", "Wayne", 30}

	//You can name the fields when assigning values to them, here you can order it anyway
	person2 := person{firstName: "Kevin", lastName: "Lin", age: 24}

	//Omitting any fields results in that variable types zero-value
	person3 := person{}

	fmt.Println("person1's first name: " + person1.firstName)
	fmt.Println("Person2's last name: " + person2.lastName)
	fmt.Println("person3's age : ", person3.age)
}
