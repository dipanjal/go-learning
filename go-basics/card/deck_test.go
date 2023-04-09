package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var deck = newDeck()
var numOfCards = 52

func TestNewDeck(t *testing.T) {
	err := doCommonDeckAssertion(deck)
	if err != nil {
		t.Error(err)
	}
}

// added this test case just to increase test coverage
func TestPrint(t *testing.T) {
	deck.print()
}

func TestToString(t *testing.T) {
	expected := "Ace of Spades,2 of Spades,3 of Spades,4 of Spades,5 of Spades,6 of Spades,7 of Spades,8 of Spades,9 of Spades,10 of Spades,Jack of Spades,Queen of Spades,King of Spades,Ace of Hearts,2 of Hearts,3 of Hearts,4 of Hearts,5 of Hearts,6 of Hearts,7 of Hearts,8 of Hearts,9 of Hearts,10 of Hearts,Jack of Hearts,Queen of Hearts,King of Hearts,Ace of Diamonds,2 of Diamonds,3 of Diamonds,4 of Diamonds,5 of Diamonds,6 of Diamonds,7 of Diamonds,8 of Diamonds,9 of Diamonds,10 of Diamonds,Jack of Diamonds,Queen of Diamonds,King of Diamonds,Ace of Clubs,2 of Clubs,3 of Clubs,4 of Clubs,5 of Clubs,6 of Clubs,7 of Clubs,8 of Clubs,9 of Clubs,10 of Clubs,Jack of Clubs,Queen of Clubs,King of Clubs"
	actual := deck.toString()
	if actual != expected {
		t.Error("Not matched")
	}
}

func TestDeal(t *testing.T) {
	handSize := 12
	cardsInHand, cardsRemaining := deal(deck, handSize)

	if len(cardsInHand) != handSize {
		t.Errorf("Cards In Hand Size | expected %v, actual %v", handSize, len(cardsInHand))
	}

	expCardsRemSize := numOfCards - handSize
	if len(cardsRemaining) != expCardsRemSize {
		t.Errorf("Cards Remaining Size | expected %v, actual %v", expCardsRemSize, len(cardsRemaining))
	}
}

func TestSaveAndCreateNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	cleanUpTestFiles()

	err := deck.saveToFile(fileName)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}
	readError, deckFromFile := createNewDeckFromFile(fileName)
	if readError != nil {
		t.Errorf("Unexpected Error: %s", err)
	}
	deckError := doCommonDeckAssertion(deckFromFile)
	if deckError != nil {
		t.Error(deckError)
	}
	cleanUpTestFiles()
}

func TestSaveAndReadFromFileFailed(t *testing.T) {
	fileName := "_decktesting"
	cleanUpTestFiles()

	err := deck.saveToFile(fileName)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}
	readError, _ := createNewDeckFromFile("random_file")
	if readError == nil {
		t.Error("Expected read error, got: nil")
	}
	cleanUpTestFiles()
}

func TestShuffle(t *testing.T) {
	aDeck := newDeck()
	aDeck.shuffle()
	if reflect.DeepEqual(deck, aDeck) {
		t.Error("Deck is not shuffled")
	}
}

// private functions
func doCommonDeckAssertion(deck Deck) error {
	if len(deck) != numOfCards {
		return fmt.Errorf("length test | expected=%v, actual=%v\n", numOfCards, len(deck))
	}

	firstValueExpected := "Ace of Spades"
	firstValueActual := deck[0]
	if firstValueActual != firstValueExpected {
		return fmt.Errorf("first value | expected=%v, actual=%v\n", firstValueExpected, firstValueActual)
	}

	lastValueExpected := "King of Clubs"
	lastValueActual := deck[len(deck)-1]
	if lastValueActual != lastValueExpected {
		return fmt.Errorf("last value | expected=%v, actual=%v\n", lastValueExpected, firstValueActual)
	}

	if reflect.DeepEqual(deck, newDeck()) == false {
		return fmt.Errorf("invalid deck\n")
	}

	return nil
}

func cleanUpTestFiles() {
	os.RemoveAll(DirFixed)
}
