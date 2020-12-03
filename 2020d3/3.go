package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
)

func main() {
	grid := advent.InputLines(advent.MustGetArg(1))
	dxs := []int{1, 3, 5, 7, 1}
	dys := []int{1, 1, 1, 1, 2}
	var results []int
	for i := range dxs {
		dx := dxs[i]
		dy := dys[i]
		x := 0
		y := 0
		treesTotal := 0
		for y < len(grid) {
			if grid[y][x] == '#' {
				treesTotal++
			}
			x = (x + dx) % len(grid[0])
			y += dy
		}
		results = append(results, treesTotal)
	}

	ans := results[0]
	for i := 1; i < len(results); i++ {
		ans *= results[i]
	}
	fmt.Println("part 2:", ans)
}
