package main

import (
	"flag"
	"fmt"

	"github.com/eleniums/blackjack/engine"
	"github.com/eleniums/blackjack/game"
)

func main() {
	flag.BoolVar(&game.UseCardSymbols, "use-card-symbols", false, "set to display card symbols instead of text")
	printCardsTest := flag.Bool("print-cards-test", false, "set to display all cards (for testing purposes)")
	numDecks := flag.Int("num-decks", 6, "number of shuffled decks to use")
	// maxDiscard := flag.Int("max-discard", 20, "number of cards allowed in discard pile before shuffling them back in")
	// startingMoney := flag.Int("starting-money", 100, "amount of money players start with")
	// minBet := flag.Int("min-bet", 15, "minimum bet allowed")
	// maxBet := flag.Int("max-bet", 15, "maximum bet allowed")
	flag.Parse()

	if *printCardsTest {
		displayAllCards()
		return
	}

	if *numDecks < 0 {
		fmt.Println("Number of decks has to be 1 or greater.")
		return
	}

	blackjack := engine.NewBlackjack(*numDecks)

	for blackjack.Update() {
		// do nothing
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
