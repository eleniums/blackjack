package engine

import (
	"fmt"
	"os"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// ML contains methods for generating machine learning training data.
type ML struct {
	data *os.File
}

// NewML will create a new ML instance with an open file for storing training data.
func NewML(trainingDataFile string) *ML {
	data, err := os.Create(trainingDataFile)
	if err != nil {
		panic(err)
	}
	return &ML{
		data: data,
	}
}

// Close the open file handle. This should always be called.
func (m *ML) Close() {
	m.data.Close()
}

// WriteAction to a file for machine learning training purposes.
func (m *ML) WriteAction(dealer, player *game.Hand, action game.Action) {
	if m == nil {
		return
	}

	// a copy is used so the cards will not stay revealed
	d := game.NewHand(dealer.Cards...)

	// we need to reveal all the cards so this works for the dealer
	d.Cards[0].Hidden = false
	d.Cards[1].Hidden = false

	m.data.WriteString(fmt.Sprintf("D:%s-P:%s-%v", m.formatHand(d), m.formatHand(player), action))
}

// WriteResult to a file for machine learning training purposes.
func (m *ML) WriteResult(result game.Result) {
	if m == nil {
		return
	}
	m.data.WriteString(fmt.Sprintf("_%v\n", result))
}

// formatHand for training data.
func (m *ML) formatHand(hand *game.Hand) string {
	cleaned := strings.TrimSpace(hand.String())
	cleaned = strings.ReplaceAll(cleaned, "♣", "")
	cleaned = strings.ReplaceAll(cleaned, "♠", "")
	cleaned = strings.ReplaceAll(cleaned, "♥", "")
	cleaned = strings.ReplaceAll(cleaned, "♦", "")
	cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	return cleaned
}
