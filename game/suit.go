package game

// Suit represents a suit on a card.
type Suit int

// Suit types.
const (
	SuitClubs Suit = iota + 1
	SuitSpades
	SuitHearts
	SuitDiamonds
)

// String returns the symbol representation of the suit.
func (s Suit) String() string {
	switch s {
	case SuitClubs:
		return "♣"
	case SuitSpades:
		return "♠"
	case SuitHearts:
		return "♥"
	case SuitDiamonds:
		return "♦"
	default:
		return "X"
	}
}
