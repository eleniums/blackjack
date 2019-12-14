package engine

import (
	"github.com/eleniums/blackjack/game"
)

// Player interface.
type Player interface {
	// Action returns the action the player wants to make with his hand.
	Action(dealer game.Hand, player game.Hand) game.Action

	// PlaceBet returns the player's bet.
	PlaceBet(minBet, maxBet, totalMoney int) int
}
