package main

import "fmt"

func main() {

	// couple ways to declare
	// first way
	colors := map[string]string{
		"red":   "#ff0000",
		"kelly": "#4CBB17",
	}

	// to add more key value pairs
	colors["white"] = "#ffffff"

	fmt.Println(colors)

	// to remove a key value pair
	delete(colors, "red")

	fmt.Println(colors)

	// second way
	var colorsAgain map[string]string
	fmt.Println(colorsAgain)

	// another way
	colorsMore := make(map[string]string)
	fmt.Println(colorsMore)

	// func that iterates over a map
	printMap(colors)
}

// iterate over the map to gain access to each pair
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is:", value)
	}
}
