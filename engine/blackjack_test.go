package engine

import (
	"testing"

	"github.com/eleniums/blackjack/game"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Blackjack_NewBlackjack(t *testing.T) {
	// arrange
	numDecks := 5
	maxDiscard := 20
	minBet := 3.0
	maxBet := 5.0
	dealer := &Player{Name: "Dealer"}
	players := []*Player{
		&Player{Name: "Player 1"},
		&Player{Name: "Player 2"},
	}

	// act
	blackjack := NewBlackjack(numDecks, maxDiscard, minBet, maxBet, dealer, players...)

	// assert
	assert.NotNil(t, blackjack)
	assert.NotNil(t, blackjack.shuffler)
	assert.Equal(t, dealer, blackjack.dealer)
	assert.NotNil(t, blackjack.discard)
	assert.Empty(t, blackjack.discard.Cards)
	assert.Equal(t, maxDiscard, blackjack.maxDiscard)
	assert.Equal(t, minBet, blackjack.minBet)
	assert.Equal(t, maxBet, blackjack.maxBet)

	for i, v := range blackjack.players {
		assert.Equal(t, players[i], v)
	}
}

func Test_Unit_Blackjack_determineWinner(t *testing.T) {
	testCases := []struct {
		description string
		player      *Player
		hand        *game.Hand
		dealerTotal int
		win         int
		loss        int
		tie         int
		money       float64
	}{
		{
			description: "Surrendered",
			player:      &Player{},
			hand: &game.Hand{
				Bet:         10,
				Surrendered: true,
			},
			dealerTotal: 1,
			win:         0,
			loss:        1,
			tie:         0,
			money:       -10,
		},
		{
			description: "Natural_Blackjack",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, game.RankJack),
					game.NewCard(game.SuitClubs, game.RankAce),
				},
			},
			dealerTotal: 1,
			win:         1,
			loss:        0,
			tie:         0,
			money:       15,
		},
		{
			description: "Player_Bust",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, 9),
					game.NewCard(game.SuitClubs, 8),
					game.NewCard(game.SuitClubs, 5),
				},
			},
			dealerTotal: 1,
			win:         0,
			loss:        1,
			tie:         0,
			money:       -10,
		},
		{
			description: "Dealer_Bust",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, 2),
				},
			},
			dealerTotal: 22,
			win:         1,
			loss:        0,
			tie:         0,
			money:       10,
		},
		{
			description: "Player_Loss",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, game.RankJack),
					game.NewCard(game.SuitClubs, game.RankQueen),
				},
			},
			dealerTotal: 21,
			win:         0,
			loss:        1,
			tie:         0,
			money:       -10,
		},
		{
			description: "Player_Win",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, 9),
					game.NewCard(game.SuitClubs, 8),
					game.NewCard(game.SuitClubs, 4),
				},
			},
			dealerTotal: 20,
			win:         1,
			loss:        0,
			tie:         0,
			money:       10,
		},
		{
			description: "Push",
			player:      &Player{},
			hand: &game.Hand{
				Bet: 10,
				Cards: []game.Card{
					game.NewCard(game.SuitClubs, 9),
					game.NewCard(game.SuitClubs, 8),
					game.NewCard(game.SuitClubs, 4),
				},
			},
			dealerTotal: 21,
			win:         0,
			loss:        0,
			tie:         1,
			money:       0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			blackjack := &Blackjack{}

			// act
			blackjack.determineWinner(tc.player, tc.hand, tc.dealerTotal)

			// assert
			assert.Equal(t, tc.win, tc.player.Win)
			assert.Equal(t, tc.loss, tc.player.Loss)
			assert.Equal(t, tc.tie, tc.player.Tie)
			assert.Equal(t, tc.money, tc.player.Money)
		})
	}
}
