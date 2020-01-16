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
