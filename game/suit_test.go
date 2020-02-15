package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Suit_String(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		expected    string
	}{
		{
			description: "Suit_Invalid",
			suit:        0,
			expected:    "X",
		},
		{
			description: "Suit_Clubs",
			suit:        SuitClubs,
			expected:    "♣",
		},
		{
			description: "Suit_Spades",
			suit:        SuitSpades,
			expected:    "♠",
		},
		{
			description: "Suit_Hearts",
			suit:        SuitHearts,
			expected:    "♥",
		},
		{
			description: "Suit_Diamonds",
			suit:        SuitDiamonds,
			expected:    "♦",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			result := tc.suit.String()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
