package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	foo()
	bar("Kevin")
	s := woo("Paul")
	fmt.Println(s)
	a, b := ohyeah("Ellie", "Bear")
	fmt.Println(a)
	fmt.Println(b)
	variadicParams(1,2,3,4,5,6,7,8,9)
}

// func (r receiver) identifier(parameter(s)) (return(s)) {... code ...}

// receiver ties the function to a type

// parameters vs. arguments
// define the func with parameters
// when a func is called you pass arguments

func foo() {
	fmt.Println("Does something in foo...")
}

// everything in go is pass by value

func bar(s string) {
	fmt.Printf("Hello %s, from bar...\n", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func ohyeah(fn string, ln string) (string, bool) {
	s := fmt.Sprint(fn, " ", ln, ` says "Hello"`)
	b := true
	return s, b
}

func variadicParams(x ...int) { // turns x into a slice
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	
	fmt.Println(sum)
}
