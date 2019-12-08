package game

import (
	"fmt"
)

type Card struct {
	Suite Suite
	Rank  int
}

func NewCard(suite Suite, rank int) Card {
	return Card{
		Suite: suite,
		Rank:  rank,
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.Rank, c.Suite)
}
