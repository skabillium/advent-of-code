// Description: https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
		digits := ""
		for _, c := range line {
			char := string(c)
			if char >= "0" && char <= "9" {
				digits += char
			}
		}

		value := string(digits[0]) + string(digits[len(digits)-1])

		valueInt, _ := strconv.Atoi(value)
		sum += valueInt
	}

	return sum

}

func main() {
	const INPUT = "day1.txt"
	expected := 54968
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
