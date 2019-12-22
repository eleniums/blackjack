package engine

import (
	"github.com/eleniums/blackjack/game"
)

// Player is someone playing against the dealer.
type Player struct {
	Name  string
	Money int
	AI    AI
	Hand  []*game.Hand
	Bet   []int
}

// NewPlayer will create a new player instance.
func NewPlayer(name string, money int, ai AI) *Player {
	return &Player{
		Name:  name,
		Money: money,
		AI:    ai,
		Hand:  []*game.Hand{},
		Bet:   []int{},
	}
}
