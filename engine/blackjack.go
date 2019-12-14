package engine

import (
	"github.com/eleniums/blackjack/game"
)

// Blackjack is the engine for a game of Blackjack.
type Blackjack struct {
	shuffler game.Shuffler
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int) *Blackjack {
	shuffler := game.NewShuffler()

	for i := 0; i < numDecks; i++ {
		deck := game.NewDeck()
		shuffler.Add(deck.Cards...)
	}

	return &Blackjack{
		shuffler: shuffler,
	}
}

// PlayRound will run a single round of blackjack.
func (b *Blackjack) PlayRound() {
	// TODO: this is the game engine main loop, return true to keep looping
}
