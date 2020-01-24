package ai

import (
	"github.com/eleniums/blackjack/game"
)

// Soft17Dealer is a dealer that will hit on a soft 17.
type Soft17Dealer struct{}

// NewSoft17Dealer will create a new dealer AI.
func NewSoft17Dealer() *Soft17Dealer {
	return &Soft17Dealer{}
}

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (ai *Soft17Dealer) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	// dealer hits on soft 17
	for dealer.Total() < 17 || (dealer.Total() == 17 && dealer.Soft()) {
		return game.ActionHit
	}
	return game.ActionStay
}

// PlaceBet returns the player's bet.
func (ai *Soft17Dealer) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	// dealer does not place any bets
	return 0
}
