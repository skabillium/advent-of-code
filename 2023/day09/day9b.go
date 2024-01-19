package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func parseNumbers(str string) []int {
	numbers := []int{}

	for i := 0; i < len(str); i++ {
		char := rune(str[i])
		if unicode.IsDigit(char) || string(char) == "-" {
			start := i
			end := i
			for unicode.IsDigit(char) || string(char) == "-" {
				end++
				if i == len(str)-1 {
					break
				}
				i++
				char = rune(str[i])
			}

			num, _ := strconv.Atoi(str[start:end])
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func Part2(input string) int {

	file, err := os.Open(input)
	if err != nil {
		panic("No file: " + input)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sequences := [][]int{}
		sequences = append(sequences, parseNumbers(scanner.Text()))

		seq := sequences[0]

	createDiffs:
		for {
			next := make([]int, len(seq)-1)
			allZeros := true
			for i := 0; i < len(next); i++ {
				diff := seq[i+1] - seq[i]
				if diff != 0 {
					allZeros = false
				}

				next[i] = diff
			}

			sequences = append(sequences, next)
			if allZeros == true {
				break createDiffs
			}
			seq = next
		}

		// Calculate the previous sequence number by subtracting the next difference
		// from the first number
		for i := len(sequences) - 2; i >= 0; i-- {
			current := sequences[i]
			previous := sequences[i+1]
			nextElement := current[0] - previous[0]
			sequences[i] = append([]int{nextElement}, sequences[i]...)

			if i == 0 {
				sum += sequences[i][0]
			}
		}

	}

	return sum
}

// Description https://adventofcode.com/2023/day/9#part2
func main() {
	const INPUT = "day9.txt"
	expected := 1026

	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
