// Description: https://adventofcode.com/2023/day/7#part2
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

	return Hand{cards: cards, bid: bid, handType: handType}
}

func parseHandType(cards string) int {
	cardCounts := map[string]int{}
	for _, c := range cards {
		cardCounts[string(c)]++
	}

	jokers := cardCounts["J"]

	switch len(cardCounts) {
	case 5:
		if jokers != 0 {
			return OnePair
		}
		return HighCard
	case 4:
		if jokers != 0 {
			return ThreeOfAKind
		}
		return OnePair
	case 3:
		// Could be TwoPair, ThreeOfAKind
		max := 0
		for _, count := range cardCounts {
			if count > max {
				max = count
			}
		}

		if max == 3 {
			if jokers != 0 {
				return FourOfAKind
			} else {
				return ThreeOfAKind
			}
		} else {
			if jokers == 0 {
				return TwoPair
			} else if jokers == 1 {
				return FullHouse
			} else {
				return FourOfAKind
			}
		}

	case 2:
		// Could be FourOfAkind, FullHouse
		if jokers != 0 {
			return FiveOfAKind
		}

		max := 0
		for _, count := range cardCounts {
			if count > max {
				max = count
			}
		}

		if max == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}

	case 1:
		return FiveOfAKind

	}

	return HighCard
}

func compareHands(a Hand, b Hand) int {
	CardOrder := []string{
		"J",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
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

func Part2(input string) int {
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
	expected := 251224870
	winnings := Part2(INPUT)
	if winnings != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", winnings, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", winnings)
}
