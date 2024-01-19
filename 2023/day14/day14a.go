// Description: https://adventofcode.com/2023/day/14
package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(input string) int {
	filebytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Convert slice of strings to slice of byte slices (byte matrix)
	platform := [][]byte{}
	lines := strings.Split(string(filebytes), "\n")
	for _, line := range lines {
		platform = append(platform, []byte(line))
	}

	for i, row := range platform {
		for j := 0; j < len(row); j++ {
			rock := row[j]
			if rock == '.' || rock == '#' {
				continue
			}

			up := i - 1
			r := i
			for {
				if up < 0 {
					break
				}

				if platform[up][j] == '.' {
					platform[r][j] = '.'
					platform[up][j] = 'O'
				} else {
					break
				}

				up--
				r--
			}
		}
	}

	load := 0
	for i, row := range platform {
		load += strings.Count(string(row), "O") * (len(platform) - i)
	}

	return load
}

func main() {
	const INPUT = "day14.txt"
	expected := 108614
	load := Part1(INPUT)
	if load != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", load, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", load)
}
