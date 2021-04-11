package main

import (
	"fmt"
)

/*
Assign a function to a variable
*/
func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar() // x is assigned to a function now
	fmt.Printf("%T\n", x)

	// to get the value from the return of the function returned, we need to execute it
	fmt.Printf("%d\n", x())
	
	// you could do something like this
	// where you call bar() which returns a function that you need to execute with another ()
	fmt.Printf("%d\n", bar()())
	
}

func foo() string {
	return "hello world"
}

// the return type is "func int()"
func bar() func() int {
	return func() int {
		return 451
	}
}
