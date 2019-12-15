package players

import (
	"fmt"
	"strings"

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
	var action game.Action
	for action == 0 {
		fmt.Printf("Hit, Stay, Split, or Double: ")
		input := game.ReadInput()
		input = strings.ToLower(input)

		switch input {
		case "hit", "h":
			action = game.ActionHit
			break
		case "stay", "s":
			action = game.ActionStay
			break
		case "split", "p":
			action = game.ActionSplit
			break
		case "double", "d":
			action = game.ActionDouble
			break
		default:
			action = 0
			break
		}
	}

	return action
}

// PlaceBet returns the player's bet.
func (hp *HumanPlayer) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
