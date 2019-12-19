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
	players  []Player
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int, players ...Player) *Blackjack {
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
	for i, p := range b.players {
		fmt.Printf("** %s's turn. **\n", p.Name())

		var action game.Action
		for action != game.ActionStay && action != game.ActionDouble {
			b.displayPlayerHand(p.Name(), b.hands[i])

			if b.hands[i].Total() == 21 {
				fmt.Printf("%s has blackjack!\n", p.Name())
				break
			} else if b.hands[i].Total() > 21 {
				fmt.Printf("%s busted.\n", p.Name())
				break
			}

			action = p.Action(b.dealer, b.hands[i])
			switch action {
			case game.ActionHit:
				card := b.dealCard(b.hands[i], false)
				fmt.Printf("%s hit and was dealt: %v\n", p.Name(), card)
				break
			case game.ActionStay:
				fmt.Printf("%s chose to stay.\n", p.Name())
				break
			case game.ActionSplit:
				// TODO: implement split
				break
			case game.ActionDouble:
				card := b.dealCard(b.hands[i], false)
				// TODO: double bet
				fmt.Printf("%s doubled down and was dealt: %v\n", p.Name(), card)
				break
			default:
				break
			}
		}

		fmt.Println()
	}

	// take actions for dealer
	// TODO: first reveal facedown dealer card
	// TODO: figure out rules for when dealer should hit or stay (soft 17?)
	fmt.Println("** Dealer's turn. **")
	for b.dealer.Total() <= 17 {
		card := b.dealCard(b.dealer, false)
		fmt.Printf("Dealer hit and was dealt: %v\n", card)
	}

	fmt.Println()

	// determine winners
	// TODO: determine who won and lost and collect bets
}

// displayAll will display all cards on the table.
func (b *Blackjack) displayAll() {
	fmt.Printf("Dealer: %v= %d\n", b.dealer, b.dealer.Total())
	for i, v := range b.hands {
		fmt.Printf("%s: %v= %d\n", b.players[i].Name(), v, v.Total())
	}
}

// displayPlayerHand will display the dealer hand and the player's given hand.
func (b *Blackjack) displayPlayerHand(name string, hand *game.Hand) {
	fmt.Printf("Dealer: %v= %d\n", b.dealer, b.dealer.Total())
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
