// Description: https://adventofcode.com/2023/day/15#part2
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	label    string
	focalLen int
}

type Box []Lens

func remove(box Box, s int) Box {
	return append(box[:s], box[s+1:]...)
}

func hash(str string) int {
	value := 0
	for i := 0; i < len(str); i++ {
		c := str[i]
		value += int(c)
		value *= 17
		value %= 256
	}

	return value
}

func Part2(input string) int {

	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	steps := strings.Split(string(file), ",")

	boxes := make([]Box, 256)

	for _, step := range steps {
		if strings.ContainsRune(step, '=') {
			command := strings.Split(step, "=")
			label := command[0]
			focalLen, _ := strconv.Atoi(command[1])

			lens := Lens{label: command[0], focalLen: focalLen}

			boxIdx := hash(label)

			lensIdx := slices.IndexFunc(boxes[boxIdx], func(l Lens) bool { return l.label == label })
			if lensIdx != -1 {
				// Lens found
				boxes[boxIdx][lensIdx] = lens
			} else {
				// Not found
				boxes[boxIdx] = append(boxes[boxIdx], lens)
			}

		} else {
			label := strings.Split(step, "-")[0]
			boxIdx := hash(label)
			lensIdx := slices.IndexFunc(boxes[boxIdx], func(l Lens) bool { return l.label == label })

			if lensIdx != -1 {
				boxes[boxIdx] = remove(boxes[boxIdx], lensIdx)
			}
		}
	}

	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focalLen
		}
	}

	return sum
}

func main() {
	const INPUT = "day15.txt"
	expected := 260530
	sum := Part2(INPUT)
	if sum != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", sum, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", sum)
}
