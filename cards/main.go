package main

import (
	"fmt"
)

func main() {
	deck := newDeck()
	deck.shuffle()

	fmt.Println("deck", deck)
	// fmt.Println("random number", rand.Intn(100))
}
