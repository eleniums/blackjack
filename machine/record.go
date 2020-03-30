package machine

import (
	"fmt"
	"os"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// IncludePlayerName will add the player's name to the training data if true.
var IncludePlayerName = false

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
		Dealer: formatHand(dealer),
		Player: formatHand(player),
	}
}

// Write record to the given file.
func (r *Record) Write(file *os.File) {
	if IncludePlayerName {
		file.WriteString(fmt.Sprintf("%v_%v,%s,%s,%s\n", r.Action, r.Result, r.Dealer, r.Player, r.Name))
	} else {
		file.WriteString(fmt.Sprintf("%v_%v,%s,%s\n", r.Action, r.Result, r.Dealer, r.Player))
	}
}

// formatHand will return a formatted hand string.
func formatHand(hand *game.Hand) string {
	// a copy is used so the cards will not stay revealed
	local := game.NewHand(hand.Cards...)

	// we need to reveal all the cards
	for i := range local.Cards {
		local.Cards[i].Hidden = false
	}

	// remove suits as they are not needed
	cleaned := strings.TrimSpace(local.String())
	cleaned = strings.ReplaceAll(cleaned, "♣", "")
	cleaned = strings.ReplaceAll(cleaned, "♠", "")
	cleaned = strings.ReplaceAll(cleaned, "♥", "")
	cleaned = strings.ReplaceAll(cleaned, "♦", "")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")

	return cleaned
}
