package engine

import (
	"github.com/eleniums/blackjack/game"
)

// Player is someone playing against the dealer.
type Player struct {
	Name  string
	AI    AI
	Hand  *game.Hand
	Money float64
	Bet   float64
	Win   int
	Loss  int
	Tie   int
}

// NewPlayer will create a new player instance.
func NewPlayer(name string, money float64, ai AI) *Player {
	return &Player{
		Name:  name,
		AI:    ai,
		Hand:  game.NewHand(),
		Money: money,
		Bet:   0,
		Win:   0,
		Loss:  0,
		Tie:   0,
	}
}
