package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	/* one way to declare a struct */
	// kevin := person{firstName: "Kevin", lastName: "Griffith"}
	// fmt.Println(kevin)

	/* different way to declare and populate a struct */
	// var kevin person

	// kevin.firstName = "Kevin"
	// kevin.lastName = "Griffith"

	// fmt.Println(kevin)
	// fmt.Printf("%+v", kevin)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)
}
