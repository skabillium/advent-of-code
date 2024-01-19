// Description: https://adventofcode.com/2023/day/12#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cache = map[string]int{}

func intSliceToString(ints []int) []string {
	strs := []string{}
	for _, i := range ints {
		strs = append(strs, strconv.Itoa(i))
	}

	return strs
}

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

	// Memoize results to avoid re-calculations
	// Example key: "..#?.:2"
	key := sequence + ":" + strings.Join(intSliceToString(broken), ",")
	if result, found := cache[key]; found {
		return result
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
			if broken[0] == len(sequence) {
				sequence = ""
			} else {
				sequence = sequence[broken[0]+1:]
			}
			result += count(sequence, broken[1:])
		}
	}

	cache[key] = result
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

func Part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		repeatedSeq, repeatedBroken := parse(scanner.Text())

		sequences := []string{}
		broken := make([]int, len(repeatedBroken)*5)

		for i := 0; i < 5; i++ {
			sequences = append(sequences, repeatedSeq)
			copy(broken[i*len(repeatedBroken):], repeatedBroken)
		}

		sequence := strings.Join(sequences, "?")
		sum += count(sequence, broken)
	}

	return sum
}

func main() {
	const INPUT = "day12.txt"
	expected := 4546215031609
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
