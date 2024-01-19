// Description: https://adventofcode.com/2023/day/7
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	bid      int
	handType int
}

// Hand types
const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func parseHand(str string) Hand {
	hand := strings.Split(str, " ")
	cards := hand[0]
	bid, _ := strconv.Atoi(hand[1])
	handType := parseHandType(cards)

	return Hand{cards: hand[0], bid: bid, handType: handType}
}

func parseHandType(cards string) int {
	cardCounts := map[string]int{}
	for _, c := range cards {
		cardCounts[string(c)]++
	}

	for _, count := range cardCounts {
		switch count {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			if len(cardCounts) == 2 {
				return FullHouse
			} else {
				return ThreeOfAKind
			}
		case 2:
			if len(cardCounts) == 2 {
				return FullHouse
			} else if len(cardCounts) == 3 {
				return TwoPair
			} else {
				return OnePair
			}
		}

	}

	return HighCard
}

func compareHands(a Hand, b Hand) int {
	CardOrder := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}

	if a.handType != b.handType {
		return a.handType - b.handType
	}

	for i, card := range a.cards {
		cardA := string(card)
		cardB := string(b.cards[i])
		if cardA == cardB {
			continue
		}

		return slices.Index(CardOrder, cardA) - slices.Index(CardOrder, cardB)

	}
	return 0
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic("No file: " + input)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		hand := parseHand(line)
		hands = append(hands, hand)
	}

	slices.SortStableFunc(hands, compareHands)

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}

	return winnings
}

func main() {
	const INPUT = "day7.txt"
	expected := 250347426
	winnings := Part1(INPUT)
	if winnings != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", winnings, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", Part1(INPUT))
}
