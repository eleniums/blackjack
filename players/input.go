package players

import (
	"bufio"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

func readInput() string {
	input, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
}
