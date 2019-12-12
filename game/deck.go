package game

import (
	"errors"
)

type Deck []Card

func NewDeck() Deck {
	var deck Deck

	// add clubs
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteClubs, Rank(i))
		deck = append(deck, card)
	}

	// add spades
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteSpades, Rank(i))
		deck = append(deck, card)
	}

	// add hearts
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteHearts, Rank(i))
		deck = append(deck, card)
	}

	// add diamonds
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteDiamonds, Rank(i))
		deck = append(deck, card)
	}

	return deck
}

func (d *Deck) Deal() (Card, error) {
	if d.Empty() {
		return Card{}, errors.New("deck is empty")
	}

	deck := *d
	card := deck[0]
	*d = deck[1:len(deck)]

	return card, nil
}

func (d Deck) Empty() bool {
	return len(d) == 0
}
