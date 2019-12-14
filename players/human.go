package players

import (
	"github.com/eleniums/blackjack/game"
)

// HumanPlayer represents a single human player.
type HumanPlayer struct {
}

// Action returns the action the player wants to make with his hand.
func (hp *HumanPlayer) Action(dealer game.Hand, player game.Hand) game.Action {
	// TODO: implement player Action
	return 0
}

// PlaceBet returns the player's bet.
func (hp *HumanPlayer) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
