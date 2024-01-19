package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func isStartingNode(node string) bool {
	return string(node[len(node)-1]) == "A"
}

func isFinalNode(node string) bool {
	return string(node[len(node)-1]) == "Z"
}

func Part2(input string) int {
	fileBytes, err := os.ReadFile(input)
	if err != nil {
		panic("No file: " + input)
	}

	lines := strings.Split(string(fileBytes), "\n")
	instructions := lines[0]

	mapp := map[string][2]string{}

	nodes := []string{}

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		node := string(line[0:3])
		left := string(line[7:10])
		right := string(line[12:15])

		mapp[node] = [2]string{left, right}

		if isStartingNode(node) {
			nodes = append(nodes, node)
		}
	}

	steps := make([]int, len(nodes))

	for i, node := range nodes {
		st := 0
	nodeSearch:
		for {
			for _, ins := range instructions {
				if isFinalNode(node) {
					break nodeSearch
				}
				st++
				instruction := string(ins)
				if instruction == "L" {
					node = mapp[node][0]
				} else {
					node = mapp[node][1]
				}

			}
		}

		steps[i] = st
	}

	// Find LCM from steps, could also be via the Euclidean algorithm,
	// see: https://en.wikipedia.org/wiki/Least_common_multiple
	slices.Sort(steps)
	largest := steps[len(steps)-1]
	lcm := largest

	for {
		found := true
		for _, step := range steps {
			if lcm%step != 0 {
				found = false
				break
			}
		}
		if found == true {
			break
		}
		lcm += largest
	}

	return lcm

}

// Description: https://adventofcode.com/2023/day/8#part2
func main() {
	const INPUT = "day8.txt"
	expected := 21165830176709
	steps := Part2(INPUT)

	if steps != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", steps, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", steps)

}
