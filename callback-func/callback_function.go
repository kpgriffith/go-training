package main

import (
	"fmt"
)

func main() {	
	ii := []int{1,2,3,4,5,6,7,8,9,}
	fmt.Println("all numbers", sum(ii...))
	
	fmt.Println("even numbers", even(sum, ii...))
	
	fmt.Println("odd numbers", odd(sum, ii...))

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, y ...int) int {
	// create a slice of just even numbers
	var z []int
	for _, v := range y {
		if v % 2 == 0 {
			z = append(z, v)
		}
	}
	
	// execute the callback function that was passed in and return the value returned from it
	return f(z...)
}

func odd(f func(x ...int) int, y ...int) int {
	// create a slice of just even numbers
	var z []int
	for _, v := range y {
		if v % 2 != 0 {
			z = append(z, v)
		}
	}
	
	// execute the callback function that was passed in and return the value returned from it
	return f(z...)
}