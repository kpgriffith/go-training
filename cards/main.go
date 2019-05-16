package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zip int
}

func main() {
//	cards := newDeck()
//	cards.shuffle()
//	cards.print()
	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zip: 12345,
		},
	}

	jim.print()
	jim.updateName("Kevin")
	jim.print()
}

func (p person) print(){
	fmt.Printf("%v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
	p.print()
}
