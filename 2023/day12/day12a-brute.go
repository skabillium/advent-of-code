// Description: https://adventofcode.com/2023/day/12
// This is a brute force solution, it will be removed once I find a better one
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Recursively generate all posible combinations for a sequence, replacing all "?" with
// "." & "#"
func generateCombinations(s string, index int, currentCombination string, combinations *[]string) {
	if index == len(s) {
		*combinations = append(*combinations, currentCombination)
		return
	}

	if s[index] == '?' {
		generateCombinations(s, index+1, currentCombination+"#", combinations)
		generateCombinations(s, index+1, currentCombination+".", combinations)
	} else {
		generateCombinations(s, index+1, currentCombination+string(s[index]), combinations)
	}
}

// Count all distinct sequences of broken gears and returning a comma separated string
// like the input. This will be used to compare a generated combination with the original
// broken gear sequences to check if it is valid.
func score(sequence string) string {
	score := ""
	count := 0
	for i := 0; i < len(sequence); i++ {
		current := sequence[i]
		if current == '.' {
			if count != 0 {
				score += strconv.Itoa(count) + ","
			}
			count = 0
		}

		if current == '#' {
			count++
			if i == len(sequence)-1 {
				score += strconv.Itoa(count)
			}
		}

	}

	if len(score) != 0 && score[len(score)-1] == ',' {
		score = score[:len(score)-1]
	}

	return score
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		sequence := split[0]
		broken := strings.ReplaceAll(split[1], "\n", "")

		// Generate all possible combinations for a sequence
		var combinations []string
		generateCombinations(sequence, 0, "", &combinations)

		// Count how many combinations match the broken sequences from the input
		valid := 0
		for _, combination := range combinations {
			if score(combination) == broken {
				valid++
			}
		}

		sum += valid
	}

	return sum
}

func main() {
	const INPUT = "day12.txt"
	expected := 6981
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
