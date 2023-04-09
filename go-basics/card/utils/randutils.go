package utils

import (
	"math/rand"
	"time"
)

// NewRand will return new Randomizer
// here we are using UnixNano to create the Random Seeder
// to generate unique random pattern everytime
func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
