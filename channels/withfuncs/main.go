package main

import "fmt"

func main() {

	// create the channel
	c := make(chan int) // general bi-directional channel

	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) { // went from general channel to specific one
	c <- 42
}

// recieve
func bar(c <-chan int) { // went from general channel to specific one
	fmt.Println(<-c)
}
