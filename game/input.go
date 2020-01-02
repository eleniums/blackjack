package game

import (
	"bufio"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

// ReadInput will read a string from stdin.
func ReadInput() string {
	input, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
}
