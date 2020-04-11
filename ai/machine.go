package ai

import (
	"github.com/eleniums/blackjack/game"
	"github.com/eleniums/blackjack/machine"
)

// Machine is an opponent that uses machine learning to make decisions.
type Machine struct {
	model  string
	script string
}

// NewMachine will create a new machine learning AI.
func NewMachine(modelFile string, predictScript string) *Machine {
	return &Machine{
		model:  modelFile,
		script: predictScript,
	}
}

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (ai *Machine) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	prediction := machine.Predict(dealer, player, ai.model, ai.script)
	action, result := prediction.Split()

	if allowed(actions, action) {
		// if surrender is recommended, take the surrender
		if action == game.ActionSurrender {
			return action
		}

		// if result is positive, proceed with action
		if result == game.ResultWin || result == game.ResultTie || result == game.ResultNone {
			return action
		}
	}

	return game.ActionStay
}

// PlaceBet returns the player's bet.
func (ai *Machine) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	return minBet
}
