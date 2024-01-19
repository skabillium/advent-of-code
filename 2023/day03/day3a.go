// Description: https://adventofcode.com/2023/day/3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	number int
	line   int
	start  int
	end    int
}

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func isSymbol(char string) bool {
	return char != "." && !isDigit(char)
}

func checkSymbol(chars string) bool {
	for _, c := range chars {
		char := string(c)
		if isSymbol(char) {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []Number{}
	content := []string{}
	lineNum := 0

	// sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if isDigit(char) {
				start := i
				end := i
				numStr := ""
				for isDigit(char) {
					numStr += char
					end = i
					if i == len(line)-1 {
						break
					}
					i++
					char = string(line[i])
				}
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, Number{number: num, line: lineNum, start: start, end: end})
			}
		}

		lineNum++
		content = append(content, line)
	}

	sum := 0

	// TODO: Can skip len checks for diagonals, go strings work weird: https://stackoverflow.com/questions/12311033/extracting-substrings-in-go
	for _, number := range numbers {
		if number.line != 0 {

			if checkSymbol(content[number.line-1][number.start : number.end+1]) {
				sum += number.number
				continue
			}

		}

		if number.line != len(content)-1 {
			// Check down
			if checkSymbol(content[number.line+1][number.start : number.end+1]) {
				sum += number.number
				continue
			}
		}

		if number.start != 0 {
			// Check left
			if checkSymbol(string(content[number.line][number.start-1])) {
				sum += number.number
				continue
			}
		}

		if number.end != len(content[number.line])-1 {
			// Check right
			if checkSymbol(string(content[number.line][number.end+1])) {
				sum += number.number
				continue
			}
		}

		if number.line != 0 && number.start != 0 {
			// Check up & left
			if isSymbol(string(content[number.line-1][number.start-1])) {
				sum += number.number
				continue
			}
		}

		if number.line != 0 && number.end != len(content[number.line])-1 {
			// Check up & right
			if isSymbol(string(content[number.line-1][number.end+1])) {
				sum += number.number
				continue
			}
		}

		if number.line != len(content)-1 && number.start != 0 {
			// Check down & left
			if isSymbol(string(content[number.line+1][number.start-1])) {
				sum += number.number
				continue
			}
		}

		if number.line != len(content)-1 && number.end != len(content[number.line])-1 {
			// Check down & right
			if isSymbol(string(content[number.line+1][number.end+1])) {
				sum += number.number
				continue
			}
		}
	}

	return sum
}

func main() {
	const INPUT = "day3.txt"
	expected := 544664
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
