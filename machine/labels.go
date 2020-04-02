package machine

import (
	"strings"

	"github.com/eleniums/blackjack/game"
)

// Label in a dataset.
type Label int

var labelStrings = map[Label]string{
	0:  "STAY_WIN",
	1:  "STAY_LOSS",
	2:  "STAY_TIE",
	3:  "HIT_WIN",
	4:  "HIT_LOSS",
	5:  "HIT_TIE",
	6:  "HIT_NONE",
	7:  "DOUBLE_WIN",
	8:  "DOUBLE_LOSS",
	9:  "DOUBLE_TIE",
	10: "SURRENDER_LOSS",
	11: "SPLIT_NONE",
}

// ParseLabel will return the label represented by the given string.
func ParseLabel(value string) Label {
	upper := strings.ToUpper(value)
	for k, v := range labelStrings {
		if v == upper {
			return k
		}
	}
	return -1
}

// Split will divide a label into an action and result.
func (l Label) Split() (game.Action, game.Result) {
	s := l.String()
	split := strings.Split(s, "_")
	a := game.ParseAction(split[0])
	r := game.ParseResult(split[1])
	return a, r
}

// String will return the string representation of a Label.
func (l Label) String() string {
	s, ok := labelStrings[l]
	if !ok {
		return "INVALID_INVALID"
	}
	return s
}
