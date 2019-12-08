package game

import (
	"fmt"
)

type Card struct {
	Suite Suite
	Value int
}

func NewCard(suite Suite, value int) Card {
	return Card{
		Suite: suite,
		Value: value,
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%v%d", c.Suite, c.Value)
}
