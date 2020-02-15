package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Shuffler_NewShuffler(t *testing.T) {
	// act
	shuffler := NewShuffler()

	// assert
	assert.NotNil(t, shuffler)
	assert.Empty(t, shuffler.deck)
}

func Test_Unit_Shuffler_Add(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		count       int
		add         []Card
		expected    []Card
	}{
		{
			description: "Empty_Deck_Add_Card",
			cards:       []Card{},
			count:       1,
			add: []Card{
				NewCard(SuitClubs, RankAce),
			},
			expected: []Card{
				NewCard(SuitClubs, RankAce),
			},
		},
		{
			description: "Empty_Deck_Add_Cards",
			cards:       []Card{},
			count:       3,
			add: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			expected: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
		},
		{
			description: "Multiple_Cards_Add_Card",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			count: 4,
			add: []Card{
				NewCard(SuitClubs, RankAce),
			},
			expected: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
		},
		{
			description: "Multiple_Cards_Add_Cards",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			count: 6,
			add: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitDiamonds, 2),
				NewCard(SuitHearts, 3),
			},
			expected: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitDiamonds, 2),
				NewCard(SuitHearts, 3),
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			shuffler := NewShuffler()
			shuffler.deck.Cards = tc.cards

			// act
			shuffler.Add(tc.add...)

			// assert
			assert.Len(t, shuffler.deck.Cards, tc.count)
			for _, v := range tc.expected {
				assert.Contains(t, shuffler.deck.Cards, v)
			}
		})
	}
}

func Test_Unit_Shuffler_Deal(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		err         error
		expected    Card
	}{
		{
			description: "Empty_Deck",
			cards:       []Card{},
			err:         ErrDeckEmpty,
			expected:    Card{},
		},
		{
			description: "Single_Card",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
			},
			err:      nil,
			expected: NewCard(SuitDiamonds, RankJack),
		},
		{
			description: "Multiple_Cards",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			err:      nil,
			expected: NewCard(SuitDiamonds, RankJack),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			shuffler := NewShuffler()
			shuffler.deck.Cards = tc.cards

			// act
			result, err := shuffler.Deal()

			// assert
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
