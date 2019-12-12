package game

import (
	"strconv"
)

// Rank represents a rank (the value) on a card.
type Rank int

// Rank types.
const (
	RankAce  Rank = 1
	RankJack Rank = iota + 10
	RankQueen
	RankKing
)

// String returns the string representation of this rank.
func (r Rank) String() string {
	if r >= 2 && r <= 10 {
		return strconv.Itoa(int(r))
	}

	switch r {
	case RankAce:
		return "A"
	case RankJack:
		return "J"
	case RankQueen:
		return "Q"
	case RankKing:
		return "K"
	default:
		return "invalid rank"
	}
}
