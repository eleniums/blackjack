package main

import (
	"fmt"

	"github.com/eleniums/blackjack/game"
)

func main() {
	card := game.NewCard(game.SuiteSpades, 11)
	fmt.Printf("card: %v\n", card)
	fmt.Println("🃑")

	fmt.Println(int(game.RankAce))
	fmt.Println(int(game.RankJack))
	fmt.Println(int(game.RankQueen))
	fmt.Println(int(game.RankKing))
}
