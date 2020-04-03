package machine

import (
	"fmt"
	"os"

	"github.com/eleniums/blackjack/game"
)

// Debug will add the player's name to the training data and print the training data to stdout.
var Debug = false

// Record of a single player action.
type Record struct {
	Name   string
	Dealer string
	Player string
	Action game.Action
	Result game.Result
}

// NewRecord will create a new record.
func NewRecord(dealer, player *game.Hand, name string) *Record {
	return &Record{
		Name:   name,
		Dealer: FormatHand(dealer),
		Player: FormatHand(player),
	}
}

// Write record to the given file.
func (r *Record) Write(file *os.File) {
	line := fmt.Sprintf("%v_%v,%s,%s", r.Action, r.Result, r.Dealer, r.Player)

	if Debug {
		line = fmt.Sprintf("%s,%s", line, r.Name)
		fmt.Println(line)
	}

	file.WriteString(line + "\n")
}
