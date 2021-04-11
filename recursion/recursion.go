package main

import (
	"fmt"
)

func main() {
	fmt.Println("Count up using Fibonacci's Sequence...")
	fmt.Printf("%d\n%d\n", 0, 1)
	getNext(0, 1, 10)
}

func getNext(f int, s int, m int) {
	n := f + s
	if n < m {
		fmt.Printf("%d\n", n)
		getNext(s, n, m)
	}
}
