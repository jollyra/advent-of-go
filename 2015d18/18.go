package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"github.com/jollyra/go-grid"
)

type p = grid.Point

var on = byte('#')
var off = byte('.')

func isCorner(point p) bool {
	corners := []p{p{0, 0}, p{0, 99}, p{99, 0}, p{99, 99}}
	for _, c := range corners {
		if c.Equals(point) {
			return true
		}
	}
	return false
}

func update(grid *grid.Grid) *grid.Grid {
	nextGrid := grid.Copy()
	for y := range grid.Grid {
		for x := range grid.Grid[y] {

			if isCorner(p{x, y}) {
				grid.Grid[y][x] = on
				continue
			}

			ns := grid.Neighbours8(x, y)
			var numOn int
			for _, n := range ns {
				if grid.Grid[n.Y][n.X] == on {
					numOn++
				}
			}

			if grid.Grid[y][x] == on {
				if numOn != 2 && numOn != 3 {
					nextGrid.Grid[y][x] = off
				}
			} else {
				if numOn == 3 {
					nextGrid.Grid[y][x] = on
				}
			}

		}
	}

	return nextGrid
}

func gridFromLines(lines []string) *grid.Grid {
	grid := grid.NewGrid(len(lines), len(lines))
	for y := range lines {
		for x := range lines[y] {
			grid.Grid[y][x] = lines[y][x]
		}
	}
	return grid
}

func countOnLights(g *grid.Grid) int {
	count := 0
	for y := range g.Grid {
		for x := range g.Grid[y] {
			if g.Grid[y][x] == on {
				count++
			}
		}
	}
	return count
}

func animationLoop(grid *grid.Grid, steps int) *grid.Grid {
	for i := 0; i < steps; i++ {
		grid = update(grid)
	}
	return grid
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	grid := gridFromLines(lines)
	fmt.Println(grid)
	grid = animationLoop(grid, 100)
	fmt.Println(grid)
	fmt.Println(countOnLights(grid))
}
