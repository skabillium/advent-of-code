// Description: https://adventofcode.com/2023/day/6#part2
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT = "src/day06/day6.txt"

func parseNumber(str string) int {
	digits := ""
	for _, c := range str {
		char := string(c)
		if "0" <= char && char <= "9" {
			digits += char
		}
	}

	number, _ := strconv.Atoi(digits)
	return number
}

func Part2(input string) int {
	fileBytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	time := parseNumber(lines[0])
	recordDistance := parseNumber(lines[1])

	total := 0
	for j := 0; j < time; j++ {
		buttonTime := j
		remainingTime := time - buttonTime
		distance := remainingTime * buttonTime
		if distance > recordDistance {
			total++
		}
	}

	return total

}

func main() {
	const INPUT = "day6.txt"
	expected := 42550411
	total := Part2(INPUT)
	if total != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", total, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", total)
}
