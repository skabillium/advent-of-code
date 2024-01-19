// Description: https://adventofcode.com/2023/day/14#part2
package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func toString(platform [][]byte) string {
	strs := make([]string, len(platform))
	for i, bytes := range platform {
		strs[i] = string(bytes)
	}

	return strings.Join(strs, "\n")
}

func cycle(platform [][]byte) [][]byte {
	// Tilt North
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

	// Tilt West
	for i, row := range platform {
		for j := 0; j < len(row); j++ {
			rock := row[j]
			if rock == '.' || rock == '#' {
				continue
			}

			left := j - 1
			r := j
			for {
				if left < 0 {
					break
				}
				if platform[i][left] == '.' {
					platform[i][r] = '.'
					platform[i][left] = 'O'
				} else {
					break
				}

				left--
				r--
			}
		}
	}

	// Tilt South
	for i := len(platform) - 1; i > -1; i-- {
		row := platform[i]
		for j := 0; j < len(row); j++ {
			rock := row[j]
			if rock == '.' || rock == '#' {
				continue
			}

			down := i + 1
			r := i
			for {
				if down > len(platform)-1 {
					break
				}

				if platform[down][j] == '.' {
					platform[r][j] = '.'
					platform[down][j] = 'O'
				} else {
					break
				}

				down++
				r++
			}
		}
	}

	// Tilt East
	for i, row := range platform {
		for j := len(row) - 1; j > -1; j-- {
			rock := row[j]
			if rock == '.' || rock == '#' {
				continue
			}

			right := j + 1
			r := j
			for {
				if right > len(row)-1 {
					break
				}
				if platform[i][right] == '.' {
					platform[i][r] = '.'
					platform[i][right] = 'O'
				} else {
					break
				}

				right++
				r++
			}
		}
	}

	return platform
}

func Part2(input string) int {
	filebytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Convert slice of strings to slice of byte slices (byte matrix) to flip chars more easily
	platform := [][]byte{}
	lines := strings.Split(string(filebytes), "\n")
	for _, line := range lines {
		platform = append(platform, []byte(line))
	}

	results := []string{toString(platform)}

	// Perform the North->West->South->East cycles unti you find a repeated value.
	// After that we know the period l and we can calculate the result based
	// on the offset from the start of the pattern.
	i := 0
	found := -1
	for {
		res := toString(cycle(platform))
		found = slices.Index(results, res)

		if found != -1 {
			break
		}

		results = append(results, res)
		i++
	}

	// Note for me: Add one to the length of the cycle to adjust for the last element
	// If we did not do that, then when the mod result would be 0 the array element would be the
	// First element instead of the last
	// Max offset = cycle length
	index := found + ((1_000_000_000 - found) % (i - found + 1))
	final := strings.Split(results[index], "\n")

	load := 0
	for i, row := range final {
		load += strings.Count(string(row), "O") * (len(final) - i)
	}

	return load
}

func main() {
	const INPUT = "day14.txt"
	expected := 96447
	load := Part2(INPUT)
	if load != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", load, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", load)

}
