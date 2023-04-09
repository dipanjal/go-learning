package main

import (
	"fmt"
	"log"
)

func main() {
	myDeck := spawnADeck()
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
func spawnADeck() Deck {
	fileName := "my-cards.txt"
	err, deck := createNewDeckFromFile(fileName)
	if deck != nil {
		return deck
	}
	log.Printf("Unexpected error occurred while reading file: %s\n", err)
	return createNewDeckAndSaveIntoFile(fileName)
}

func createNewDeckAndSaveIntoFile(fileName string) Deck {
	newDeck := newDeck()
	err := newDeck.saveToFile(fileName)
	if err != nil {
		log.Printf("Unexpected error occurred while saving deck into file %s\n", err)
	}
	return newDeck
}
