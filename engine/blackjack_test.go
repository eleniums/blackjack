package engine

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Blackjack_NewBlackjack(t *testing.T) {
	// arrange
	numDecks := 5
	maxDiscard := 20
	minBet := 3.0
	maxBet := 5.0
	dealer := &Player{Name: "Dealer"}
	players := []*Player{
		&Player{Name: "Player 1"},
		&Player{Name: "Player 2"},
	}

	// act
	blackjack := NewBlackjack(numDecks, maxDiscard, minBet, maxBet, dealer, players...)

	// assert
	assert.NotNil(t, blackjack)
	assert.NotNil(t, blackjack.shuffler)
	assert.Equal(t, dealer, blackjack.dealer)
	assert.NotNil(t, blackjack.discard)
	assert.Empty(t, blackjack.discard.Cards)
	assert.Equal(t, maxDiscard, blackjack.maxDiscard)
	assert.Equal(t, minBet, blackjack.minBet)
	assert.Equal(t, maxBet, blackjack.maxBet)

	for i, v := range blackjack.players {
		assert.Equal(t, players[i], v)
	}
}
