package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	lengthExpected := 52
	if len(deck) != lengthExpected {
		t.Errorf("length test | expected=%v, actual=%v\n", lengthExpected, len(deck))
	}

	firstValueExpected := "Ace of Spades"
	firstValueActual := deck[0]
	if firstValueActual != firstValueExpected {
		t.Errorf("first value | expected=%v, actual=%v\n", firstValueExpected, firstValueActual)
	}

	lastValueExpected := "King of Clubs"
	lastValueActual := deck[len(deck)-1]
	if lastValueActual != lastValueExpected {
		t.Errorf("last value | expected=%v, actual=%v\n", lastValueExpected, firstValueActual)
	}
}
