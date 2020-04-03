package machine

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/eleniums/blackjack/game"
)

const model = "./machine/model.bin"

// Predict will feed a dealer hand and player hand into a model and return the resulting label.
func Predict(dealer *game.Hand, player *game.Hand) Label {
	d := strconv.Itoa(ConvertHand(formatHand(dealer)))
	p := strconv.Itoa(ConvertHand(formatHand(player)))

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
