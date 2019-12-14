package players

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// HumanPlayer represents a single human player.
type HumanPlayer struct {
	name   string
	reader *bufio.Reader
}

// NewHumanPlayer will create a new human player.
func NewHumanPlayer(name string) *HumanPlayer {
	reader := bufio.NewReader(os.Stdin)
	return &HumanPlayer{
		name:   name,
		reader: reader,
	}
}

// Name of player.
func (hp *HumanPlayer) Name() string {
	return hp.name
}

// Action returns the action the player wants to make with his hand.
func (hp *HumanPlayer) Action(dealer *game.Hand, player *game.Hand) game.Action {
	var action game.Action
	for action == 0 {
		fmt.Printf("Hit, Stay, Split, or Double: ")
		input := hp.readInput()
		input = strings.ToLower(input)

		switch input {
		case "hit", "h":
			action = game.ActionHit
			break
		case "stay", "s":
			action = game.ActionStay
			break
		case "split", "p":
			action = game.ActionSplit
			break
		case "double", "d":
			action = game.ActionDouble
			break
		default:
			action = 0
			break
		}
	}

	return action
}

// PlaceBet returns the player's bet.
func (hp *HumanPlayer) PlaceBet(minBet, maxBet, totalMoney int) int {
	// TODO: implement player PlaceBet
	return 15
}

func (hp *HumanPlayer) readInput() string {
	input, err := hp.reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\nerror reading input: %v\n", err)
	}
	input = strings.TrimSpace(input)
	return input
}
