// Description: https://adventofcode.com/2023/day/2#part2
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

type Game struct {
	id    string
	draws []Draw
}

func isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

func Part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		gameStart := strings.Index(line, ":")
		draws := strings.Split(line[gameStart+2:], ";")

		maxValues := Draw{
			red:   0,
			green: 0,
			blue:  0,
		}

		for _, d := range draws {
			draw := Draw{
				red:   0,
				green: 0,
				blue:  0,
			}
			d = strings.TrimSpace(d)
			for i, cubes := range strings.Split(d, ",") {

				if i == i {

				}

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

			if draw.red > maxValues.red {
				maxValues.red = draw.red
			}

			if draw.green > maxValues.green {
				maxValues.green = draw.green
			}

			if draw.blue > maxValues.blue {
				maxValues.blue = draw.blue
			}

		}

		sum += maxValues.red * maxValues.green * maxValues.blue

	}

	return sum
}

func main() {
	const INPUT = "day2.txt"
	expected := 83105
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
