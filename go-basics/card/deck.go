package main

import (
	"fmt"
	"go-basics/card/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Deck is a decorated type of []string slice
type Deck []string

const DirFixed = "./decks"
const Separator = ","

// a receiver function which can be used by any type of Deck
// the argument deck is like this in java or self in python
// in go it's a convention to use the actual name instead of this or self
func (deck Deck) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
	fmt.Println("No of Cards in the Deck: ", len(deck))
}

func deal(deck Deck, handSize int) (Deck, Deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck Deck) toString() string {
	return strings.Join(deck, Separator)
}

func (deck Deck) saveToFile(fileName string) error {
	filePath := filepath.Join(DirFixed, fileName)

	dirError := utils.CreateDirectory(DirFixed)
	if dirError != nil {
		return dirError
	}

	fileError := utils.SaveToFile(filePath, []byte(deck.toString()))
	if fileError == nil {
		log.Println("successfully saved into", filePath)
	}
	return fileError
}

func createNewDeckFromFile(fileName string) Deck {
	filePath := filepath.Join(DirFixed, fileName)
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicf("Unexpected error occurred while reading file: %s\n", err)
	}
	return strings.Split(string(bytes), Separator)
}
