package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Card_Suit(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		hidden      bool
		expected    Suit
	}{
		{
			description: "Suit_Clubs_Visible",
			suit:        SuitClubs,
			hidden:      false,
			expected:    SuitClubs,
		},
		{
			description: "Suit_Clubs_Hidden",
			suit:        SuitClubs,
			hidden:      true,
			expected:    0,
		},
		{
			description: "Suit_Spades_Visible",
			suit:        SuitSpades,
			hidden:      false,
			expected:    SuitSpades,
		},
		{
			description: "Suit_Spades_Hidden",
			suit:        SuitSpades,
			hidden:      true,
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			card := NewCard(tc.suit, 0)
			card.Hidden = tc.hidden

			// act
			result := card.Suit()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_Unit_Card_Rank(t *testing.T) {
	testCases := []struct {
		description string
		rank        Rank
		hidden      bool
		expected    Rank
	}{
		{
			description: "Rank_Ace_Visible",
			rank:        RankAce,
			hidden:      false,
			expected:    RankAce,
		},
		{
			description: "Rank_Ace_Hidden",
			rank:        RankAce,
			hidden:      true,
			expected:    0,
		},
		{
			description: "Rank_King_Visible",
			rank:        RankKing,
			hidden:      false,
			expected:    RankKing,
		},
		{
			description: "Rank_King_Hidden",
			rank:        RankKing,
			hidden:      true,
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			card := NewCard(0, tc.rank)
			card.Hidden = tc.hidden

			// act
			result := card.Rank()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
