package game

import (
	"fmt"
	"strconv"
)

// UseCardSymbol determines whether to use the unicode card symbols for the text representation or straight text.
var UseCardSymbol = false

// Card represents a single card.
type Card struct {
	Suite Suite
	Rank  Rank
}

// NewCard will return a new card with the given suite and rank.
func NewCard(suite Suite, rank Rank) Card {
	return Card{
		Suite: suite,
		Rank:  rank,
	}
}

// String will return the string representation of this card.
func (c Card) String() string {
	if UseCardSymbol {
		return c.Symbol()
	}
	return c.Text()
}

// Text will return the text representation of this card.
func (c Card) Text() string {
	return fmt.Sprintf("%v%v", c.Rank, c.Suite)
}

// Symbol will return the unicode card symbol representation of this card.
func (c Card) Symbol() string {
	code := "1F0"

	switch c.Suite {
	case SuiteClubs:
		code += "D"
		break
	case SuiteSpades:
		code += "A"
		break
	case SuiteHearts:
		code += "B"
		break
	case SuiteDiamonds:
		code += "C"
		break
	}

	// skip the "knight" playing card symbol
	if c.Rank < 12 {
		code = fmt.Sprintf("%s%X", code, int(c.Rank))
	} else {
		code = fmt.Sprintf("%s%X", code, int(c.Rank+1))
	}

	return unicodeToString(code)
}

// unicodeToString will convert a hex string to the unicode symbol it is associated with.
func unicodeToString(code string) string {
	i, _ := strconv.ParseInt(code, 16, 32)
	r := rune(i)
	return string(r)
}
