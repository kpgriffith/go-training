package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		// use a select to pull values off the channels
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from the comma ok channel:", v, ok)
				return
			} else {
				fmt.Println("from the comma ok channel:", v, ok)
			}
		}
	}
}
