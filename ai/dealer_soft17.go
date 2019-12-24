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

// Action returns the action the player wants to make with his hand.
func (ai *Soft17Dealer) Action(dealer *game.Hand, player *game.Hand) game.Action {
	// TODO: implement soft 17 dealer Action
	return 0
}

// PlaceBet returns the player's bet.
func (ai *Soft17Dealer) PlaceBet(minBet, maxBet, totalMoney int) int {
	// dealer does not place any bets
	return 0
}
