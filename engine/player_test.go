package engine

import (
	"testing"

	"github.com/eleniums/blackjack/ai"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Player_NewPlayer(t *testing.T) {
	// arrange
	name := "Test Name"
	money := 0.0
	ai := ai.NewRandom()

	// act
	player := NewPlayer(name, money, ai)

	// assert
	assert.NotNil(t, player)
	assert.Equal(t, name, player.Name)
	assert.Equal(t, ai, player.AI)
	assert.NotNil(t, player.Hand)
	assert.Empty(t, player.SplitHands)
	assert.Equal(t, money, player.Money)
	assert.Zero(t, player.Win)
	assert.Zero(t, player.Loss)
	assert.Zero(t, player.Tie)
}
