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

func Part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	lineNum := 0
	distances := [][]int{}
	pipes := [][]bool{}
	start := State{}

	for scanner.Scan() {
		line := scanner.Text()
		pipe := make([]bool, len(line))
		for i, c := range line {
			if c == 'S' {
				start.row = lineNum
				start.col = i
				start.char = 'S'
				pipe[i] = true
			} else {
				pipe[i] = false
			}
		}
		lines = append(lines, line)
		distances = append(distances, make([]int, len(line)))
		pipes = append(pipes, pipe)
		lineNum++
	}

	// Find where it can move
	next := State{}
	{
		// Check up
		if start.row != 0 && slices.Index([]byte{'|', '7', 'F'}, lines[start.row-1][start.col]) != -1 {
			next = State{
				row:      start.row - 1,
				col:      start.col,
				char:     lines[start.row-1][start.col],
				previous: South,
			}
		}

		// Check down
		if start.row != len(lines)-1 && slices.Index([]byte{'|', 'L', 'J'}, lines[start.row+1][start.col]) != -1 {
			next = State{
				row:      start.row + 1,
				col:      start.col,
				char:     lines[start.row+1][start.col],
				previous: North,
			}
		}

		// Check left
		if start.col != 0 && slices.Index([]byte{'-', 'L', 'F'}, lines[start.row][start.col-1]) != -1 {
			next = State{
				row:      start.row,
				col:      start.col - 1,
				char:     lines[start.row][start.col-1],
				previous: East,
			}
		}

		// Check right
		if start.col != len(lines[0])-1 && slices.Index([]byte{'-', '7', 'J'}, lines[start.row][start.col+1]) != -1 {
			next = State{
				row:      start.row,
				col:      start.col + 1,
				char:     lines[start.row][start.col+1],
				previous: West,
			}
		}

	}

	pipes[next.row][next.col] = true

	state := next
	points := []State{start}
	for {
		next = getNextState(state)
		next.char = lines[next.row][next.col]
		pipes[next.row][next.col] = true

		points = append(points, state)

		if next.char == 'S' {
			break
		}

		state = next
	}

	// Count the ammount of vertical segments ('|', 'L', '7', 'F') encountered when
	// traversing the rows of the matrix. If the number of crossings is odd, then a point
	// is inside the polygon. If it is even, it is outside
	between := 0
	verticalPipes := []byte{'|', 'S', '7', 'F'} // Skip 'J' since we're parsing left to right
	for i, pipeline := range pipes {
		verticals := 0
		for j := 0; j < len(pipeline); j++ {
			isLoop := pipeline[j]
			if isLoop && slices.Contains(verticalPipes, lines[i][j]) {
				verticals++
			}

			if !isLoop && j != len(pipeline)-1 && verticals%2 == 1 {
				between++
			}
		}
	}

	return between

}

func main() {
	const INPUT = "day10.txt"
	expected := 563
	between := Part2(INPUT)
	if between != expected {
		fmt.Printf("[ERROR]: Got: %d, Expected: %d\n", between, expected)
		os.Exit(1)
	}

	fmt.Println("Part 2:", between)
}
