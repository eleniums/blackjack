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
		return "C"
	case SuiteSpades:
		return "S"
	case SuiteHearts:
		return "H"
	case SuiteDiamonds:
		return "D"
	default:
		return "invalid suite"
	}
}
