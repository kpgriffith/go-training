package main

import (
	"fmt"
)

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	ret := sum(xi...) // because we already have a slice, we need to unravel it or unroll it, basically it undoes the slice to many ints
	fmt.Println(ret)
	
	// we could also pass in nothing
	sum1 := sum()
	fmt.Println(sum1)
	
	newSum := helloSum("kevin", xi...)
	fmt.Println(newSum)
}

func sum(x ...int) int { // turns x into a slice
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	
	return sum
}

// if there are multiple parameters and a variadic parameter is one of them, it must be the final parameter, because you can pass 0 or more
func helloSum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	fmt.Println(s)
	
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	
	return sum
}