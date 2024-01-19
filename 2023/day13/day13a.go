// Description: https://adventofcode.com/2023/day/13
package main

import (
	"fmt"
	"os"
	"strings"
)

func splitPatterns(lines []string) [][]string {
	patterns := [][]string{}
	pattern := []string{}

	for i, line := range lines {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = []string{}
			continue
		}

		pattern = append(pattern, line)

		if i == len(lines)-1 {
			patterns = append(patterns, pattern)
		}

	}

	return patterns
}

func Part1(input string) int {
	filebytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(filebytes), "\n")
	patterns := splitPatterns(lines)

	sum := 0

	for _, pattern := range patterns {
		// Check rows
		for i := 0; i < len(pattern)-1; i++ {
			current := pattern[i]
			next := pattern[i+1]
			if current == next {
				reflection := i
				isReflection := false
				low := i - 1
				high := i + 2
				for {
					if low < 0 || high > len(pattern)-1 {
						isReflection = true
						break
					}
					if pattern[low] == pattern[high] {
						low--
						high++
					} else {
						isReflection = false
						break
					}
				}

				if isReflection {
					sum += (reflection + 1) * 100
				}
			}
		}

		// Check columns
		for i := 0; i < len(pattern[0])-1; i++ {
			current := pattern[0][i]
			next := pattern[0][i+1]

			if current == next {
				reflection := i
				isReflection := false
				low := i
				high := i + 1
				for {
					if low < 0 || high > len(pattern[0])-1 {
						isReflection = true
						break
					}

					checked := true
					for _, line := range pattern {
						if line[low] != line[high] {
							checked = false
						}
					}

					if checked {
						low--
						high++
					} else {
						isReflection = false
						break
					}
				}
				if isReflection {
					sum += reflection + 1
				}
			}
		}

	}

	return sum
}

func main() {
	const INPUT = "day13.txt"
	expected := 37561
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
