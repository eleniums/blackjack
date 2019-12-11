package game

import (
	"fmt"
	"strconv"
)

var UseCardSymbol = false

type Card struct {
	Suite Suite
	Rank  Rank
}

func NewCard(suite Suite, rank Rank) Card {
	return Card{
		Suite: suite,
		Rank:  rank,
	}
}

func (c Card) String() string {
	if UseCardSymbol {
		return c.Symbol()
	}
	return c.Text()
}

func (c Card) Text() string {
	return fmt.Sprintf("%v%v", c.Rank, c.Suite)
}

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

func unicodeToString(code string) string {
	i, _ := strconv.ParseInt(code, 16, 32)
	r := rune(i)
	return string(r)
}
