package game

import (
	"strings"
)

// Action a player can take.
type Action int

// Actions a player can take.
const (
	ActionHit = iota + 1
	ActionStay
	ActionSplit
	ActionDouble
	ActionSurrender
	ActionStats
	ActionExit
	ActionInvalid
)

var actionStrings = map[Action]string{
	ActionHit:       "HIT",
	ActionStay:      "STAY",
	ActionSplit:     "SPLIT",
	ActionDouble:    "DOUBLE",
	ActionSurrender: "SURRENDER",
	ActionStats:     "STATS",
	ActionExit:      "EXIT",
	ActionInvalid:   "INVALID",
}

// ParseAction will return the Action represented by the given string.
func ParseAction(value string) Action {
	upper := strings.ToUpper(value)
	for k, v := range actionStrings {
		if v == upper {
			return k
		}
	}
	return ActionInvalid
}

// String will return the string representation of an Action.
func (a Action) String() string {
	s, ok := actionStrings[a]
	if !ok {
		return actionStrings[ActionInvalid]
	}
	return s
}
