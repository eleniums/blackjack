package ai

import (
	"testing"

	"github.com/eleniums/blackjack/game"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Common_within(t *testing.T) {
	testCases := []struct {
		description string
		value       int
		low         int
		high        int
		expected    bool
	}{
		{
			description: "Less_Than",
			value:       1,
			low:         2,
			high:        4,
			expected:    false,
		},
		{
			description: "Greater_Than",
			value:       5,
			low:         2,
			high:        4,
			expected:    false,
		},
		{
			description: "Lower_Bound",
			value:       2,
			low:         2,
			high:        4,
			expected:    true,
		},
		{
			description: "Upper_Bound",
			value:       4,
			low:         2,
			high:        4,
			expected:    true,
		},
		{
			description: "Between_Bounds",
			value:       3,
			low:         2,
			high:        4,
			expected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			result := within(tc.value, tc.low, tc.high)

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}

func Test_Unit_Common_allowed(t *testing.T) {
	testCases := []struct {
		description string
		actions     []game.Action
		action      game.Action
		expected    bool
	}{
		{
			description: "Empty_Actions",
			actions:     []game.Action{},
			action:      game.ActionDouble,
			expected:    false,
		},
		{
			description: "Action_Allowed",
			actions: []game.Action{
				game.ActionHit,
				game.ActionDouble,
				game.ActionStay,
			},
			action:   game.ActionDouble,
			expected: true,
		},
		{
			description: "Action_Not_Allowed",
			actions: []game.Action{
				game.ActionHit,
				game.ActionStay,
			},
			action:   game.ActionDouble,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// act
			result := allowed(tc.actions, tc.action)

			// assert
			assert.Equal(t, tc.expected, result)
		})
	}
}
