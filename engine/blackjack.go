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
	b.display()

	// take actions for each player
	for i, p := range b.players {
		if b.hands[i].Total() == 21 {
			fmt.Printf("%s has blackjack!", p.Name())
			continue
		} else if b.hands[i].Total() > 21 {
			fmt.Printf("%s busted.", p.Name())
			continue
		}

		action := p.Action(b.dealer, b.hands[i])
		switch action {
		case game.Hit:
			break
		case game.Stay:
			break
		case game.Split:
			break
		case game.Double:
			break
		default:
			break
		}
	}
}

// display all cards on the table.
func (b *Blackjack) display() {
	fmt.Printf("Dealer: %v\n\n", b.dealer)
	for i, v := range b.hands {
		fmt.Printf("%s: %v\n\n", b.players[i].Name(), v)
	}
}

func (b *Blackjack) emptyHands() {
	b.dealer.Cards = b.dealer.Cards[:0]
	for _, v := range b.hands {
		v.Cards = v.Cards[:0]
	}
}

func (b *Blackjack) dealInitialCards() {
	var card game.Card
	var err error

	// deal first card to each player face up
	for _, v := range b.hands {
		card, err = b.shuffler.Deal()
		if err != nil {
			panic(err)
		}
		v.Add(card)
	}

	// deal first card to dealer face down
	card, err = b.shuffler.Deal()
	if err != nil {
		panic(err)
	}
	card.Hidden = true
	b.dealer.Add(card)

	// deal second card to each player face up
	for _, v := range b.hands {
		card, err = b.shuffler.Deal()
		if err != nil {
			panic(err)
		}
		v.Add(card)
	}

	// deal second card to dealer face up
	card, err = b.shuffler.Deal()
	if err != nil {
		panic(err)
	}
	b.dealer.Add(card)
}
