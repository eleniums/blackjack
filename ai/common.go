package ai

import (
	"github.com/eleniums/blackjack/game"
)

// within will return true if the value is within the inclusive range [low, high].
func within(value, low, high int) bool {
	return value >= low && value <= high
}

// allowed will return true if the action is allowed based on the available actions.
func allowed(actions []game.Action, action game.Action) bool {
	for _, v := range actions {
		if v == action {
			return true
		}
	}
	return false
}
