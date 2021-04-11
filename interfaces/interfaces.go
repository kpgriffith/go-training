package main

import (
	"fmt"
)

// an interface allows us to define behavior

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// the receiver attaches the function to any VALUE of the receiver type
// this is called a "method"
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - the person speak")
}

// the interface defining the behavior
type human interface {
	speak()
}

// any other type that has the method speak() is also of type human
// a value can be of more than one type

// to use this we can do something like
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I'm a humannnnnnn", h.(person).first) // assertion example
	case secretAgent:
		fmt.Println("I'm a humannnnnnn", h.(secretAgent).first) // assertion example
	}
	fmt.Println("I'm a human", h)
}

// conversion example
type hotdog int

func main() {
	// this is both a secretAgent type and of type human, because it implements the speak() method
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	// converstion example
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) // conversion is here
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
