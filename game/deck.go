package game

import (
	"errors"
	"fmt"
)

// ErrDeckEmpty means the deck is empty.
var ErrDeckEmpty = errors.New("deck is empty")

// Deck represents a deck of cards.
type Deck struct {
	Cards []Card
}

// NewDeck creates a new deck of 52 cards. The cards are not in random order.
func NewDeck() *Deck {
	var deck Deck

	// add clubs
	for i := 1; i <= 13; i++ {
		card := NewCard(SuitClubs, Rank(i))
		deck.Cards = append(deck.Cards, card)
	}

	// add spades
	for i := 1; i <= 13; i++ {
		card := NewCard(SuitSpades, Rank(i))
		deck.Cards = append(deck.Cards, card)
	}

	// add hearts
	for i := 1; i <= 13; i++ {
		card := NewCard(SuitHearts, Rank(i))
		deck.Cards = append(deck.Cards, card)
	}

	// add diamonds
	for i := 1; i <= 13; i++ {
		card := NewCard(SuitDiamonds, Rank(i))
		deck.Cards = append(deck.Cards, card)
	}

	return &deck
}

// NewEmptyDeck will create an empty deck.
func NewEmptyDeck() *Deck {
	return &Deck{}
}

// Count returns the number of cards in the deck.
func (d *Deck) Count() int {
	return len(d.Cards)
}

// Add a card to the deck at the specific index.
func (d *Deck) Add(i int, card Card) {
	d.Cards = append(d.Cards, Card{})
	copy(d.Cards[i+1:], d.Cards[i:])
	d.Cards[i] = card
}

// Deal the top card from the deck. This removes the card from the deck.
func (d *Deck) Deal() (Card, error) {
	if d.Count() == 0 {
		return Card{}, ErrDeckEmpty
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}

// String will return a string representation of the deck.
func (d *Deck) String() string {
	var s string
	for _, v := range d.Cards {
		s += fmt.Sprintf("%v  ", v)
	}
	return s
}
