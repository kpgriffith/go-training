package main

import (
	"fmt"
)

func main() {

	// this will show that x's scope is different between a and b
	
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// the scope of x is the function
// but also available in the function that is returned from this function
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

