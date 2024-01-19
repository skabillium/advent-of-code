// Description: https://adventofcode.com/2023/day/4
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		parsingWinning := true
		winning := []int{}
		drawn := []int{}
		for i := 8; i < len(line); i++ { // Start from first ":"
			char := string(line[i])
			if char == " " {
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
				if parsingWinning {
					winning = append(winning, num)
				} else {
					drawn = append(drawn, num)
				}
			}

			if char == "|" {
				parsingWinning = false
			}

		}
		lineNum++

		matches := 0
		for _, w := range winning {
			contains := slices.Contains(drawn, w)
			if contains {
				matches++
			}
		}

		if matches != 0 {
			sum += int(math.Pow(2, float64(matches-1)))
		}

	}

	return sum
}

func main() {
	const INPUT = "day4.txt"
	expected := 32001
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
