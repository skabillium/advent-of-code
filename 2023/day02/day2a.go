// Description: https://adventofcode.com/2023/day/2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Draw contains the number of cubes in each draw of a game
type Draw struct {
	red   int
	green int
	blue  int
}

func isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	target := Draw{
		red:   12,
		green: 13,
		blue:  14,
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		gameStart := strings.Index(line, ":")
		gameId, _ := strconv.Atoi(line[5:gameStart])
		draws := strings.Split(line[gameStart+2:], ";")
		isValid := true
		for _, d := range draws {
			draw := Draw{
				red:   0,
				green: 0,
				blue:  0,
			}
			d = strings.TrimSpace(d)
			for _, cubes := range strings.Split(d, ",") {
				val := strings.Split(strings.TrimSpace(cubes), " ")
				color := val[1]
				num, _ := strconv.Atoi(val[0])

				switch color {
				case "red":
					draw.red = num
				case "green":
					draw.green = num
				case "blue":
					draw.blue = num
				}

			}

			if draw.red > target.red || draw.green > target.green || draw.blue > target.blue {
				isValid = false
			}

		}

		if isValid {
			sum += gameId
		}

	}

	return sum
}

func main() {
	const INPUT = "day2.txt"
	expected := 1931
	sum := Part1(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", sum)
}
