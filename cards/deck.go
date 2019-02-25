package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

// create a new type 'deck' that is a slice of strings
type deck []string

func newDeck() deck {  // no receiver because this will create a new deck
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNames := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range cardSuites {
		for _, n := range cardNames {
			cards = append(cards, n + " of " + s)
		}
	}

	return cards
}

// This is a function with a 'deck' receiver
// Any variable define as type 'deck' can now call 'print()'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
The print function is similar to the following in Java

public class Deck {

	private ArrayList<String> cards = new ArrayList<String>();

	public void print(cards){

		for (String card : cards){
			System.out.println(card);
		}

	}

}
 */

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error
		fmt.Println("Error:", err)
		// exit with error
		os.Exit(1)
		// or we could send a default deck
		//return newDeck()
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}