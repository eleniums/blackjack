package game

import (
	"bufio"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Input_ReadInput(t *testing.T) {
	testCases := []struct {
		description string
		line        string
		expected    string
	}{
		{
			description: "Newline_Only",
			line:        "\n",
			expected:    "",
		},
		{
			description: "Full_Line",
			line:        "Simple line with spaces.\n",
			expected:    "Simple line with spaces.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// arrange
			stdin = bufio.NewReader(strings.NewReader(tc.line))

			// act
			result := ReadInput()

			// assert
			assert.Equal(t, tc.expected, result)
		})

	}
}

func Test_Unit_Input_ReadInput_Panic(t *testing.T) {
	// arrange
	stdin = bufio.NewReader(strings.NewReader(""))

	defer func() {
		if r := recover(); r != nil {
			assert.True(t, true)
		}
	}()

	// act
	ReadInput()

	// assert
	assert.Fail(t, "expected panic from ReadInput")
}
