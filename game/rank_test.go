package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Rank_String(t *testing.T) {
	testCases := []struct {
		description string
		rank        Rank
		expected    string
	}{
		{
			description: "Rank_Ace",
			rank:        RankAce,
			expected:    "A",
		},
		{
			description: "Rank_Two",
			rank:        2,
			expected:    "2",
		},
		{
			description: "Rank_Ten",
			rank:        10,
			expected:    "10",
		},
		{
			description: "Rank_Jack",
			rank:        RankJack,
			expected:    "J",
		},
		{
			description: "Rank_Queen",
			rank:        RankQueen,
			expected:    "Q",
		},
		{
			description: "Rank_King",
			rank:        RankKing,
			expected:    "K",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			result := tc.rank.String()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
