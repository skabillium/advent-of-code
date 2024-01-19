// Description: https://adventofcode.com/2023/day/5
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	SeedRow = iota
	SoilRow
	FertilizerRow
	WaterRow
	LightRow
	TempRow
	HumidityRow
	LocationRow
)

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func isLetter(char string) bool {
	return "a" <= char && char <= "z" || "A" <= char && char <= "Z"
}

func parseWords(str string) []string {
	words := []string{}

	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if isLetter(char) {
			start := i
			end := i
			for isLetter(char) {
				end++
				if i == len(str)-1 {
					break
				}
				i++
				char = string(str[i])
			}
			words = append(words, str[start:end])
		}
	}

	return words
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

	seeds := parseNumbers(lines[0])

	// Initialize almanac
	almanac := [8][]int{}
	almanac[SeedRow] = seeds
	for i := 1; i < 8; i++ {
		almanac[i] = make([]int, len(seeds))
	}

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if isLetter(string(line[0])) {

			words := parseWords(line)
			mapFrom := words[0]

			mappings := [][]int{}

			i++
			line = lines[i]

			for isDigit(string(line[0])) {
				mappings = append(mappings, parseNumbers(line))
				if i == len(lines)-1 || len(lines[i+1]) == 0 {
					break
				}

				i++
				line = lines[i]
			}

			var row int
			switch mapFrom {
			case "seed":
				row = SeedRow
			case "soil":
				row = SoilRow
			case "fertilizer":
				row = FertilizerRow
			case "water":
				row = WaterRow
			case "light":
				row = LightRow
			case "temperature":
				row = TempRow
			case "humidity":
				row = HumidityRow

			}

			for j, num := range almanac[row] {
				foundInMapping := false
				for _, mapping := range mappings {
					fromRangeStart := mapping[1]
					toRangeStart := mapping[0]
					rangeLen := mapping[2] - 1
					// toRangeEnd := toRangeStart + rangeLen
					fromRangeEnd := fromRangeStart + rangeLen
					if fromRangeStart <= num && num <= fromRangeEnd {
						offset := num - fromRangeStart
						almanac[row+1][j] = toRangeStart + offset
						foundInMapping = true
					}
				}

				if !foundInMapping {
					almanac[row+1][j] = num
				}
			}

		}

	}

	return slices.Min(almanac[LocationRow])
}

func main() {
	const INPUT = "day5.txt"
	expected := 157211394
	min := Part1(INPUT)
	if min != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", min, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", min)
}
