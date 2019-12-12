package game

import (
	"errors"
)

type Shuffler struct {
	deck Deck
}

func NewShuffler() Shuffler {
	return Shuffler{}
}

func (s Shuffler) Add(cards ...Card) {
	// TODO: shuffle cards into the internal deck
}

func (s Shuffler) Deal() (Card, error) {
	// TODO: implement deal. Might want to switch to a pointer so we can return nil
	return Card{}, errors.New("there are no cards left to deal")
}
