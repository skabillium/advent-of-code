package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntArrayHeap [][]int

func (h IntArrayHeap) Len() int           { return len(h) }
func (h IntArrayHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntArrayHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntArrayHeap) Push(x interface{}) {
	item := x.([]int)
	*h = append(*h, item)
}

func (h *IntArrayHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func getKey(block []int) string {
	strs := make([]string, len(block))
	for i, val := range block {
		strs[i] = strconv.Itoa(val)
	}
	return strings.Join(strs, ",")
}

func findPath(grid []string) {
	seen := map[string]bool{}
	pq := &IntArrayHeap{{0, 0, 0, 0, 0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		block := heap.Pop(pq).([]int)
		hl, r, c, dr, dc, n := block[0], block[1], block[2], block[3], block[4], block[5]

		if r == len(grid)-1 && c == len(grid[0])-1 {
			fmt.Println(hl)
			break
		}

		key := getKey(block)
		isSeen, found := seen[key]
		if isSeen || found {
			continue
		}
		seen[key] = true

		if n < 3 && (dr != 0 && dc != 0) {
			nr := r + dr
			nc := c + dc
			if (0 <= nr && nr < len(grid)) && (0 <= nc && nc < len(grid[0])) {
				loss, _ := strconv.Atoi(string(grid[nr][nc]))
				heap.Push(pq, []int{hl + loss, nr, nc, dr, dc, n + 1})
			}
		}

		directions := [][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		for _, dir := range directions {
			ndr, ndc := dir[0], dir[1]
			if (ndr != dr && ndc != dc) && (ndr != -dr && ndc != -dc) {
				nr := r + ndr
				nc := c + ndc
				if (0 <= nr && nr < len(grid)) && (0 <= nc && nc < len(grid[0])) {
					loss, _ := strconv.Atoi(string(grid[nr][nc]))
					heap.Push(pq, []int{hl + loss, nr, nc, ndr, ndc, 1})
				}
			}
		}
	}

}

func Part1(input string) int {
	file, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	grid := strings.Split(string(file), "\n")

	findPath(grid)

	return 12
}

func main() {
	const INPUT = "day17-test.txt"

	loss := Part1(INPUT)

	fmt.Println("Part 1:", loss)
}
