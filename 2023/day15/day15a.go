// Description: https://adventofcode.com/2023/day/15
package main

import (
	"fmt"
	"os"
	"strings"
)

// Hash algorithm:
// - Determine the ASCII code for the current character of the string.
// - Increase the current value by the ASCII code you just determined.
// - Set the current value to itself multiplied by 17.
// - Set the current value to the remainder of dividing itself by 256.

func hash(str string) int {
	value := 0
	for i := 0; i < len(str); i++ {
		c := str[i]
		value += int(c)
		value *= 17
		value %= 256
	}

	return value
}

func Part1(input string) int {

	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	steps := strings.Split(string(file), ",")

	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}

	return sum
}

func main() {
	const INPUT = "day15.txt"
	expected := 519041
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
