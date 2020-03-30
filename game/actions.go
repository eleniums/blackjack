package game

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

// String will return the string representation of an Action.
func (a Action) String() string {
	s, ok := actionStrings[a]
	if !ok {
		return actionStrings[ActionInvalid]
	}
	return s
}
