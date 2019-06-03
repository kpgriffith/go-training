package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){

	fn := os.Args[1]

	fmt.Println("Using the os package")
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
	}

	b := make([]byte, 50)
	f.Read(b)
	fmt.Printf("%s\n\n", string(b))

	fmt.Println("Using the io/ioutil package")
	bs, error := ioutil.ReadFile(fn)
	if error != nil {
		fmt.Println("Error:", error)
	}
	fmt.Printf("%s\n", string(bs))

}
