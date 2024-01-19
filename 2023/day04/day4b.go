// Description: https://adventofcode.com/2023/day/4#part2
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	num     int
	copies  int
	score   int
	winning []int
	drawn   []int
}

const (
	PhaseGame = iota
	PhaseWinning
	PhaseDrawn
)

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func isLetter(char string) bool {
	return "a" <= char && char <= "z" || "A" <= char && char <= "Z"
}

func countMatches(card Card) int {
	matches := 0
	for _, w := range card.winning {
		contains := slices.Contains(card.drawn, w)
		if contains {
			matches++
		}
	}

	return matches
}

func initializeCardsFromText(text string) []Card {
	cards := []Card{}

	for _, line := range strings.Split(text, "\n") {

		card := Card{}
		parsingPhase := PhaseGame
		for i := 0; i < len(line); i++ { // Start from first ":"
			char := string(line[i])
			if char == " " && isLetter(char) {
				continue
			}

			if isDigit(char) {
				numStr := ""
				// Parse number
				for isDigit(char) {
					numStr += char
					if i == len(line)-1 {
						break
					}
					i++
					char = string(line[i])
				}

				num, _ := strconv.Atoi(numStr)

				switch parsingPhase {
				case PhaseGame:
					card.num = num
				case PhaseWinning:
					card.winning = append(card.winning, num)
				case PhaseDrawn:
					card.drawn = append(card.drawn, num)
				}

			}

			if char == ":" {
				// Start to parse winning numbers
				parsingPhase = PhaseWinning
			}

			if char == "|" {
				parsingPhase = PhaseDrawn
			}

		}

		card.copies = 1
		card.score = countMatches(card)
		cards = append(cards, card)

	}

	return cards
}

func Part2(input string) int {
	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(input)
	}

	cards := initializeCardsFromText(string(bytes))

	copies := 0
	for i := 0; i < len(cards); i++ {
		card := cards[i]
		copies += card.copies
		for j := 1; j <= card.score; j++ {
			cards[i+j].copies += card.copies
		}
	}

	return copies
}

func main() {
	const INPUT = "day4.txt"
	expected := 5037841
	copies := Part2(INPUT)
	if copies != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", copies, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", copies)
}
