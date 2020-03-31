package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/eleniums/blackjack/game"
)

// Simple program to transform the generated human-readable training data to integers more appropriate for machine learning.
// Usage: go run ./machine/training/convert.go ./machine/testdata/training.csv ./machine/testdata/output.csv
func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), ",")
		fmt.Println(parsed)

		converted := fmt.Sprintf("%v,%v,%v", convertResult(parsed[0]), convertHand(parsed[1]), convertHand(parsed[2]))
		fmt.Println(converted)

		output.WriteString(converted + "\n")
	}
}

func convertResult(result string) int {
	split := strings.Split(result, "_")

	a := game.ParseAction(split[0])
	r := game.ParseResult(split[1])

	return int(a) + int(r)*100
}

func convertHand(hand string) int {
	total := 0.0
	split := strings.Split(hand, " ")

	for i, v := range split {
		var n int
		switch v {
		case "J", "Q", "K":
			n = 10
		case "A":
			n = 11
		default:
			n, _ = strconv.Atoi(v)
		}
		total += float64(n) * math.Pow(100.0, float64(i))
	}

	return int(total)
}
