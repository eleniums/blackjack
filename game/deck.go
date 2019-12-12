package game

import (
	"errors"
	"fmt"
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

func (d *Deck) Add(index int, card Card) {
	deck := *d

	before := deck[0:index]
	after := deck[index:]

	deck = append(before, card)
	deck = append(deck, after...)

	*d = deck
}

func (d *Deck) Deal() (Card, error) {
	if d.Empty() {
		return Card{}, errors.New("deck is empty")
	}

	deck := *d
	card := deck[0]
	*d = deck[1:]

	return card, nil
}

func (d Deck) Empty() bool {
	return len(d) == 0
}

func (d Deck) String() {
	for _, v := range d {
		fmt.Printf("%v  ", v)
	}
}
