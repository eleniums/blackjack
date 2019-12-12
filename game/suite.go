package game

// Suite represents a suite on a card.
type Suite int

// Suite types.
const (
	SuiteClubs Suite = iota
	SuiteSpades
	SuiteHearts
	SuiteDiamonds
)

// String returns the symbol representation of the suite.
func (s Suite) String() string {
	switch s {
	case SuiteClubs:
		return "♣"
	case SuiteSpades:
		return "♠"
	case SuiteHearts:
		return "♥"
	case SuiteDiamonds:
		return "♦"
	default:
		return "invalid suite"
	}
}
