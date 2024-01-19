// Description: https://adventofcode.com/2023/day/16
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Beam struct {
	row int
	col int
	dir int
}

func Part1(input string) int {
	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	// Initialize a 2d array for the energized map
	energized := make([][]byte, len(lines))
	for i, line := range lines {
		energized[i] = make([]byte, len(line))
		for j := range line {
			energized[i][j] = '.'
		}
	}

	shift := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	queue := []Beam{
		{row: 0, col: 0, dir: 1},
	}
	looped := map[string]bool{}

	for len(queue) > 0 {
		last := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		row, col, dir := last.row, last.col, last.dir

		key := strings.Join([]string{strconv.Itoa(last.row), strconv.Itoa(last.col), strconv.Itoa(last.dir)}, ",")
		if looped[key] {
			continue
		}
		looped[key] = true

		if row >= 0 && row < len(lines) && col >= 0 && col < len(lines[0]) {
			directions := []int{}
			energized[last.row][last.col] = '#'

			switch lines[row][col] {
			case '|':
				if dir%2 != 0 {
					directions = append(directions, 0, 2)
				} else {
					directions = append(directions, dir)
				}
			case '-':
				if dir%2 == 0 {
					directions = append(directions, 1, 3)
				} else {
					directions = append(directions, dir)
				}
			case '/':
				directions = append(directions, dir^1)
			case '\\':
				directions = append(directions, dir^3)
			default:
				directions = append(directions, dir)
			}

			for _, d := range directions {
				queue = append(queue, Beam{
					row: row + shift[d][0],
					col: col + shift[d][1],
					dir: d,
				})
			}
		}

	}

	total := 0
	for _, row := range energized {
		for _, cell := range row {
			if cell == '#' {
				total++
			}
		}
	}

	return total
}

func main() {
	const INPUT = "day16.txt"
	expected := 7046
	total := Part1(INPUT)
	if total != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", total, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", total)
}
