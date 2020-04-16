package machine

import (
	"github.com/eleniums/blackjack/game"
)

// Prediction contains a single outcome from a model.
type Prediction struct {
	Action game.Action
	Result game.Result
	Score  float64
}
