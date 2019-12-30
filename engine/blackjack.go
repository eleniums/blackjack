package engine

import (
	"fmt"
	"os"

	"github.com/eleniums/blackjack/game"
)

// Blackjack is the engine for a game of Blackjack.
type Blackjack struct {
	shuffler   game.Shuffler
	dealer     *Player
	players    []*Player
	discard    *game.Deck
	maxDiscard int
	minBet     int
	maxBet     int
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int, maxDiscard int, minBet int, maxBet int, dealer *Player, players ...*Player) *Blackjack {
	shuffler := game.NewShuffler()

	deck := game.NewDeck()
	for i := 0; i < numDecks; i++ {
		shuffler.Add(deck.Cards...)
	}

	discard := game.NewEmptyDeck()

	return &Blackjack{
		shuffler:   shuffler,
		dealer:     dealer,
		players:    players,
		discard:    discard,
		maxDiscard: maxDiscard,
		minBet:     minBet,
		maxBet:     maxBet,
	}
}

// PlayRound will run a single round of blackjack.
func (b *Blackjack) PlayRound() {
	// place bets for each player
	for _, p := range b.players {
		b.placeBet(p)
		fmt.Println()
	}

	// deal initial hands
	b.emptyHands()
	b.dealInitialCards()
	b.displayAll()
	fmt.Println()

	// take actions for each player
	busted := true
	for _, p := range b.players {
		busted = b.playerTurn(p)
		fmt.Println()
	}

	// take actions for dealer
	if busted {
		fmt.Println("Skipping dealer since all players busted.")
	} else {
		b.dealerTurn()
	}
	fmt.Println()

	// determine winners
	b.determineWinners()
	fmt.Println()
}

// DisplayStats will show all player stats.
func (b *Blackjack) DisplayStats() {
	for _, p := range b.players {
		b.displayPlayerStats(p)
	}
}

func (b *Blackjack) placeBet(player *Player) {
	fmt.Printf("%s has $%d. Min bet is $%d and max bet is $%d.\n", player.Name, player.Money, b.minBet, b.maxBet)
	player.Bet = player.AI.PlaceBet(b.minBet, b.maxBet, player.Money)
	if player.Bet >= b.minBet && player.Bet <= b.maxBet {
		fmt.Printf("%s placed a bet of $%d.\n", player.Name, player.Bet)
	} else {
		fmt.Printf("%s tried to place an invalid bet of $%d. Will use minimum bet of $%d.\n", player.Name, player.Bet, b.minBet)
		player.Bet = b.minBet
	}
}

