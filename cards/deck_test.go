package main

import (
	"errors"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLen := 52
	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %d %d", expectedLen, len(d))
	}

	if d[0] != "ace of hearts" {
		t.Errorf("First element should be ace of hearts but got %v", d[0])
	}

	if d[len(d)-1] != "king of spades" {
		t.Errorf("Last element should be ace of hearts but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	testFilePath := "_decktesting"
	os.Remove(testFilePath)

	newDeck().saveToFile(testFilePath)
	if _, err := os.Stat(testFilePath); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected that a file would exist at path %v", testFilePath)
	}

	loadedDeck := newDeckFromFile(testFilePath)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, only got %v", len(loadedDeck))
	}

	os.Remove(testFilePath)
}

func TestDeal(t *testing.T) {
	d := newDeck()
	dealSize := 10
	hand, rest := deal(d, dealSize)

	if len(d) != len(hand)+len(rest) {
		t.Errorf("Expected that the len(d) %v should equal len(hand) %v + len(rest) %v", len(d), len(hand), len(rest))
	}

	if dealSize != len(hand) {
		t.Errorf("Expected that dealing %v cards results in a hand of 10 cards", dealSize)
	}
}
