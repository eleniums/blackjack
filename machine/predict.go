package machine

import (
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/eleniums/blackjack/game"
)

const model = "./machine/model.bin"

// Predict will feed a dealer hand and player hand into a model and return the resulting label.
func Predict(dealer *game.Hand, player *game.Hand) Label {
	d := strconv.Itoa(convertHand(formatHand(dealer)))
	p := strconv.Itoa(convertHand(formatHand(player)))

	cmd := exec.Command("python3", "./machine/predict.py", model, d, p)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		panic(err)
	}

	return Label(n)
}

// convertHand will take a string representation of a hand and convert it into an integer.
func convertHand(hand string) int {
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
