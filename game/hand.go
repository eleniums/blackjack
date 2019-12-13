package game

import (
	"fmt"
)

// Hand represents a hand of cards.
type Hand struct {
	Cards []Card
}

// NewHand will create a new hand with the given cards.
func NewHand(cards ...Card) *Hand {
	var hand Hand

	for _, v := range cards {
		hand.Cards = append(hand.Cards, v)
	}

	return &hand
}

// Count will return the number of cards in the hand.
func (h *Hand) Count() int {
	return len(h.Cards)
}

// Add a single card to the hand.
func (h *Hand) Add(card Card) {
	h.Cards = append(h.Cards, card)
}

// Total returns the highest hand total without busting, unless busting is unavoidable.
func (h *Hand) Total() int {
	total := 0
	numAces := 0
	for _, v := range h.Cards {
		if v.Rank() > RankAce && v.Rank() < RankJack {
			total += int(v.Rank())
		} else if v.Rank() >= RankJack && v.Rank() <= RankKing {
			total += 10
		} else if v.Rank() == RankAce {
			numAces++
		}
	}

	foundAces := 0
	for _, v := range h.Cards {
		if v.Rank() == RankAce {
			foundAces++
			if foundAces == numAces && total+11 <= 21 {
				total += 11
			} else {
				total++
			}
		}
	}

	return total
}

// String will return a string representation of the hand.
func (h *Hand) String() string {
	var s string
	for _, v := range h.Cards {
		s += fmt.Sprintf("%v  ", v)
	}
	return s
}
