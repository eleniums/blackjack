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
	minBet     float64
	maxBet     float64
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int, maxDiscard int, minBet float64, maxBet float64, dealer *Player, players ...*Player) *Blackjack {
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
	// start clean
	b.emptyHands()

	// place bets for each player
	for _, p := range b.players {
		b.placeBet(p)
		fmt.Println()
	}

	// deal initial hands
	b.dealInitialCards()
	b.displayAll()
	fmt.Println()

	if b.dealer.Hand.IsNatural() {
		// dealer has blackjack, so skip to winners/losers
		b.handleDealerNatural()
		fmt.Println()
		return
	}

	// take actions for each player
	busted := true
	for _, p := range b.players {
		busted = b.playerTurn(p)
		fmt.Println()
	}

	// take actions for dealer
	if busted {
		fmt.Println("Skipping dealer since all players busted or surrendered.")
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
	fmt.Printf("%s has $%.2f. Min bet is $%.2f and max bet is $%.2f.\n", player.Name, player.Money, b.minBet, b.maxBet)
	player.Hand.Bet = player.AI.PlaceBet(b.minBet, b.maxBet, player.Money)
	if player.Hand.Bet >= b.minBet && player.Hand.Bet <= b.maxBet {
		fmt.Printf("%s placed a bet of $%.2f.\n", player.Name, player.Hand.Bet)
	} else {
		fmt.Printf("%s tried to place an invalid bet of $%.2f. Will use minimum bet of $%.2f.\n", player.Name, player.Hand.Bet, b.minBet)
		player.Hand.Bet = b.minBet
	}
}

// playerTurn will take actions for a single player and return true if player busted.
func (b *Blackjack) playerTurn(player *Player) bool {
	fmt.Printf("** %s's turn **\n", player.Name)
	busted := b.playHand(player, player.Hand)

	for i := 0; i < len(player.SplitHands); i++ {
		fmt.Printf("\n** Split hand for %s **\n", player.Name)
		splitBusted := b.playHand(player, player.SplitHands[i])
		if !splitBusted {
			busted = false
		}
	}

	return busted
}

// playHand will play out a single hand from a player.
func (b *Blackjack) playHand(player *Player, hand *game.Hand) bool {
	var action game.Action
	for action != game.ActionStay && action != game.ActionDouble {
		b.displayHand("Dealer", b.dealer.Hand)
		b.displayHand(player.Name, hand)

		if hand.Total() == 21 {
			fmt.Printf("%s has blackjack!\n", player.Name)
			return false
		} else if hand.Total() > 21 {
			fmt.Printf("%s busted with a total of %d.\n", player.Name, hand.Total())
			return true
		}

		action = player.AI.Action(b.dealer.Hand, hand)
		switch action {
		case game.ActionHit:
			card := b.dealCard(hand, false)
			fmt.Printf("%s hit and was dealt: %v\n", player.Name, card)

		case game.ActionStay:
			fmt.Printf("%s chose to stay with a total of %d.\n", player.Name, hand.Total())

		case game.ActionSplit:
			if !hand.CanSplit() {
				fmt.Println("Splitting is only allowed if the starting hand has two cards with equal rank.")
				action = game.ActionInvalid
				continue
			}

			// split hand
			splitHand := game.NewHand(hand.Cards[1])
			splitHand.Bet = hand.Bet
			hand.Cards = hand.Cards[:1]
			player.SplitHands = append(player.SplitHands, splitHand)

			// TODO: deal second card to each split hand

		case game.ActionDouble:
			if !hand.CanDouble() {
				fmt.Println("Doubling down is only allowed on the original two cards.")
				action = game.ActionInvalid
				continue
			}
			card := b.dealCard(hand, false)
			player.Hand.Bet *= 2
			fmt.Printf("%s doubled their bet to $%.2f and was dealt: %v\n", player.Name, player.Hand.Bet, card)
			b.displayHand(player.Name, hand)
			return hand.Total() > 21

		case game.ActionSurrender:
			if !hand.CanDouble() || len(player.SplitHands) > 0 {
				fmt.Println("Surrendering is only allowed on the original two cards before doubling or splitting.")
				action = game.ActionInvalid
				continue
			}
			player.Hand.Bet /= 2
			fmt.Printf("%s surrendered their hand and reduced their bet to $%.2f.\n", player.Name, player.Hand.Bet)
			return true

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
	fmt.Println("** Dealer's turn **")
	b.dealer.Hand.Cards[1].Hidden = false
	fmt.Printf("Dealer revealed their facedown card: %v\n", b.dealer.Hand.Cards[1])
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
		b.determineWinner(p, p.Hand, dealerTotal)
		for _, h := range p.SplitHands {
			b.determineWinner(p, h, dealerTotal)
		}
	}
}

// determineWinner will determine whether a player beat the dealer.
func (b *Blackjack) determineWinner(player *Player, hand *game.Hand, dealerTotal int) {
	playerTotal := hand.Total()
	if hand.IsNatural() {
		fmt.Printf("%s has a natural blackjack!\n", player.Name)
		player.Win++
		player.Money += hand.Bet * 1.5
	} else if playerTotal > 21 {
		fmt.Printf("%s busted with a total of %d.\n", player.Name, playerTotal)
		player.Loss++
		player.Money -= hand.Bet
	} else if dealerTotal > 21 {
		fmt.Printf("%s wins with %d because dealer busted with a total of %d!\n", player.Name, playerTotal, dealerTotal)
		player.Win++
		player.Money += hand.Bet
	} else if playerTotal < dealerTotal {
		fmt.Printf("%s has %d, which loses to dealer's %d.\n", player.Name, playerTotal, dealerTotal)
		player.Loss++
		player.Money -= hand.Bet
	} else if playerTotal == dealerTotal {
		fmt.Printf("Push, %s and dealer both have %d.\n", player.Name, playerTotal)
		player.Tie++
	} else if playerTotal > dealerTotal {
		fmt.Printf("%s has %d, which beats dealer's %d!\n", player.Name, playerTotal, dealerTotal)
		player.Win++
		player.Money += hand.Bet
	}
}

// handleDealerNatural will determine which players won or lost after dealer got a natural blackjack.
func (b *Blackjack) handleDealerNatural() {
	fmt.Println("Dealer has a natural blackjack.")
	for _, p := range b.players {
		if p.Hand.IsNatural() {
			fmt.Printf("Push, %s also has a natural blackjack.\n", p.Name)
			p.Tie++
		} else {
			fmt.Printf("%s loses to dealer's natural blackjack.\n", p.Name)
			p.Loss++
			p.Money -= p.Hand.Bet
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
	total := player.Win + player.Loss + player.Tie
	fmt.Printf("%s (%T)\n", player.Name, player.AI)
	fmt.Printf("  Win: %d (%%%.1f) | Loss: %d (%%%.1f) | Tie: %d (%%%.1f) | $%.2f\n", player.Win, percent(player.Win, total), player.Loss, percent(player.Loss, total), player.Tie, percent(player.Tie, total), player.Money)
}

// percent will calculate a percentage from the given values.
func percent(numerator, denominator int) float32 {
	if denominator == 0 {
		return 0
	}
	return float32(numerator) / float32(denominator) * 100.0
}

// emptyHands will empty all hands in the game and add cards to discard pile.
func (b *Blackjack) emptyHands() {
	for _, c := range b.dealer.Hand.Cards {
		b.discard.Add(0, c)
	}
	b.dealer.Hand.Clear()

	for _, v := range b.players {
		for _, c := range v.Hand.Cards {
			b.discard.Add(0, c)
		}
		v.Hand.Clear()

		for _, s := range v.SplitHands {
			for _, c := range s.Cards {
				b.discard.Add(0, c)
			}
			s.Clear()
		}
		v.SplitHands = v.SplitHands[:0]
	}

	// check if discard pile is too full
	if b.discard.Count() > b.maxDiscard {
		b.shuffler.Add(b.discard.Cards...)
		b.discard.Cards = b.discard.Cards[:0]
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

	// deal first card to dealer face up
	b.dealCard(b.dealer.Hand, false)

	// deal second card to each player face up
	for _, v := range b.players {
		b.dealCard(v.Hand, false)
	}

	// deal second card to dealer face down
	b.dealCard(b.dealer.Hand, true)
}
