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
