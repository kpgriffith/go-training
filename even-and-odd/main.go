package main

import "fmt"

func main(){
	mySlice := []int{
		0,1,2,3,4,5,6,7,8,9,10,
	}

	for myInt := range mySlice {
		if myInt % 2 == 0 {
			fmt.Println(myInt, "is even")
		} else {
			fmt.Println(myInt, "is odd")
		}
	}
}
