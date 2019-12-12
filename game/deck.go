package game

import (
	"errors"
	"fmt"
)

// Deck represents a deck of cards.
type Deck struct {
	cards []Card
}

// NewDeck creates a new deck of 52 cards. The cards are not in random order.
func NewDeck() *Deck {
	var deck Deck

	// add clubs
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteClubs, Rank(i))
		deck.cards = append(deck.cards, card)
	}

	// add spades
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteSpades, Rank(i))
		deck.cards = append(deck.cards, card)
	}

	// add hearts
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteHearts, Rank(i))
		deck.cards = append(deck.cards, card)
	}

	// add diamonds
	for i := 1; i <= 13; i++ {
		card := NewCard(SuiteDiamonds, Rank(i))
		deck.cards = append(deck.cards, card)
	}

	return &deck
}

// Count returns the number of cards in the deck.
func (d *Deck) Count() int {
	return len(d.cards)
}

// Cards will return the cards in this deck.
func (d *Deck) Cards() []Card {
	return d.cards
}

// Add a card to the deck at the specific index.
func (d *Deck) Add(i int, card Card) {
	d.cards = append(d.cards, Card{})
	copy(d.cards[i+1:], d.cards[i:])
	d.cards[i] = card
}

// Deal the top card from the deck. This removes the card from the deck.
func (d *Deck) Deal() (Card, error) {
	if d.Empty() {
		return Card{}, errors.New("deck is empty")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]

	return card, nil
}

// Empty returns true if the deck is empty.
func (d *Deck) Empty() bool {
	return len(d.cards) == 0
}

// String will return a string representation of the deck.
func (d *Deck) String() string {
	var s string
	for _, v := range d.cards {
		s += fmt.Sprintf("%v  ", v)
	}
	return s
}
