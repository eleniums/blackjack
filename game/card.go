package game

import (
	"fmt"
	"strconv"
)

// UseCardSymbols determines whether to use the unicode card symbols for the text representation or straight text.
var UseCardSymbols = false

// Card represents a single card.
type Card struct {
	suit  Suit
	rank   Rank
	Hidden bool
}

// NewCard will return a new card with the given suit and rank.
func NewCard(suit Suit, rank Rank) Card {
	return Card{
		suit: suit,
		rank:  rank,
	}
}

// Suit will return the card's suit or a default value if the card is hidden.
func (c Card) Suit() Suit {
	if c.Hidden {
		return 0
	}
	return c.suit
}

// Rank will return the card's rank or a default value if the card is hidden.
func (c Card) Rank() Rank {
	if c.Hidden {
		return 0
	}
	return c.rank
}

// String will return the string representation of the card.
func (c Card) String() string {
	if UseCardSymbols {
		return c.Symbol()
	}
	return c.Text()
}

// Text will return the text representation of the card.
func (c Card) Text() string {
	return fmt.Sprintf("%v%v", c.Rank(), c.Suit())
}

// Symbol will return the unicode card symbol representation of the card.
func (c Card) Symbol() string {
	code := "1F0"

	switch c.Suit() {
	case SuitClubs:
		code += "D"
	case SuitSpades:
		code += "A"
	case SuitHearts:
		code += "B"
	case SuitDiamonds:
		code += "C"
	default:
		code += "A"
	}

	// skip the "knight" playing card symbol
	if c.Rank() < 12 {
		code = fmt.Sprintf("%s%X", code, int(c.Rank()))
	} else {
		code = fmt.Sprintf("%s%X", code, int(c.Rank()+1))
	}

	return unicodeToString(code)
}

// unicodeToString will convert a hex string to the unicode symbol it is associated with.
func unicodeToString(code string) string {
	i, _ := strconv.ParseInt(code, 16, 32)
	r := rune(i)
	return string(r)
}
