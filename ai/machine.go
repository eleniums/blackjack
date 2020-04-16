package ai

import (
	"github.com/eleniums/blackjack/game"
	"github.com/eleniums/blackjack/machine"
)

// Machine is an opponent that uses machine learning to make decisions.
type Machine struct {
	model *machine.Model
}

// NewMachine will create a new machine learning AI.
func NewMachine(modelFile string) *Machine {
	return &Machine{
		model: machine.NewModel(modelFile),
	}
}

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (ai *Machine) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	p := ai.model.Predict(dealer, player)

	if allowed(actions, p.Action) {
		// if surrender is recommended, take the surrender
		if p.Action == game.ActionSurrender {
			return p.Action
		}

		// if result is positive, proceed with action
		if p.Result == game.ResultWin || p.Result == game.ResultTie || p.Result == game.ResultNone {
			return p.Action
		}
	}

	return game.ActionStay
}

// PlaceBet returns the player's bet.
func (ai *Machine) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	return minBet
}
