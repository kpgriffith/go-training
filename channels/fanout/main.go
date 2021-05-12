package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
fanning out
Having a lot of work to do and fanning the work out in many go routines,
then bringing the result back in to 1 channel
*/

func main() {
	todd()
}

func todd() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		// we add a different wait group for each value in the channel
		wg.Add(1)
		// this is fanning out to separate go routines
		go func(v2 int) {
			c2 <- timeConsumingWork(v2) // this brings the values back to a single channel
			wg.Done()
		}(v)
	}
	wg.Wait() // wait for all the go routines to be done
	close(c2) // close the channel, which we are looping over above
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
