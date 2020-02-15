package game

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Deck_NewDeck(t *testing.T) {
	// act
	deck := NewDeck()

	// assert
	assert.NotNil(t, deck)
	assert.Len(t, deck.Cards, 52)
	for i := 1; i <= 13; i++ {
		expected := NewCard(SuitClubs, Rank(i))
		assert.Equal(t, expected, deck.Cards[i-1])
	}
	for i := 1; i <= 13; i++ {
		expected := NewCard(SuitSpades, Rank(i))
		assert.Equal(t, expected, deck.Cards[i+12])
	}
	for i := 1; i <= 13; i++ {
		expected := NewCard(SuitHearts, Rank(i))
		assert.Equal(t, expected, deck.Cards[i+25])
	}
	for i := 1; i <= 13; i++ {
		expected := NewCard(SuitDiamonds, Rank(i))
		assert.Equal(t, expected, deck.Cards[i+38])
	}
}

func Test_Unit_Deck_NewEmptyDeck(t *testing.T) {
	// act
	deck := NewEmptyDeck()

	// assert
	assert.NotNil(t, deck)
	assert.Empty(t, deck.Cards)
}

func Test_Unit_Deck_Add(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		index       int
		card        Card
		expected    []Card
	}{
		{
			description: "Empty_Deck",
			cards:       []Card{},
			index:       0,
			card:        NewCard(SuitClubs, RankAce),
			expected: []Card{
				NewCard(SuitClubs, RankAce),
			},
		},
		{
			description: "Multiple_Cards_Add_Top",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			index: 0,
			card:  NewCard(SuitClubs, RankAce),
			expected: []Card{
				NewCard(SuitClubs, RankAce),
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
		},
		{
			description: "Multiple_Cards_Add_Bottom",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			index: 3,
			card:  NewCard(SuitClubs, RankAce),
			expected: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
				NewCard(SuitClubs, RankAce),
			},
		},
		{
			description: "Multiple_Cards_Add_Middle",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			index: 1,
			card:  NewCard(SuitClubs, RankAce),
			expected: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitClubs, RankAce),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			deck := &Deck{
				Cards: tc.cards,
			}

			// act
			deck.Add(tc.index, tc.card)

			// assert
			assert.Equal(t, tc.expected, deck.Cards)
		})
	}
}

func Test_Unit_Deck_Count(t *testing.T) {
	testCases := []struct {
		description string
		cards       []Card
		expected    int
	}{
		{
			description: "Empty_Deck",
			cards:       []Card{},
			expected:    0,
		},
		{
			description: "Single_Card",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
			},
			expected: 1,
		},
		{
			description: "Multiple_Cards",
			cards: []Card{
				NewCard(SuitDiamonds, RankJack),
				NewCard(SuitSpades, RankQueen),
				NewCard(SuitHearts, RankKing),
			},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			deck := &Deck{
				Cards: tc.cards,
			}

			// act
			result := deck.Count()

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Deck_Deal(t *testing.T) {
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
			deck := &Deck{
				Cards: tc.cards,
			}

			// act
			result, err := deck.Deal()

			// assert
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Deck_String(t *testing.T) {
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
			deck := &Deck{
				Cards: tc.cards,
			}

			// act
			result := deck.String()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}
