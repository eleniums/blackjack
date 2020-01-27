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

// Action returns the action the player wants to make with his hand from the given array of possible actions.
func (ai *Standard) Action(dealer *game.Hand, player *game.Hand, actions []game.Action) game.Action {
	dealerTotal := dealer.Total()
	playerTotal := player.Total()
	soft := player.Soft()
	hard := player.Hard()

	// check for surrender
	if allowed(actions, game.ActionSurrender) {
		if hard && playerTotal == 16 && within(dealerTotal, 9, 11) && player.Cards[0].Rank() != 8 {
			return game.ActionSurrender
		}
		if hard && playerTotal == 15 && dealerTotal == 10 {
			return game.ActionSurrender
		}
	}

	// check for split
	if allowed(actions, game.ActionSplit) {
		if player.Cards[0].Rank() == game.RankAce {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 8 {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 2 && within(dealerTotal, 2, 7) {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 3 && within(dealerTotal, 2, 7) {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 4 && within(dealerTotal, 5, 6) {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 6 && within(dealerTotal, 2, 6) {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 7 && within(dealerTotal, 2, 7) {
			return game.ActionSplit
		}
		if player.Cards[0].Rank() == 9 && within(dealerTotal, 2, 9) && dealerTotal != 7 {
			return game.ActionSplit
		}
	}

	// check for double
	if allowed(actions, game.ActionDouble) {
		if hard && playerTotal == 9 && within(dealerTotal, 3, 6) {
			return game.ActionDouble
		}
		if hard && playerTotal == 10 && !within(dealerTotal, 10, 11) {
			return game.ActionDouble
		}
		if hard && playerTotal == 11 && dealerTotal != 11 {
			return game.ActionDouble
		}
		if soft && within(playerTotal, 13, 14) && within(dealerTotal, 5, 6) {
			return game.ActionDouble
		}
		if soft && within(playerTotal, 15, 16) && within(dealerTotal, 4, 6) {
			return game.ActionDouble
		}
		if soft && within(playerTotal, 17, 18) && within(dealerTotal, 3, 6) {
			return game.ActionDouble
		}
	}

	// check for hit or stand
	if hard && playerTotal <= 11 {
		return game.ActionHit
	}
	if hard && playerTotal == 12 && within(dealerTotal, 4, 6) {
		return game.ActionStay
	}
	if hard && within(playerTotal, 13, 16) && within(dealerTotal, 2, 6) {
		return game.ActionStay
	}
	if hard && playerTotal >= 17 {
		return game.ActionStay
	}
	if soft && playerTotal <= 17 {
		return game.ActionHit
	}
	if soft && playerTotal == 18 && !within(dealerTotal, 9, 11) {
		return game.ActionStay
	}
	if soft && playerTotal >= 19 {
		return game.ActionStay
	}

	return game.ActionHit
}

// PlaceBet returns the player's bet.
func (ai *Standard) PlaceBet(minBet, maxBet, totalMoney float64) float64 {
	return minBet
}
