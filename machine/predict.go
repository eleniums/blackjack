package machine

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// Predict will feed a dealer hand and player hand into a model and return the resulting label.
func Predict(dealer *game.Hand, player *game.Hand, modelFile string, predictScript string) Label {
	d := strconv.Itoa(ConvertHand(FormatHand(dealer)))
	p := strconv.Itoa(ConvertHand(FormatHand(player)))

	cmd := exec.Command(predictScript, modelFile, d, p)
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
