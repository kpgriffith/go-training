package main

import (
	"fmt"
)

/*
new type that is an interface

any other types that implement the getGreeting() string function
belong to type bot
*/
type bot interface {
	getGreeting() string
}

type englishBot struct {

}

type spanishBot struct {

}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

//func printGreetingEnglish(eb englishBot){
//	fmt.Println(eb.getGreeting())
//}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}


//func printGreetingSpanish(sb spanishBot){
//	fmt.Println(sb.getGreeting())
//}


// printGreeting is the same for both structs
// I want to use 1 function for either struct
// enter interfaces...

// new function that uses the interface as the input parameter.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
