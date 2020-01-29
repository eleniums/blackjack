package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Hand_IsInitialHand(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    bool
	}{
		{
			description: "Two_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
			},
			expected: true,
		},
		{
			description: "Less_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
			},
			expected: false,
		},
		{
			description: "More_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.IsInitialHand()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}
