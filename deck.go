package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"clubs", "hearts", "diamonds", "spades"}
	values := []string{"ace", "two", "three", "queen", "king"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) print() {

	for i, cardName := range d {
		fmt.Println(i, cardName)
	}
	fmt.Println("finished iterating deck of card")
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func saveToFile(d deck, fileName string) {
	ioutil.WriteFile(fileName, []byte(d.toString()), 0644)
}

func createNewDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	strArray := strings.Split(string(bs), ",")
	return deck(strArray)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
