package game

import (
	"fmt"
)

// Hand represents a hand of cards.
type Hand struct {
	ID          string
	Cards       []Card
	Bet         float64
	Surrendered bool
}

// NewHand will create a new hand with the given cards.
func NewHand(cards ...Card) *Hand {
	var hand Hand

	hand.ID = NewUUID()

	for _, v := range cards {
		hand.Cards = append(hand.Cards, v)
	}

	return &hand
}

// Count will return the number of cards in the hand.
func (h *Hand) Count() int {
	return len(h.Cards)
}

// Clear will reset a hand to empty.
func (h *Hand) Clear() {
	h.Cards = h.Cards[:0]
	h.Bet = 0
	h.Surrendered = false
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

// Soft returns true if the hand is a soft hand, meaning it includes an ace that is counted as 11.
func (h *Hand) Soft() bool {
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
				return true
			}
			total++
		}
	}

	return false
}

// Hard returns true if the hand is a hard hand, meaning it includes no aces or only aces that are counted as 1.
func (h *Hand) Hard() bool {
	return !h.Soft()
}

// IsInitialHand will return true if hand only has two cards.
func (h *Hand) IsInitialHand() bool {
	return h.Count() == 2
}

// CanDouble will return true if doubling down is allowed.
func (h *Hand) CanDouble() bool {
	return h.IsInitialHand()
}

// CanSplit will return true if splitting is allowed.
func (h *Hand) CanSplit() bool {
	return h.IsInitialHand() && h.Cards[0].Rank() == h.Cards[1].Rank()
}

// IsNatural will return true if this hand is a natural blackjack (21 from two cards).
func (h *Hand) IsNatural() bool {
	if h.IsInitialHand() {
		// a copy is used so the cards will not stay revealed
		local := NewHand(h.Cards...)

		// we need to reveal all the cards so this works for the dealer
		local.Cards[0].Hidden = false
		local.Cards[1].Hidden = false

		return local.Total() == 21
	}

	return false
}

// String will return a string representation of the hand.
func (h *Hand) String() string {
	var s string
	for _, v := range h.Cards {
		s += fmt.Sprintf("%v  ", v)
	}
	return s
}
