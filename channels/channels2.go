package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://walmart.com", "http://google.com", "http://golang.org", "http://stackoverflow.com",}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}

	// infinite loop
	//for { // infinite loop
	//	go makeRequest(<-c, c)  // receiving a value from a channel is a blocking operation
	//}

	// use of a range loop with a channel
	// instead of the infinite loop above
	// it assigns what is coming out of the channel to l and iterates as long as something
	// is coming out of the channel
	for l := range c {
		// using a Function Literal to make sure the sleep is in the
		// correct go routine.  But I don't want it as part of the makeRequest() function itself
		//go func() {
		//	time.Sleep(2 * time.Second)
		//	// PROBLEM with this...  Both the parent and child routine are pointing to the same location for l
		//	// need to pass in a value so it is separate for the child
		//	makeRequest(l, c)
		//}()

		// new version passing in the link
		go func(lnk string) {
			time.Sleep(2 * time.Second)
			// PROBLEM with this...  Both the parent and child routine are pointing to the same location for l
			// need to pass in a value so it is separate for the child
			makeRequest(lnk, c)
		}(l)
	}
}

func makeRequest(l string, c chan string){
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s is down!!\n", l)
		c <- l
		return
	}
	fmt.Printf("%s is up!!\n", l)
	c <- l
}

