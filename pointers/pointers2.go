package main

import (
	"fmt"
)

// everything is pass by value

func main() {
	x := 0
	foo(x)
	fmt.Println("x after passing without pointer", x)
	
	bar(&x)
	fmt.Println("address of x after bar", &x)
	fmt.Println("value of x after bar", x)
}

// no pointer
func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

// passed the pointer to a variable.
func bar(z *int) {
	fmt.Println(z)   // this will give the address because it's a pointer
	fmt.Println(*z)  // derefense the pointer to get the value
	*z = 43          // dereference to get the value and change the value
	fmt.Println(z)
	fmt.Println(*z)
}
