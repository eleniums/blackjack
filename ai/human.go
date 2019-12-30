package ai

import (
	"fmt"
	"strconv"
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
		case "stay", "s":
			action = game.ActionStay
		case "split", "p":
			action = game.ActionSplit
		case "double", "d":
			action = game.ActionDouble
		case "stats":
			action = game.ActionStats
		case "exit", "quit":
			action = game.ActionExit
		default:
			action = 0
		}
	}

	return action
}

// PlaceBet returns the player's bet.
func (h *Human) PlaceBet(minBet, maxBet, totalMoney int) int {
	var err error
	bet := -1
	for err != nil || bet < minBet || bet > maxBet {
		fmt.Printf("Place bet: ")
		input := game.ReadInput()
		input = strings.ToLower(input)
		bet, err = strconv.Atoi(input)
	}
	return bet
}
