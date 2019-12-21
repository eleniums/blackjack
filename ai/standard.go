package ai

import (
	"github.com/eleniums/blackjack/game"
)

// Standard is an opponent that uses a standard strategy.
type Standard struct {
}

// NewStandard will create a new standard AI.
func NewStandard() *Standard {
	return &Standard{}
}

// Name of player.
func (ai *Standard) Name() string {
	return "Larry"
}

// Action returns the action the player wants to make with his hand.
func (ai *Standard) Action(dealer *game.Hand, player *game.Hand) game.Action {
	// TODO: implement player Action
	return 0
}

// PlaceBet returns the player's bet.
func (ai *Standard) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
