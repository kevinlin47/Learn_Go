package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	kevin := person{
		firstName: "Kevin",
		lastName:  "Lin",
		contactInfo: contactInfo{
			email:   "example@example.com",
			zipCode: 1234,
		},
	}
	kevinPtr := &kevin
	kevinPtr.updateEmail("myNewEmail@example.com")
	kevin.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateEmail(newEmail string) {
	(*p).contactInfo.email = newEmail
}
