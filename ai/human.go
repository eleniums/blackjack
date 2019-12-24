package ai

import (
	"fmt"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// Human represents a single human player.
type Human struct{}

// NewHuman will create a new human player.
func NewHuman() *Human {
	return &Human{}
}

// Action returns the action the player wants to make with his hand.
func (h *Human) Action(dealer *game.Hand, player *game.Hand) game.Action {
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
func (h *Human) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}
