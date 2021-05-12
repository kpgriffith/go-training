package main

import (
	"fmt"
)

func main() {
	// create the RECEIVE ONLY channel with 2 capacity
	c := make(<-chan int, 2)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
