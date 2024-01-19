package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(input string) int {
	fileBytes, err := os.ReadFile(input)
	if err != nil {
		panic("No file: " + input)
	}

	lines := strings.Split(string(fileBytes), "\n")
	instructions := lines[0]

	mapp := map[string][2]string{}

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		node := string(line[0:3])
		left := string(line[7:10])
		right := string(line[12:15])

		mapp[node] = [2]string{left, right}
	}

	current := "AAA"
	target := "ZZZ"
	steps := 0

nodeSearch:
	for {
		for _, ins := range instructions {
			if current == target {
				break nodeSearch
			}

			steps++
			instruction := string(ins)
			if instruction == "L" {
				current = mapp[current][0]
			} else {
				current = mapp[current][1]
			}

		}
	}

	return steps

}

// Description: https://adventofcode.com/2023/day/8
func main() {
	const INPUT = "day8.txt"
	expected := 21409
	steps := Part1(INPUT)
	if steps != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", steps, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", steps)

}
