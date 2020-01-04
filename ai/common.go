package ai

// within will return true if the value is within the inclusive range [low, high].
func within(value, low, high int) bool {
	return value >= low && value <= high
}
