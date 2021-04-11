package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // type of the address is a pointer (*int)

	/*
		a is type int
		&a gives the address to where the value is stores
		the TYPE of the address (&a) is *int (pointer)
	*/

	// I can assign the address to a variable
	b := &a
	// so now b is an address not the value
	fmt.Println(b)
	// to get the value of b
	fmt.Println(*b) // this is an operator that de-references the address and gives me the value

	// we could also use this syntax to de-reference an address
	fmt.Println(*&a)
	
	// let's change the value at the address using b
	// and see what happens when we print a
	*b = 43 // set the value that is at the address of b and set it to 43
		
	fmt.Println(a)

}
