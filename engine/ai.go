package engine

import (
	"github.com/eleniums/blackjack/game"
)

// AI interface.
type AI interface {
	// Action returns the action the player wants to make with his hand from the given array of possible actions.
	Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action

	// PlaceBet returns the player's bet.
	PlaceBet(minBet, maxBet, totalMoney float64) float64
}
