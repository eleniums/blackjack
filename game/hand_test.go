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

func Test_Unit_Hand_CanDouble(t *testing.T) {
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
			result := hand.CanDouble()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_CanSplit(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    bool
	}{
		{
			description: "Two_Matching_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
			},
			expected: true,
		},
		{
			description: "Two_Mismatched_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankKing),
			},
			expected: false,
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
			result := hand.CanSplit()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_IsNatural(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    bool
	}{
		{
			description: "Natural_Blackjack",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankAce),
			},
			expected: true,
		},
		{
			description: "Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankQueen),
			},
			expected: false,
		},
		{
			description: "Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 7),
			},
			expected: false,
		},
		{
			description: "21_With_More_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 6),
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.IsNatural()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}
