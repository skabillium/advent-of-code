// Description: https://adventofcode.com/2023/day/16#part2
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

var shift = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func Part2(input string) int {
	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	entrypoints := []Beam{}
	for i := 0; i < len(lines); i++ {
		entrypoints = append(entrypoints, Beam{row: i, col: 0, dir: 1})
		entrypoints = append(entrypoints, Beam{row: i, col: len(lines) - 1, dir: 3})
	}

	for i := 0; i < len(lines[0]); i++ {
		entrypoints = append(entrypoints, Beam{row: 0, col: i, dir: 2})
		entrypoints = append(entrypoints, Beam{row: len(lines[0]) - 1, col: i, dir: 0})
	}

	max := 0
	for _, entry := range entrypoints {
		// Initialize a 2d array for the energized map
		energized := make([][]byte, len(lines))
		for i, line := range lines {
			energized[i] = make([]byte, len(line))
			for j := range line {
				energized[i][j] = '.'
			}
		}

		queue := []Beam{entry}
		looped := map[string]bool{}

		for len(queue) > 0 {
			last := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			row, col, dir := last.row, last.col, last.dir

			key := strings.Join([]string{strconv.Itoa(row), strconv.Itoa(col), strconv.Itoa(dir)}, ",")
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

		if total > max {
			max = total
		}
	}

	return max
}

func main() {
	const INPUT = "day16.txt"
	expected := 7313
	max := Part2(INPUT)
	if max != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", max, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", max)
}
