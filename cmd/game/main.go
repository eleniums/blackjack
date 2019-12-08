package main

import (
	"fmt"

	"github.com/eleniums/blackjack/game"
)

func main() {
	card := game.NewCard(game.SuiteSpades, 11)
	fmt.Printf("card: %v\n", card)
	fmt.Println("🃑")
}
