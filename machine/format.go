package machine

import (
	"math"
	"strconv"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// ConvertHand will take a string representation of a hand and convert it into an integer.
func ConvertHand(hand string) int {
	total := 0.0
	split := strings.Split(hand, " ")

	for i, v := range split {
		var n int
		switch v {
		case "J", "Q", "K":
			n = 10
		case "A":
			n = 11
		default:
			n, _ = strconv.Atoi(v)
		}
		total += float64(n) * math.Pow(100.0, float64(i))
	}

	return int(total)
}

// ConvertResult will take a string representation of a result and convert it into an integer.
func ConvertResult(result string) int {
	return int(ParseLabel(result))
}

// FormatHand will return a formatted hand string. The returned string is used as the input to ConvertHand.
func FormatHand(hand *game.Hand) string {
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
