package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// the receiver attaches the function to any VALUE of the receiver type
// this is called a "method"
func (s secretAgent) speak() {
	fmt.Println(s)
}

func main() {
	sa1 := secretAgent {
		person: person {
			"James",
			"Bond",
		},
		ltk: true,
	}
	
	sa1.speak()
}
