package main

import (
	"fmt"
	"log"
)

var cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}

var cardValues = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

//var cardValues = []string{"Ace", "2", "3", "4"}

func main() {
	newDeck := newDeck()
	fileName := "my-deck.txt"
	err := newDeck.saveToFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	deckFromFile := createNewDeckFromFile(fileName)
	fmt.Println(deckFromFile.toString())

	cardsInHand, cardsRemaining := deal(deckFromFile, 12)
	fmt.Println("Cards in hand....")
	cardsInHand.print()
	fmt.Println("Cards Remaining....")
	cardsRemaining.print()

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
