package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	suits := []string{"hearts", "clubs", "diamonds", "spades"}
	cards := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	d := deck{}
	for _, suit := range suits {
		for _, card := range cards {
			d = append(d, card+" of "+suit)
		}
	}

	return d
}

func newDeckFromFile(path string) deck {
	bs, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// func (d deck) deal(handSize int) (deck, deck) {
// indicates that maybe deck is being mutated? deck.deal(5)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) _shuffle(r *rand.Rand) {
	for i := range d {
		swap := r.Intn(i + 1)
		d[i], d[swap] = d[swap], d[i]
	}
}

func (d deck) shuffle() {
	d._shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(path string) error {
	return os.WriteFile(path, []byte(d.toString()), 0666)
}
