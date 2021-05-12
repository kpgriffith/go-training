package main

import "fmt"

func main() {

	// create the channel
	c := make(chan int) // general bi-directional channel

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive using a range
	// it will loop until the channel is closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
