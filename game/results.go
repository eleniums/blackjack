package game

import (
	"strings"
)

// Result of playing a hand.
type Result int

// Results of playing a hand.
const (
	ResultWin = iota + 1
	ResultLoss
	ResultTie
	ResultNone
	ResultInvalid
)

var resultStrings = map[Result]string{
	ResultWin:     "WIN",
	ResultLoss:    "LOSS",
	ResultTie:     "TIE",
	ResultNone:    "NONE",
	ResultInvalid: "INVALID",
}

// ParseResult will return the Result represented by the given string.
func ParseResult(value string) Result {
	upper := strings.ToUpper(value)
	for k, v := range resultStrings {
		if v == upper {
			return k
		}
	}
	return ResultInvalid
}

// String will return the string representation of a Result.
func (a Result) String() string {
	s, ok := resultStrings[a]
	if !ok {
		return resultStrings[ResultInvalid]
	}
	return s
}
