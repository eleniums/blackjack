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

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (h *Human) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	var action game.Action
	for action == 0 {
		h.displayPossibleActions(actions)
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
		case "surrender":
			action = game.ActionSurrender
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
func (h *Human) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	var err error
	bet := -1.0
	for err != nil || bet < minBet || bet > maxBet {
		fmt.Printf("Place bet: ")
		input := game.ReadInput()
		input = strings.ToLower(input)
		if input == "" {
			return minBet
		}
		bet, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid value for bet.")
		} else if bet < minBet {
			fmt.Printf("Bet cannot be less than minimum of $%.2f.\n", minBet)
		} else if bet > maxBet {
			fmt.Printf("Bet cannot be greater than maximum of $%.2f.\n", maxBet)
		}
	}
	return bet
}

// displayPossibleActions will print the available player actions.
func (h *Human) displayPossibleActions(actions []game.Action) {
	display := []string{}

	for _, v := range actions {
		switch v {
		case game.ActionHit:
			display = append(display, "Hit")
		case game.ActionStay:
			display = append(display, "Stay")
		case game.ActionDouble:
			display = append(display, "Double")
		case game.ActionSplit:
			display = append(display, "Split")
		case game.ActionSurrender:
			display = append(display, "Surrender")
		}
	}

	var prompt strings.Builder
	for i, v := range display {
		if i == len(display)-1 {
			prompt.WriteString("or ")
			prompt.WriteString(v)
			prompt.WriteString(": ")
		} else {
			prompt.WriteString(v)
			if len(display) == 2 {
				prompt.WriteString(" ")
			} else {
				prompt.WriteString(", ")
			}
		}
	}

	fmt.Printf(prompt.String())
}
