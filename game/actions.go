package game

// Action a player can take.
type Action int

// Actions a player can take.
const (
	// Hit to draw another card.
	Hit = iota + 1

	// Stay to stop drawing cards.
	Stay

	// Split a hand.
	Split

	// Double a bet and draw one more card.
	Double
)
