// Description: https://adventofcode.com/2023/day/6
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT = "src/day06/day6.txt"

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func parseNumbers(str string) []int {
	numbers := []int{}

	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if isDigit(char) {
			start := i
			end := i
			for isDigit(char) {
				end++
				if i == len(str)-1 {
					break
				}
				i++
				char = string(str[i])
			}

			num, _ := strconv.Atoi(str[start:end])
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func Part1(input string) int {
	fileBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	times := parseNumbers(lines[0])
	distances := parseNumbers(lines[1])

	prod := 1
	for i, time := range times {
		waysToWin := 0
		recordDistance := distances[i]
		for j := 0; j < time; j++ {
			// TODO: Could count until distance falls below record
			buttonTime := j
			remainingTime := time - buttonTime
			distance := remainingTime * buttonTime
			if distance > recordDistance {
				waysToWin++
			}
		}

		prod *= waysToWin
	}

	return prod
}

func main() {
	const INPUT = "day6.txt"
	expected := 1195150
	prod := Part1(INPUT)
	if prod != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", prod, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", prod)
}
