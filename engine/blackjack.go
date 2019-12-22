package engine

import (
	"fmt"

	"github.com/eleniums/blackjack/game"
)

// Blackjack is the engine for a game of Blackjack.
type Blackjack struct {
	shuffler game.Shuffler
	dealer   *game.Hand
	hands    []*game.Hand
	players  []AI
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int, players ...AI) *Blackjack {
	shuffler := game.NewShuffler()

	deck := game.NewDeck()
	for i := 0; i < numDecks; i++ {
		shuffler.Add(deck.Cards...)
	}

	hands := []*game.Hand{}
	for range players {
		hands = append(hands, game.NewHand())
	}

	return &Blackjack{
		shuffler: shuffler,
		dealer:   game.NewHand(),
		hands:    hands,
		players:  players,
	}
}

// PlayRound will run a single round of blackjack.
func (b *Blackjack) PlayRound() {
	// deal initial hands
	b.emptyHands()
	b.dealInitialCards()
	b.displayAll()
	fmt.Println()

	// take actions for each player
	busted := true
	for i, p := range b.players {
		busted = b.playerTurn(p, b.hands[i])
		fmt.Println()
	}

	// take actions for dealer
	if busted {
		fmt.Println("All players busted.")
	} else {
		b.dealerTurn()
	}
	fmt.Println()

	// determine winners
	// TODO: determine who won and lost and collect bets
}

// playerTurn will take actions for a single player and return true if player busted.
func (b *Blackjack) playerTurn(player AI, hand *game.Hand) bool {
	fmt.Printf("** %s's turn. **\n", player.Name())

	var action game.Action
	for action != game.ActionStay && action != game.ActionDouble {
		b.displayHand("Dealer", b.dealer)
		b.displayHand(player.Name(), hand)

		if hand.Total() == 21 {
			fmt.Printf("%s has blackjack!\n", player.Name())
			return false
		} else if hand.Total() > 21 {
			fmt.Printf("%s busted.\n", player.Name())
			return true
		}

		action = player.Action(b.dealer, hand)
		switch action {
		case game.ActionHit:
			card := b.dealCard(hand, false)
			fmt.Printf("%s hit and was dealt: %v\n", player.Name(), card)
			break
		case game.ActionStay:
			fmt.Printf("%s chose to stay.\n", player.Name())
			break
		case game.ActionSplit:
			// TODO: implement split
			break
		case game.ActionDouble:
			card := b.dealCard(hand, false)
			// TODO: double bet
			fmt.Printf("%s doubled down and was dealt: %v\n", player.Name(), card)
			b.displayHand(player.Name(), hand)
			break
		default:
			break
		}
	}

	return false
}

// dealerTurn will take actions for the dealer.
func (b *Blackjack) dealerTurn() {
	fmt.Println("** Dealer's turn. **")
	b.dealer.Cards[0].Hidden = false
	fmt.Printf("Dealer revealed their facedown card: %v\n", b.dealer.Cards[0])
	b.displayHand("Dealer", b.dealer)

	// dealer hits on soft 17
	for b.dealer.Total() < 17 || (b.dealer.Total() == 17 && b.dealer.Soft()) {
		card := b.dealCard(b.dealer, false)
		fmt.Printf("Dealer hit and was dealt: %v\n", card)
		b.displayHand("Dealer", b.dealer)
	}
}

// displayAll will display all cards on the table.
func (b *Blackjack) displayAll() {
	b.displayHand("Dealer", b.dealer)
	for i, v := range b.hands {
		b.displayHand(b.players[i].Name(), v)
	}
}

// displayHand will display the given hand.
func (b *Blackjack) displayHand(name string, hand *game.Hand) {
	fmt.Printf("%s: %v= %d\n", name, hand, hand.Total())
}

func (b *Blackjack) emptyHands() {
	b.dealer.Cards = b.dealer.Cards[:0]
	for _, v := range b.hands {
		v.Cards = v.Cards[:0]
	}
}

func (b *Blackjack) dealCard(hand *game.Hand, faceDown bool) game.Card {
	card, err := b.shuffler.Deal()
	if err != nil {
		panic(err)
	}
	card.Hidden = faceDown
	hand.Add(card)
	return card
}

func (b *Blackjack) dealInitialCards() {
	// deal first card to each player face up
	for _, v := range b.hands {
		b.dealCard(v, false)
	}

	// deal first card to dealer face down
	b.dealCard(b.dealer, true)

	// deal second card to each player face up
	for _, v := range b.hands {
		b.dealCard(v, false)
	}

	// deal second card to dealer face up
	b.dealCard(b.dealer, false)
}
