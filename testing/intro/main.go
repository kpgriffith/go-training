package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 5 + 6 = ", mySum(4, 5, 6))
}

func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
