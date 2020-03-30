package engine

import (
	"os"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// ML contains methods for generating machine learning training data.
type ML struct {
	data   *os.File
	record *record
}

// NewML will create a new ML instance with an open file for storing training data.
func NewML(trainingDataFile string) *ML {
	data, err := os.Create(trainingDataFile)
	if err != nil {
		panic(err)
	}
	data.WriteString("Dealer,Player,Result\n")
	return &ML{
		data:   data,
		record: nil,
	}
}

// Close the open file handle. This should always be called.
func (m *ML) Close() {
	m.data.Close()
}

// StartRecord will begin a new record.
func (m *ML) StartRecord(dealer, player *game.Hand, action game.Action) {
	if m == nil {
		return
	}

	// close out any existing record
	if m.record != nil {
		m.WriteRecord(game.ResultNone)
	}

	// a copy is used so the cards will not stay revealed
	d := game.NewHand(dealer.Cards...)

	// we need to reveal all the cards so this works for the dealer
	d.Cards[0].Hidden = false
	d.Cards[1].Hidden = false

	m.record = &record{
		dealer: m.formatHand(d),
		player: m.formatHand(player),
		action: action,
	}
}

// WriteRecord will write a completed record to a file.
func (m *ML) WriteRecord(result game.Result) {
	if m == nil || m.record == nil {
		return
	}

	m.record.result = result
	m.record.Write(m.data)
	m.record = nil
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
