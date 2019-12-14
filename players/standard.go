package players

import (
	"github.com/eleniums/blackjack/game"
)

// StandardAI is an opponent that uses a standard strategy.
type StandardAI struct {
}

// NewStandardAI will create a new standard AI.
func NewStandardAI() *StandardAI {
	return &StandardAI{}
}

// Name of player.
func (ai *StandardAI) Name() string {
	return "Larry"
}

// Action returns the action the player wants to make with his hand.
func (ai *StandardAI) Action(dealer game.Hand, player game.Hand) game.Action {
	// TODO: implement player Action
	return 0
}

// PlaceBet returns the player's bet.
func (ai *StandardAI) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
