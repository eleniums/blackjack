package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Card_NewCard(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		rank        Rank
	}{
		{
			description: "Clubs_King",
			suit:        SuitClubs,
			rank:        RankKing,
		},
		{
			description: "Spades_Queen",
			suit:        SuitSpades,
			rank:        RankQueen,
		},
		{
			description: "Hearts_Jack",
			suit:        SuitHearts,
			rank:        RankJack,
		},
		{
			description: "Diamonds_Ace",
			suit:        SuitDiamonds,
			rank:        RankAce,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			card := NewCard(tc.suit, tc.rank)

			// assert
			assert.NotNil(t, card)
			assert.Equal(t, tc.suit, card.suit)
			assert.Equal(t, tc.rank, card.rank)
		})
	}
}

func Test_Unit_Card_Suit(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		hidden      bool
		expected    Suit
	}{
		{
			description: "Clubs_Visible",
			suit:        SuitClubs,
			hidden:      false,
			expected:    SuitClubs,
		},
		{
			description: "Clubs_Hidden",
			suit:        SuitClubs,
			hidden:      true,
			expected:    0,
		},
		{
			description: "Spades_Visible",
			suit:        SuitSpades,
			hidden:      false,
			expected:    SuitSpades,
		},
		{
			description: "Spades_Hidden",
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
			description: "Ace_Visible",
			rank:        RankAce,
			hidden:      false,
			expected:    RankAce,
		},
		{
			description: "Ace_Hidden",
			rank:        RankAce,
			hidden:      true,
			expected:    0,
		},
		{
			description: "King_Visible",
			rank:        RankKing,
			hidden:      false,
			expected:    RankKing,
		},
		{
			description: "King_Hidden",
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

func Test_Unit_Card_Text(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		rank        Rank
		hidden      bool
		expected    string
	}{
		{
			description: "Clubs_Ace_Visible",
			suit:        SuitClubs,
			rank:        RankAce,
			hidden:      false,
			expected:    "A♣",
		},
		{
			description: "Clubs_Ace_Hidden",
			suit:        SuitClubs,
			rank:        RankAce,
			hidden:      true,
			expected:    "XX",
		},
		{
			description: "Spades_Jack_Visible",
			suit:        SuitSpades,
			rank:        RankJack,
			hidden:      false,
			expected:    "J♠",
		},
		{
			description: "Hearts_Queen_Visible",
			suit:        SuitHearts,
			rank:        RankQueen,
			hidden:      false,
			expected:    "Q♥",
		},
		{
			description: "Diamonds_King_Visible",
			suit:        SuitDiamonds,
			rank:        RankKing,
			hidden:      false,
			expected:    "K♦",
		},
		{
			description: "Diamonds_Two_Visible",
			suit:        SuitDiamonds,
			rank:        2,
			hidden:      false,
			expected:    "2♦",
		},
		{
			description: "Diamonds_Ten_Visible",
			suit:        SuitDiamonds,
			rank:        10,
			hidden:      false,
			expected:    "10♦",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			card := NewCard(tc.suit, tc.rank)
			card.Hidden = tc.hidden

			// act
			result := card.Text()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_Unit_Card_Symbol(t *testing.T) {
	testCases := []struct {
		description string
		suit        Suit
		rank        Rank
		hidden      bool
		expected    string
	}{
		{
			description: "Clubs_Ace_Visible",
			suit:        SuitClubs,
			rank:        RankAce,
			hidden:      false,
			expected:    "🃑",
		},
		{
			description: "Clubs_Ace_Hidden",
			suit:        SuitClubs,
			rank:        RankAce,
			hidden:      true,
			expected:    "🂠",
		},
		{
			description: "Spades_Jack_Visible",
			suit:        SuitSpades,
			rank:        RankJack,
			hidden:      false,
			expected:    "🂫",
		},
		{
			description: "Hearts_Queen_Visible",
			suit:        SuitHearts,
			rank:        RankQueen,
			hidden:      false,
			expected:    "🂽",
		},
		{
			description: "Diamonds_King_Visible",
			suit:        SuitDiamonds,
			rank:        RankKing,
			hidden:      false,
			expected:    "🃎",
		},
		{
			description: "Diamonds_Two_Visible",
			suit:        SuitDiamonds,
			rank:        2,
			hidden:      false,
			expected:    "🃂",
		},
		{
			description: "Diamonds_Ten_Visible",
			suit:        SuitDiamonds,
			rank:        10,
			hidden:      false,
			expected:    "🃊",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			card := NewCard(tc.suit, tc.rank)
			card.Hidden = tc.hidden

			// act
			result := card.Symbol()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
