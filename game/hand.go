package game

// Hand represents a hand of cards.
type Hand struct {
	cards []Card
}

// NewHand will create a new hand with the given cards.
func NewHand(cards ...Card) *Hand {
	var hand Hand

	for _, v := range cards {
		hand.cards = append(hand.cards, v)
	}

	return &hand
}

// Count will return the number of cards in the hand.
func (h *Hand) Count() int {
	return len(h.cards)
}

// Cards will return the cards in the hand.
func (h *Hand) Cards() []Card {
	return h.cards
}

// Add a single card to the hand.
func (h *Hand) Add(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) Total() int {
	// TODO: return total value of cards in hand
	return 0
}

func (h *Hand) CanHit() bool {
	// TODO: return true if not busted
	return false
}
