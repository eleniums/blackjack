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
		if playerTotal == 19 {
			if dealerTotal == 6 && player.CanDouble() {
				return game.ActionDouble
			}
			return game.ActionStay
		}
		if playerTotal == 18 {
			if within(dealerTotal, 2, 6) && player.CanDouble() {
				return game.ActionDouble
			}
			if within(dealerTotal, 9, 11) {
				return game.ActionHit
			}
			return game.ActionStay
		}
		if playerTotal == 17 {
			if within(dealerTotal, 3, 6) && player.CanDouble() {
				return game.ActionDouble
			}
		}
		if within(playerTotal, 15, 16) {
			if within(dealerTotal, 4, 6) && player.CanDouble() {
				return game.ActionDouble
			}
		}
		if within(playerTotal, 13, 14) {
			if within(dealerTotal, 5, 6) && player.CanDouble() {
				return game.ActionDouble
			}
		}

		return game.ActionHit
	}

	// no aces in hand
	if playerTotal >= 17 {
		return game.ActionStay
	}
	if within(playerTotal, 13, 16) && within(dealerTotal, 2, 6) {
		return game.ActionStay
	}
	if playerTotal == 12 && within(dealerTotal, 4, 6) {
		return game.ActionStay
	}
	if playerTotal == 11 && player.CanDouble() {
		return game.ActionDouble
	}
	if playerTotal == 10 && within(dealerTotal, 2, 9) && player.CanDouble() {
		return game.ActionDouble
	}
	if playerTotal == 9 && within(dealerTotal, 3, 6) && player.CanDouble() {
		return game.ActionDouble
	}

	return game.ActionHit
}

// PlaceBet returns the player's bet.
func (ai *Standard) PlaceBet(minBet, maxBet, totalMoney int) int {
	return minBet
}
