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

	// take actions for each player
	for i, p := range b.players {
		var action game.Action
		for action != game.ActionStay && action != game.ActionDouble {
			b.display()

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
				b.dealCard(b.hands[i], false)
				break
			case game.ActionStay:
				break
			case game.ActionSplit:
				break
			case game.ActionDouble:
				break
			default:
				break
			}
		}
		fmt.Println()
	}
}

// display all cards on the table.
func (b *Blackjack) display() {
	fmt.Printf("Dealer: %v\n", b.dealer)
	for i, v := range b.hands {
		fmt.Printf("%s: %v\n", b.players[i].Name(), v)
	}
}

func (b *Blackjack) emptyHands() {
	b.dealer.Cards = b.dealer.Cards[:0]
	for _, v := range b.hands {
		v.Cards = v.Cards[:0]
	}
}

func (b *Blackjack) dealCard(hand *game.Hand, faceDown bool) {
	card, err := b.shuffler.Deal()
	if err != nil {
		panic(err)
	}
	card.Hidden = faceDown
	hand.Add(card)
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
