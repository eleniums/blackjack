package main

import (
	"flag"
	"fmt"

	"github.com/eleniums/blackjack/game"
)

func main() {
	flag.BoolVar(&game.UseCardSymbols, "use-card-symbols", false, "set to display card symbols instead of text")
	printCardsTest := flag.Bool("print-cards-test", false, "set to display all cards (for testing purposes)")
	flag.Parse()

	if *printCardsTest {
		displayAllCards()
		return
	}
}

func displayAllCards() {
	fmt.Println("Clubs:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuiteClubs, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Spades:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuiteSpades, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Hearts:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuiteHearts, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Diamonds:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuiteDiamonds, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n")
}
