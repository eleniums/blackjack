package machine

import (
	"fmt"
	"os"

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
