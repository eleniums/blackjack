package game

import (
	"errors"
)

// Shuffler will shuffle a deck of cards and deal from the top.
type Shuffler struct {
	deck *Deck
}

// NewShuffler will create a new shuffler.
func NewShuffler() Shuffler {
	return Shuffler{
		deck: &Deck{},
	}
}

// Add will shuffle the given cards into the shuffler.
func (s Shuffler) Add(cards ...Card) {
	// TODO: shuffle cards into the internal deck
}

// Deal will deal the top card from the shuffler.
func (s Shuffler) Deal() (Card, error) {
	// TODO: implement deal. Might want to switch to a pointer so we can return nil
	return Card{}, errors.New("there are no cards left to deal")
}
