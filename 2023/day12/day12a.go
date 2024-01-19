// Description: https://adventofcode.com/2023/day/12
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func count(sequence string, broken []int) int {
	// If no characters remain in the sequence, check if any broken sequence remains.
	// If none remain then it is a valid string
	if sequence == "" {
		if len(broken) == 0 {
			return 1
		} else {
			return 0
		}
	}

	// If no broken sequences remain to be matched, check if any "#" remain in the sequence.
	// If there are then there is no match
	if len(broken) == 0 {
		if strings.Contains(sequence, "#") {
			return 0
		} else {
			return 1
		}
	}

	result := 0
	if slices.Contains([]byte{'.', '?'}, sequence[0]) {
		result += count(sequence[1:], broken)
	}

	// Check if it the start of a block, a block is valid if these conditions are met:
	// 1. There are enough springs left
	// 2. All of the springs within the first n must be broken, where n=broken[0]
	// 3. The next sprint after the block must be operational
	if slices.Contains([]byte{'#', '?'}, sequence[0]) {
		if broken[0] <= len(sequence) && !strings.Contains(sequence[:broken[0]], ".") &&
			(broken[0] == len(sequence) || sequence[broken[0]] != '#') {
			// fmt.Println("Str:", len(sequence), "Br[0]:", broken[0])
			if broken[0] == len(sequence) {
				sequence = ""
			} else {
				sequence = sequence[broken[0]+1:]
			}
			result += count(sequence, broken[1:])
		}
	}

	return result
}

func parse(str string) (string, []int) {
	split := strings.Split(str, " ")
	brokenStr := strings.Split(split[1], ",")

	broken := []int{}
	for _, b := range brokenStr {
		num, _ := strconv.Atoi(b)
		broken = append(broken, num)
	}

	return split[0], broken
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		sequence, broken := parse(scanner.Text())
		sum += count(sequence, broken)
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
