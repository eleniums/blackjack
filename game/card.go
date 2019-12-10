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
	return unicodeToString("1F0A1")
}

func unicodeToString(code string) string {
	u := `"\u` + code + `"`
	s, _ := strconv.Unquote(u)
	return s
}
