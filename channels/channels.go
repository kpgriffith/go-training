package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://walmart.com", "http://google.com", "http://golang.org", "http://stackoverflow.com",}

	c := make(chan string)

	for _, link := range links {
		go makeRequest(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // this is blocking
	}
}

func makeRequest(l string, c chan string){
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s is down!!\n", l)
		c <- "Might be down I think"
		return
	}
	fmt.Printf("%s is up!!\n", l)
	c <- "Yep it's up: " + l
}
