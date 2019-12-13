package game

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

func (h *Hand) Total() int {
	// TODO: return total value of cards in hand
	return 0
}

func (h *Hand) CanHit() bool {
	// TODO: return true if not busted
	return false
}
