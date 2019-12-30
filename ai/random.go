package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eleniums/blackjack/game"
)

// Random is an opponent that picks random actions.
type Random struct {
	rand *rand.Rand
}

// NewRandom will create a new random AI.
func NewRandom() *Random {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return &Random{
		rand: r,
	}
}

// Action returns the action the player wants to make with his hand.
func (ai *Random) Action(dealer *game.Hand, player *game.Hand) game.Action {
	i := ai.rand.Intn(2)
	fmt.Println(i)
	if i == 0 {
		return game.ActionHit
	}
	return game.ActionStay
}

// PlaceBet returns the player's bet.
func (ai *Random) PlaceBet(minBet, maxBet, totalMoney int) int {
	return minBet
}
