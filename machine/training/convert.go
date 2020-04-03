package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eleniums/blackjack/machine"
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

		converted := fmt.Sprintf("%v,%v,%v", machine.ConvertResult(parsed[0]), machine.ConvertHand(parsed[1]), machine.ConvertHand(parsed[2]))
		fmt.Println(converted)

		output.WriteString(converted + "\n")
	}
}
