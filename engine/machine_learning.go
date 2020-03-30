package engine

import (
	"fmt"
	"os"

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

// Write hands and result to a file for machine learning training purposes.
func (m *ML) Write(dealer, player *game.Hand, action game.Action, result game.Result) {
	if m == nil {
		return
	}
	m.data.WriteString(fmt.Sprintf("D: %v - P: %v : %s_%s\n", dealer, player, action.String(), result.String()))
}
