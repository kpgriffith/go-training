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

	wg.Add(2) // increment the wait group
	go foo()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())
	go func() {
		fmt.Println("func literal")
		wg.Done()
	}()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("Print foo")
	wg.Done()
}
