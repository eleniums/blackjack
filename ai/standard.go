package ai

import (
	"github.com/eleniums/blackjack/game"
)

// Standard is an opponent that uses a standard strategy.
type Standard struct{}

// NewStandard will create a new standard AI.
func NewStandard() *Standard {
	return &Standard{}
}

// Action returns the action the player wants to make with his hand.
func (ai *Standard) Action(dealer *game.Hand, player *game.Hand) game.Action {
	dealerTotal := dealer.Total()
	playerTotal := player.Total()

	// at least one ace that is counted as 11
	if player.Soft() {
		if playerTotal >= 20 {
			return game.ActionStay
		}

		// TODO: finish soft

		return game.ActionHit
	}

	// no aces in hand
	if playerTotal >= 17 {
		return game.ActionStay
	}
	if ai.within(playerTotal, 13, 16) && ai.within(dealerTotal, 2, 6) {
		return game.ActionStay
	}
	if playerTotal == 12 && ai.within(dealerTotal, 4, 6) {
		return game.ActionStay
	}
	if playerTotal == 11 && ai.canDouble(player) {
		return game.ActionDouble
	}
	if playerTotal == 10 && ai.within(dealerTotal, 2, 9) && ai.canDouble(player) {
		return game.ActionDouble
	}
	if playerTotal == 9 && ai.within(dealerTotal, 3, 6) && ai.canDouble(player) {
		return game.ActionDouble
	}

	return game.ActionHit
}

// PlaceBet returns the player's bet.
func (ai *Standard) PlaceBet(minBet, maxBet, totalMoney int) int {
	return minBet
}

// within will return true if the value is within the inclusive range [low, high].
func (ai *Standard) within(value, low, high int) bool {
	return value >= low && value <= high
}

// canDouble will return true if the player is allowed to double down.
func (ai *Standard) canDouble(hand *game.Hand) bool {
	return hand.Count() == 2
}
