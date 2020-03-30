package machine

import (
	"fmt"
	"os"
	"strings"

	"github.com/eleniums/blackjack/game"
)

type record struct {
	dealer string
	player string
	action game.Action
	result game.Result
}

func (r *record) Write(file *os.File) {
	file.WriteString(fmt.Sprintf("%v_%v,%s,%s\n", r.action, r.result, r.dealer, r.player))
}

func (r *record) AddDealerHand(hand *game.Hand) {
	// a copy is used so the cards will not stay revealed
	local := game.NewHand(hand.Cards...)

	// we need to reveal all the dealer cards
	local.Cards[0].Hidden = false
	local.Cards[1].Hidden = false

	r.dealer = r.formatHand(local)
}

func (r *record) AddPlayerHand(hand *game.Hand) {
	r.player = r.formatHand(hand)
}

// formatHand will return a formatted hand string.
func (r *record) formatHand(hand *game.Hand) string {
	cleaned := strings.TrimSpace(hand.String())
	cleaned = strings.ReplaceAll(cleaned, "♣", "")
	cleaned = strings.ReplaceAll(cleaned, "♠", "")
	cleaned = strings.ReplaceAll(cleaned, "♥", "")
	cleaned = strings.ReplaceAll(cleaned, "♦", "")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	return cleaned
}
