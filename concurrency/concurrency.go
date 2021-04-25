package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	wg.Add(1) // increment the wait group

	go foo()
	bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())

	wg.Wait() // wait for all the wait groups to complete
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // decrement the wait group to indicate this is done
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
