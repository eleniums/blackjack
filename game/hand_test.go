package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Hand_NewHand(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    int
	}{
		{
			description: "No_Cards",
			cards:       []Card{},
			expected:    0,
		},
		{
			description: "One_Card",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
			},
			expected: 1,
		},
		{
			description: "Two_Cards",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitClubs, RankAce),
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			hand := NewHand(tc.cards...)

			// assert
			assert.NotNil(t, hand)
			assert.Equal(t, tc.expected, len(hand.Cards))
			for i, v := range hand.Cards {
				assert.Equal(t, tc.cards[i], v)
			}
		})
	}
}

func Test_Unit_Hand_Count(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    int
	}{
		{
			description: "No_Cards",
			cards:       []Card{},
			expected:    0,
		},
		{
			description: "One_Card",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
			},
			expected: 1,
		},
		{
			description: "Two_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
			},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.Count()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_Clear(t *testing.T) {
	// arrange
	hand := NewHand(NewCard(SuitHearts, RankJack), NewCard(SuitHearts, RankJack))
	hand.Bet = 15
	hand.Surrendered = true

	// act
	hand.Clear()

	// assert
	assert.Equal(t, 0, len(hand.Cards))
	assert.Equal(t, 0.0, hand.Bet)
	assert.False(t, hand.Surrendered)
}

func Test_Unit_Hand_Add(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		card        Card
		expected    int
	}{
		{
			description: "No_Cards",
			cards:       []Card{},
			card:        NewCard(SuitClubs, RankAce),
			expected:    1,
		},
		{
			description: "One_Card",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
			},
			card:     NewCard(SuitSpades, RankQueen),
			expected: 2,
		},
		{
			description: "Two_Cards",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankJack),
			},
			card:     NewCard(SuitHearts, RankQueen),
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			hand.Add(tc.card)

			// assert
			assert.Equal(t, tc.expected, len(hand.Cards))
			assert.Equal(t, tc.card, hand.Cards[len(hand.Cards)-1])
		})
	}
}

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

func Test_Unit_Hand_Soft(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    bool
	}{
		{
			description: "No_Aces_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankKing),
			},
			expected: false,
		},
		{
			description: "No_Aces_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 6),
			},
			expected: false,
		},
		{
			description: "No_Aces_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 7),
			},
			expected: false,
		},
		{
			description: "Ace_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 9),
			},
			expected: true,
		},
		{
			description: "Ace_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankJack),
			},
			expected: true,
		},
		{
			description: "Ace_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 4),
				NewCard(SuitHearts, 7),
			},
			expected: false,
		},
		{
			description: "Double_Aces_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 8),
			},
			expected: true,
		},
		{
			description: "Double_Aces_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 9),
			},
			expected: true,
		},
		{
			description: "Double_Aces_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
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
			result := hand.Soft()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_Hard(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    bool
	}{
		{
			description: "No_Aces_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, RankKing),
			},
			expected: true,
		},
		{
			description: "No_Aces_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 6),
			},
			expected: true,
		},
		{
			description: "No_Aces_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankJack),
				NewCard(SuitHearts, 5),
				NewCard(SuitHearts, 7),
			},
			expected: true,
		},
		{
			description: "Ace_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 9),
			},
			expected: false,
		},
		{
			description: "Ace_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankJack),
			},
			expected: false,
		},
		{
			description: "Ace_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 4),
				NewCard(SuitHearts, 7),
			},
			expected: true,
		},
		{
			description: "Double_Aces_Less_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 8),
			},
			expected: false,
		},
		{
			description: "Double_Aces_Equals_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, 9),
			},
			expected: false,
		},
		{
			description: "Double_Aces_Greater_Than_21",
			cards: []Card{
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankAce),
				NewCard(SuitHearts, RankJack),
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.Hard()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_Total(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    int
	}{
		{
			description: "No_Cards",
			cards:       []Card{},
			expected:    0,
		},
		{
			description: "Ace",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
			},
			expected: 11,
		},
		{
			description: "Two_Aces",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitClubs, RankAce),
			},
			expected: 12,
		},
		{
			description: "Natural_Blackjack",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitClubs, RankQueen),
			},
			expected: 21,
		},
		{
			description: "Ace_Over_21",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitClubs, 9),
				NewCard(SuitClubs, 2),
			},
			expected: 12,
		},
		{
			description: "Ace_Under_21",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitClubs, 7),
				NewCard(SuitClubs, 2),
			},
			expected: 20,
		},
		{
			description: "No_Ace_Over_21",
			cards: []Card{
				NewCard(SuitClubs, RankJack),
				NewCard(SuitClubs, 9),
				NewCard(SuitClubs, 3),
			},
			expected: 22,
		},
		{
			description: "No_Ace_Under_21",
			cards: []Card{
				NewCard(SuitClubs, RankKing),
				NewCard(SuitClubs, 8),
				NewCard(SuitClubs, 2),
			},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.Total()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Hand_String(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    string
	}{
		{
			description: "No_Cards",
			cards:       []Card{},
			expected:    "",
		},
		{
			description: "Single_Card",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
			},
			expected: "A♣  ",
		},
		{
			description: "Multiple_Cards",
			cards: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitHearts, RankQueen),
				NewCard(SuitDiamonds, RankKing),
			},
			expected: "A♣  Q♥  K♦  ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			hand := NewHand(tc.cards...)

			// act
			result := hand.String()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
