package game

type Suite int

const (
	SuiteClubs Suite = iota
	SuiteSpades
	SuiteHearts
	SuiteDiamonds
)

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
