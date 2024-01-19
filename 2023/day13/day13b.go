// Description: https://adventofcode.com/2023/day/13#part2
package main

import (
	"fmt"
	"os"
	"strings"
)

// Count the differences between 2 strings of equal length
func diff(a string, b string) int {
	d := 0
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d
}

func transpose(strs []string) []string {
	rows := len(strs)
	cols := len(strs[0])

	temp := make([][]byte, cols)
	for i := range temp {
		temp[i] = make([]byte, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp[j][i] = strs[i][j]
		}
	}

	transposed := make([]string, len(temp))
	for i, row := range temp {
		transposed[i] = string(row)
	}

	return transposed
}

// Look at all the possible reflection lines by counting the sum of total differences
// between every line and it's reflection. If at the end the sum is 0 then the horizontal
// line is found
func findMirror(pattern []string) int {
	for i := 0; i < len(pattern)-1; i++ {
		d := diff(pattern[i], pattern[i+1])
		if d <= 1 {
			reflection := i
			low := i - 1
			high := i + 2
			for {
				if low < 0 || high > len(pattern)-1 {
					break
				}

				d += diff(pattern[low], pattern[high])
				low--
				high++
			}
			if d == 1 {
				return reflection + 1
			}
		}
	}

	return 0
}

func Part2(input string) int {
	filebytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	sum := 0
	patterns := strings.Split(string(filebytes), "\n\n")
	for _, p := range patterns {
		pattern := strings.Split(p, "\n")
		sum += findMirror(pattern) * 100

		// Convert the columns to rows and run it again
		pattern = transpose(pattern)
		sum += findMirror(pattern)
	}

	return sum
}

func main() {
	const INPUT = "day13.txt"
	expected := 31108
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
