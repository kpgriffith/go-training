package main

import (
	"fmt"
)

func main() {
	// create the channel with 2 capacity
	c := make(chan int, 2)
	// put something on the channel
	c <- 42
	// take something off
	fmt.Println(<-c)

	// put something on the channel
	c <- 43
	c <- 44
	// take something off
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	// print the type
	fmt.Printf("%T\n", c)
}
