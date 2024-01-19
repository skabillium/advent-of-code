// Description: https://adventofcode.com/2023/day/1#part2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkDigitName(start int, word string, str string) bool {
	return start+len(word) <= len(str) && word == str[start:start+len(word)]
}

func Part2(input string) int {

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
		for i, c := range line {
			char := string(c)
			switch char {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				digits += char
			case "o":
				if checkDigitName(i, "one", line) {
					digits += "1"
				}
			case "t":
				if checkDigitName(i, "two", line) {
					digits += "2"
				} else if checkDigitName(i, "three", line) {
					digits += "3"
				}
			case "f":
				if checkDigitName(i, "four", line) {
					digits += "4"
				} else if checkDigitName(i, "five", line) {
					digits += "5"
				}
			case "s":
				if checkDigitName(i, "six", line) {
					digits += "6"
				} else if checkDigitName(i, "seven", line) {
					digits += "7"
				}
			case "e":
				if checkDigitName(i, "eight", line) {
					digits += "8"
				}
			case "n":
				if checkDigitName(i, "nine", line) {
					digits += "9"
				}

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
	expected := 54094
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
