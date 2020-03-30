package engine

import (
	"github.com/eleniums/blackjack/game"
	"github.com/eleniums/blackjack/machine"
)

// Player is someone playing against the dealer.
type Player struct {
	Name       string
	AI         AI
	Hand       *game.Hand
	SplitHands []*game.Hand
	Money      float64
	Win        int
	Loss       int
	Tie        int
	Records    map[string]*machine.Record
}

// NewPlayer will create a new player instance.
func NewPlayer(name string, money float64, ai AI) *Player {
	return &Player{
		Name:    name,
		AI:      ai,
		Hand:    game.NewHand(),
		Money:   money,
		Win:     0,
		Loss:    0,
		Tie:     0,
		Records: map[string]*machine.Record{},
	}
}
