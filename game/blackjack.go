package game

// Blackjack is the engine for a game of Blackjack.
type Blackjack struct {
	shuffler Shuffler
}

// NewBlackjack will create a new game engine.
func NewBlackjack(numDecks int) *Blackjack {
	shuffler := NewShuffler()

	for i := 0; i < numDecks; i++ {
		deck := NewDeck()
		shuffler.Add(deck.Cards...)
	}

	return &Blackjack{
		shuffler: shuffler,
	}
}

// Update is the main loop for the game engine.
func (b *Blackjack) Update() bool {
	// TODO: this is the game engine main loop, return true to keep looping
	return false
}
