package players

import (
	"github.com/eleniums/blackjack/game"
)

// HumanPlayer represents a single human player.
type HumanPlayer struct {
	name string
}

// NewHumanPlayer will create a new human player.
func NewHumanPlayer(name string) *HumanPlayer {
	return &HumanPlayer{
		name: name,
	}
}

// Name of player.
func (hp *HumanPlayer) Name() string {
	return hp.name
}

// Action returns the action the player wants to make with his hand.
func (hp *HumanPlayer) Action(dealer *game.Hand, player *game.Hand) game.Action {
	// TODO: implement player Action
	return 0
}

// PlaceBet returns the player's bet.
func (hp *HumanPlayer) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
