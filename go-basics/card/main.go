package main

import (
	"fmt"
	"log"
)

var cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}

var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

func main() {
	myDeck := createADeck()
	myDeck.print()

	fmt.Println("..... Cards after shuffle .....")
	myDeck.shuffle()
	myDeck.print()

	fmt.Println("..... start dealing the cards .....")
	cardsInHand, cardsRemaining := deal(myDeck, 12)
	fmt.Printf("Cards in hand (%d) .....\n", len(cardsInHand))
	fmt.Println(cardsInHand.toString())
	fmt.Printf("Cards remaining (%d) .....\n", len(cardsRemaining))
	fmt.Println(cardsRemaining.toString())
}

// try to create a deck from existing file
// is the file is not present, it will create a new deck and save into a file
func createADeck() Deck {
	fileName := "my-cards.txt"
	err, deck := createNewDeckFromFile(fileName)
	if deck != nil {
		return deck
	}
	log.Printf("Unexpected error occurred while reading file: %s\n", err)
	return createAndSaveDeck(fileName)
}

func createAndSaveDeck(fileName string) Deck {
	newDeck := newDeck()
	err := newDeck.saveToFile(fileName)
	if err != nil {
		log.Printf("Unexpected error occurred while saving deck into file %s\n", err)
	}
	return newDeck
}

func newDeck() Deck {
	cards := Deck{}
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}
