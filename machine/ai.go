package machine

import (
	"github.com/eleniums/blackjack/game"
)

// AI is an opponent that uses a previously trained model to make decisions.
type AI struct{}

// NewAI will create a new machine learning AI.
func NewAI() *AI {
	return &AI{}
}

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (ai *AI) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	// TODO
	return game.ActionStay
}

// PlaceBet returns the player's bet.
func (ai *AI) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	return minBet
}
