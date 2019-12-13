package game

import (
	"math/rand"
	"time"
)

// Shuffler will shuffle a deck of cards and deal from the top.
type Shuffler struct {
	deck *Deck
	rand *rand.Rand
}

// NewShuffler will create a new shuffler.
func NewShuffler() Shuffler {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return Shuffler{
		deck: &Deck{},
		rand: r,
	}
}

// Add will shuffle the given cards into the shuffler.
func (s Shuffler) Add(cards ...Card) {
	for _, v := range cards {
		i := s.rand.Intn(s.deck.Count() + 1)
		s.deck.Add(i, v)
	}
}

// Deal will deal the top card from the shuffler.
func (s Shuffler) Deal() (Card, error) {
	return s.deck.Deal()
}
