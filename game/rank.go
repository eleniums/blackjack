package game

import (
	"strconv"
)

type Rank int

const (
	RankAce  Rank = 1
	RankJack Rank = iota + 10
	RankQueen
	RankKing
)

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
