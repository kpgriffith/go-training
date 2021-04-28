package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "kevin",
		last:  "griffith",
		age:   42,
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).first = "James"
	(*p).last = "Bond"
	p.age = 32

	/*
		p.f is the shorthand for (*p).f
	*/
}
