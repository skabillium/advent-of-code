// Description: https://adventofcode.com/2023/day/3#part2
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	number int
	line   int
	start  int
	end    int
}

type Gear struct {
	line   int
	column int
}

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func findOverlappingNums(nums []Number, line int, from int, to int) []Number {
	overlapping := []Number{}
	for _, n := range nums {
		if n.line == line {
			if from <= n.end && n.start <= to {
				overlapping = append(overlapping, n)
			}
		}
	}

	return overlapping
}

func Part2(input string) int {
	f, err := os.ReadFile(input)
	fstr := string(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(fstr, "\n")
	numbers := []Number{}
	gears := []Gear{}

	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if isDigit(char) {
				start := j
				end := j
				numStr := ""
				for isDigit(char) {
					numStr += char
					end = j
					if j == len(line)-1 {
						break
					}
					j++
					char = string(line[j])
				}
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, Number{number: num, line: i, start: start, end: end})
			}

			if char == "*" {
				gears = append(gears, Gear{line: i, column: j})
			}
		}
	}

	sum := 0
	for _, gear := range gears {
		overlapping := []Number{}
		from := 0
		to := 0
		// Search up
		if gear.line != 0 {
			if gear.column == 0 {
				from = gear.column
			} else {
				from = gear.column - 1
			}

			if gear.column == len(lines[gear.line])-1 {
				to = gear.column
			} else {
				to = gear.column + 1
			}
			overlapping = append(overlapping, findOverlappingNums(numbers, gear.line-1, from, to)...)
		}

		// Search down
		if gear.line != len(lines)-1 {
			if gear.column == 0 {
				from = gear.column
			} else {
				from = gear.column - 1
			}

			if gear.column == len(lines[gear.line])-1 {
				to = gear.column
			} else {
				to = gear.column + 1
			}
			overlapping = append(overlapping, findOverlappingNums(numbers, gear.line+1, from, to)...)
		}

		// Seach left
		if gear.column != 0 {
			from = gear.column - 1
			to = from
			overlapping = append(overlapping, findOverlappingNums(numbers, gear.line, from, to)...)
		}

		// Search right
		if gear.column != len(lines[gear.line]) {
			from = gear.column + 1
			to = from
			overlapping = append(overlapping, findOverlappingNums(numbers, gear.line, from, to)...)
		}

		if len(overlapping) == 2 {
			sum += overlapping[0].number * overlapping[1].number
		}

	}

	return sum
}

func main() {
	const INPUT = "day3.txt"
	expected := 84495585
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
