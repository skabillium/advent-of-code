// Description: https://adventofcode.com/2023/day/10
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const (
	North = iota
	South
	West
	East
)

type State struct {
	row      int
	col      int
	char     byte
	previous int
}

// Create a state machine that computes input based on the current character and previous position
func getNextState(current State) State {
	switch current.char {
	case '|':
		if current.previous == North {
			return State{row: current.row + 1, col: current.col, previous: North}
		} else {
			return State{row: current.row - 1, col: current.col, previous: South}
		}
	case '-':
		if current.previous == East {
			return State{row: current.row, col: current.col - 1, previous: East}
		} else {
			return State{row: current.row, col: current.col + 1, previous: West}
		}
	case 'L':
		if current.previous == East {
			return State{row: current.row - 1, col: current.col, previous: South}
		} else {
			return State{row: current.row, col: current.col + 1, previous: West}
		}
	case 'J':
		if current.previous == North {
			return State{row: current.row, col: current.col - 1, previous: East}
		} else {
			return State{row: current.row - 1, col: current.col, previous: South}
		}
	case '7':
		if current.previous == South {
			return State{row: current.row, col: current.col - 1, previous: East}
		} else {
			return State{row: current.row + 1, col: current.col, previous: North}
		}
	case 'F':
		if current.previous == South {
			return State{row: current.row, col: current.col + 1, previous: West}
		} else {
			return State{row: current.row + 1, col: current.col, previous: North}
		}
	default:
		return current

	}
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	lineNum := 0
	distances := [][]int{}
	start := State{}

	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == 'S' {
				start.row = lineNum
				start.col = i
				start.char = 'S'
			}
		}
		lines = append(lines, line)
		distances = append(distances, make([]int, len(line)))
		lineNum++
	}

	// Find where it can move
	paths := []State{}
	{
		// Check up
		if start.row != 0 && slices.Index([]byte{'|', '7', 'F'}, lines[start.row-1][start.col]) != -1 {
			paths = append(paths, State{
				row:      start.row - 1,
				col:      start.col,
				char:     lines[start.row-1][start.col],
				previous: South,
			})
		}

		// Check down
		if start.row != len(lines)-1 && slices.Index([]byte{'|', 'L', 'J'}, lines[start.row+1][start.col]) != -1 {
			paths = append(paths, State{
				row:      start.row + 1,
				col:      start.col,
				char:     lines[start.row+1][start.col],
				previous: North,
			})
		}

		// Check left
		if start.col != 0 && slices.Index([]byte{'-', 'L', 'F'}, lines[start.row][start.col-1]) != -1 {
			paths = append(paths, State{
				row:      start.row,
				col:      start.col - 1,
				char:     lines[start.row][start.col-1],
				previous: East,
			})
		}

		// Check right
		if start.col != len(lines[0])-1 && slices.Index([]byte{'-', '7', 'J'}, lines[start.row][start.col+1]) != -1 {
			paths = append(paths, State{
				row:      start.row,
				col:      start.col + 1,
				char:     lines[start.row][start.col+1],
				previous: West,
			})
		}

	}

	dist := [2]int{1, 1}
	maxDist := 0
	for {
		isFinal := true
		for i, state := range paths {
			next := getNextState(state)
			next.char = lines[next.row][next.col]

			if next.char != 'S' {
				isFinal = false
			}

			dist[i]++
			currDist := distances[next.row][next.col]

			if next.char == 'S' {
				distances[state.row][state.col] = 1
			} else if currDist == 0 || dist[i] < currDist {
				distances[next.row][next.col] = dist[i]
			}

			if currDist > maxDist {
				maxDist = currDist
			}

			paths[i] = next
		}

		if isFinal {
			break
		}

	}

	return maxDist
}

func main() {
	const INPUT = "day10.txt"
	expected := 6951
	distance := Part1(INPUT)
	if distance != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", distance, expected)
		os.Exit(1)
	}

	fmt.Println("Part 1:", distance)
}
