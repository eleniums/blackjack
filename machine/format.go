package machine

import (
	"math"
	"strconv"
	"strings"
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
