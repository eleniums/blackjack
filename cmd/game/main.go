package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/eleniums/blackjack/ai"
	"github.com/eleniums/blackjack/engine"
	"github.com/eleniums/blackjack/game"
)

var version = "0.1"

func main() {
	flag.BoolVar(&game.UseCardSymbols, "use-card-symbols", false, "set to display card symbols instead of text")
	printCardsTest := flag.Bool("print-cards-test", false, "set to display all cards (for testing purposes)")
	numDecks := flag.Int("num-decks", 6, "number of shuffled decks to use")
	numRounds := flag.Int("num-rounds", 3, "number of rounds to play (0 is infinite)")
	numPlayers := flag.Int("num-players", 1, "number of human players in game")
	maxDiscard := flag.Int("max-discard", 100, "number of cards allowed in discard pile before shuffling them back in")
	startingMoney := flag.Float64("starting-money", 100, "amount of money players start with")
	minBet := flag.Float64("min-bet", 15, "minimum bet allowed")
	maxBet := flag.Float64("max-bet", 500, "maximum bet allowed")
	delay := flag.Int("delay", 0, "add a millisecond delay between rounds to slow the game down")
	addRandomAI := flag.Bool("random-ai", false, "add an ai that randomly chooses actions")
	addStandardAI := flag.Bool("standard-ai", false, "add an ai that uses a standard strategy")
	generateTrainingData := flag.Bool("generate-training-data", false, "generate and save machine learning training data")
	trainingDataFile := flag.String("training-data-file", "./training.csv", "output file for machine learning training data")
	flag.Parse()

	if *printCardsTest {
		displayAllCards()
		return
	}

	if *numDecks <= 0 {
		fmt.Println("Number of decks has to be 1 or greater.")
		return
	}

	fmt.Printf("--- Blackjack v%s ---\n\n", version)

	var players []*engine.Player

	// add human players
	for i := 0; i < *numPlayers; i++ {
		fmt.Printf("Enter player name: ")
		name := game.ReadInput()
		if strings.TrimSpace(name) == "" {
			name = fmt.Sprintf("Player %d", i+1)
		}
		player := engine.NewPlayer(name, *startingMoney, ai.NewHuman())
		players = append(players, player)
	}
	fmt.Println()

	// add computer players
	if *addRandomAI {
		randomAI := engine.NewPlayer("Larry", *startingMoney, ai.NewRandom())
		players = append(players, randomAI)
	}

	if *addStandardAI {
		standardAI := engine.NewPlayer("Joe", *startingMoney, ai.NewStandard())
		players = append(players, standardAI)
	}

	if len(players) <= 0 {
		fmt.Println("Number of players has to be 1 or greater.")
		return
	}

	// create dealer
	dealer := engine.NewPlayer("Dealer", 0, ai.NewSoft17Dealer())

	blackjack := engine.NewBlackjack(*numDecks, *maxDiscard, *minBet, *maxBet, dealer, players...)

	// check if machine learning training data should be generated
	if *generateTrainingData {
		ml := engine.NewML(*trainingDataFile)
		defer ml.Close()
		blackjack.ML = ml
	}

	// show starting stats
	if *numRounds == 0 {
		fmt.Println("Ready to begin game with infinite rounds.")
	} else {
		fmt.Printf("Ready to begin game with %d rounds.\n", *numRounds)
	}
	fmt.Printf("All players start with $%.2f.\n", *startingMoney)
	fmt.Printf("Min bet is $%.2f. Max bet is $%.2f.\n", *minBet, *maxBet)
	fmt.Println()
	blackjack.DisplayStats()
	fmt.Println()

	for i := 1; i <= *numRounds || *numRounds == 0; i++ {
		if *delay > 0 {
			time.Sleep(time.Duration(*delay) * time.Millisecond)
		}

		fmt.Printf("--- Round %d ---\n", i)
		blackjack.PlayRound()
		blackjack.DisplayStats()
		fmt.Println()
	}
}

func displayAllCards() {
	fmt.Println("Clubs:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuitClubs, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Spades:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuitSpades, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Hearts:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuitHearts, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n\n")

	fmt.Println("Diamonds:")
	for i := 1; i <= 13; i++ {
		card := game.NewCard(game.SuitDiamonds, game.Rank(i))
		fmt.Printf("%v  ", card)
	}
	fmt.Printf("\n")
}
