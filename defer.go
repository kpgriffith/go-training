package main

import (
	"fmt"
)

// good for when you open a file or a connection and want to make sure you close it
// defer can be used to close the connection, right after opening.
// this helps with good code organization

func main() {
	defer foo() // defers it to when main() exits
	bar()
}

func foo(){
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}