// playerTurn will take actions for a single player and return true if player busted.
func (b *Blackjack) playerTurn(player *Player) bool {
	fmt.Printf("** %s's turn. **\n", player.Name)

	var action game.Action
	for action != game.ActionStay && action != game.ActionDouble {
		b.displayHand("Dealer", b.dealer.Hand)
		b.displayHand(player.Name, player.Hand)

		if player.Hand.Total() == 21 {
			fmt.Printf("%s has blackjack!\n", player.Name)
			return false
		} else if player.Hand.Total() > 21 {
			fmt.Printf("%s busted with a total of %d.\n", player.Name, player.Hand.Total())
			return true
		}

		action = player.AI.Action(b.dealer.Hand, player.Hand)
		switch action {
		case game.ActionHit:
			card := b.dealCard(player.Hand, false)
			fmt.Printf("%s hit and was dealt: %v\n", player.Name, card)
		case game.ActionStay:
			fmt.Printf("%s chose to stay with a total of %d.\n", player.Name, player.Hand.Total())
		case game.ActionSplit:
			// TODO: implement split
		case game.ActionDouble:
			card := b.dealCard(player.Hand, false)
			// TODO: double bet
			fmt.Printf("%s doubled down and was dealt: %v\n", player.Name, card)
			b.displayHand(player.Name, player.Hand)
		case game.ActionStats:
			b.displayPlayerStats(player)
		case game.ActionExit:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}

	return false
}

// dealerTurn will take actions for the dealer.
func (b *Blackjack) dealerTurn() {
	fmt.Println("** Dealer's turn. **")
	b.dealer.Hand.Cards[0].Hidden = false
	fmt.Printf("Dealer revealed their facedown card: %v\n", b.dealer.Hand.Cards[0])
	b.displayHand("Dealer", b.dealer.Hand)

	// dealer hits on soft 17
	for b.dealer.Hand.Total() < 21 && b.dealer.AI.Action(b.dealer.Hand, nil) == game.ActionHit {
		card := b.dealCard(b.dealer.Hand, false)
		fmt.Printf("Dealer hit and was dealt: %v\n", card)
		b.displayHand("Dealer", b.dealer.Hand)
	}

	dealerTotal := b.dealer.Hand.Total()
	if dealerTotal <= 21 {
		fmt.Printf("Dealer finished with a total of %d.\n", dealerTotal)
	} else {
		fmt.Printf("Dealer busted with a total of %d.\n", dealerTotal)
	}
}

// determineWinners will determine which players won or lost.
func (b *Blackjack) determineWinners() {
	dealerTotal := b.dealer.Hand.Total()
	for _, p := range b.players {
		playerTotal := p.Hand.Total()
		if playerTotal > 21 {
			fmt.Printf("%s busted with a total of %d.\n", p.Name, playerTotal)
			p.Loss++
			p.Money -= p.Bet
		} else if dealerTotal > 21 {
			fmt.Printf("%s wins with %d because dealer busted with a total of %d!\n", p.Name, playerTotal, dealerTotal)
			p.Win++
			p.Money += p.Bet
		} else if playerTotal < dealerTotal {
			fmt.Printf("%s has %d, which loses to dealer's %d.\n", p.Name, playerTotal, dealerTotal)
			p.Loss++
			p.Money -= p.Bet
		} else if playerTotal == dealerTotal {
			fmt.Printf("Push, %s and dealer both have %d.\n", p.Name, playerTotal)
			p.Tie++
		} else if playerTotal > dealerTotal {
			fmt.Printf("%s has %d, which beats dealer's %d!\n", p.Name, playerTotal, dealerTotal)
			p.Win++
			p.Money += p.Bet
		}
	}
}

// displayAll will display all cards on the table.
func (b *Blackjack) displayAll() {
	b.displayHand("Dealer", b.dealer.Hand)
	for _, v := range b.players {
		b.displayHand(v.Name, v.Hand)
	}
}

// displayHand will display the given hand.
func (b *Blackjack) displayHand(name string, hand *game.Hand) {
	fmt.Printf("%s: %v= %d\n", name, hand, hand.Total())
}

// displayPlayerStats will display the stats for a single player.
func (b *Blackjack) displayPlayerStats(player *Player) {
	fmt.Printf("%s (%T)\n", player.Name, player.AI)
	fmt.Printf("  Win: %d | Loss: %d | Tie: %d | $%d\n", player.Win, player.Loss, player.Tie, player.Money)
}

func (b *Blackjack) emptyHands() {
	b.dealer.Hand.Cards = b.dealer.Hand.Cards[:0]
	for _, v := range b.players {
		v.Hand.Cards = v.Hand.Cards[:0]
	}
}

// dealCard will deal a single card to the given hand.
func (b *Blackjack) dealCard(hand *game.Hand, faceDown bool) game.Card {
	card, err := b.shuffler.Deal()
	if err != nil {
		panic(err)
	}
	card.Hidden = faceDown
	hand.Add(card)
	return card
}

// dealInitialCards will deal cards for a new round.
func (b *Blackjack) dealInitialCards() {
	// deal first card to each player face up
	for _, v := range b.players {
		b.dealCard(v.Hand, false)
	}

	// deal first card to dealer face down
	b.dealCard(b.dealer.Hand, true)

	// deal second card to each player face up
	for _, v := range b.players {
		b.dealCard(v.Hand, false)
	}

	// deal second card to dealer face up
	b.dealCard(b.dealer.Hand, false)
}
