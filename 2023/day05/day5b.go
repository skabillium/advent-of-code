// Description: https://adventofcode.com/2023/day/5#part2
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func isLetter(char string) bool {
	return "a" <= char && char <= "z" || "A" <= char && char <= "Z"
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

func Part2(input string) int {
	fileBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	locations := [][]int{}
	fileMappings := [][][]int{}

	lines := strings.Split(string(fileBytes), "\n")
	seeds := parseNumbers(lines[0])
	ranges := [][]int{}

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			continue
		}

		if isLetter(string(line[0])) {

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

			fileMappings = append(fileMappings, mappings)

		}

	}

	for j := 0; j < len(seeds); j += 2 {
		ranges = [][]int{{seeds[j], seeds[j] + seeds[j+1]}}
		results := [][]int{}
		for _, mappings := range fileMappings {
			for len(ranges) > 0 {
				rang := ranges[len(ranges)-1]
				ranges = ranges[:len(ranges)-1]
				startRange := rang[0]
				endRange := rang[1]
				found := false

				for _, mapp := range mappings {
					target := mapp[0]
					startMap := mapp[1]
					r := mapp[2]

					endMap := startMap + r
					offset := target - startMap

					// Skip if not overlapping at all
					if endMap <= startRange || endRange <= startMap {
						continue
					}

					if startRange < startMap {
						ranges = append(ranges, []int{startRange, startMap})
						startRange = startMap
					}

					if endMap < endRange {
						ranges = append(ranges, []int{endMap, endRange})
						endRange = endMap
					}

					results = append(results, []int{startRange + offset, endRange + offset})
					found = true
					break
				}

				if !found {
					results = append(results, []int{startRange, endRange})
				}
			}
			ranges = results
			results = nil
		}

		locations = append(locations, ranges...)
	}

	min := math.MaxInt
	for _, loc := range locations {
		// Only check start range for minimum
		if loc[0] < min {
			min = loc[0]
		}
	}

	return min
}

func main() {
	const INPUT = "day5.txt"
	expected := 50855035
	min := Part2(INPUT)
	if min != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", min, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", min)
}
