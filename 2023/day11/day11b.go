// Description: https://adventofcode.com/2023/day/11#part2
package main

import (
	"fmt"
	"os"
	"strings"
)

type Galaxy struct {
	row int
	col int
}

func absolute(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Part2(input string) int {

	filebytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	emptyRows := []int{}
	emptyCols := []int{}

	galaxies := []Galaxy{}

	lines := strings.Split(string(filebytes), "\n")

	colContains := make([]bool, len(lines[0]))
	for i, line := range lines {
		if strings.Index(line, "#") == -1 {
			emptyRows = append(emptyRows, i)
		}

		for j, col := range line {
			if col == '#' {
				colContains[j] = true
				galaxies = append(galaxies, Galaxy{row: i, col: j})
			}
		}
	}

	for i, col := range colContains {
		if col == false {
			emptyCols = append(emptyCols, i)
		}
	}

	for i, row := range emptyRows {
		offset := i * 999_999
		for j, g := range galaxies {
			if g.row > row+offset {
				galaxies[j].row += 999_999
			}
		}
	}

	for i, col := range emptyCols {
		offset := i * 999_999
		for j, g := range galaxies {
			if g.col > col+offset {
				galaxies[j].col += 999_999
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i; j < len(galaxies); j++ {
			b := galaxies[j]
			sum += absolute(a.row-b.row) + absolute(a.col-b.col)
		}
	}

	return sum
}

func main() {
	const INPUT = "day11.txt"
	expected := 593821230983
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
