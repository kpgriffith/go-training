package main

import "fmt"

func main(){

	colors := map[string]string{
		"red":"#ff0000",
		"kelly":"#4CBB17",
		"white":"#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string){
	for key, value := range c{
		fmt.Println("Hex code for", key, "is:", value)
	}
}