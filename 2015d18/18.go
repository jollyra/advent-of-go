package main

import (
	"fmt"
	// "github.com/jollyra/numutil"
	"github.com/jollyra/stringutil"
	"os"
	"strings"
)

type grid [100][100]byte

func (g grid) String() string {
	var b strings.Builder
	for _, row := range g {
		for _, x := range row {
			fmt.Fprintf(&b, "%c", x)
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

func (g *grid) Copy() grid {
	copy := grid{}
	for y := range g {
		for x := range g[y] {
			copy[y][x] = g[y][x]
		}
	}
	return copy
}

func (g *grid) isInBounds(x, y int) bool {
	size := len(g)
	if 0 <= x && x < size && 0 <= y && y < size {
		return true
	}
	return false
}

type point struct{ X, Y int }

func (p *point) equals(q point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

func (g *grid) Neighbours8(x, y int) []point {
	points := []point{
		point{x + 1, y},
		point{x + 1, y + 1},
		point{x, y + 1},
		point{x - 1, y + 1},
		point{x - 1, y},
		point{x - 1, y - 1},
		point{x, y - 1},
		point{x + 1, y - 1},
	}

	pointsInBounds := make([]point, 0)
	for _, p := range points {
		if g.isInBounds(p.X, p.Y) {
			pointsInBounds = append(pointsInBounds, p)
		}
	}

	return pointsInBounds
}

func newGrid(lines []string) grid {
	grid := grid{}
	for y := range lines {
		for x := range lines[y] {
			grid[y][x] = lines[y][x]
		}
	}
	return grid
}
func isCorner(p point) bool {
	corners := []point{point{0, 0}, point{0, 99}, point{99, 0}, point{99, 99}}
	for _, c := range corners {
		if c.equals(p) {
			return true
		}
	}
	return false
}

var on = byte('#')
var off = byte('.')

func update(grid grid) grid {
	nextGrid := grid.Copy()
	for y := range grid {
		for x := range grid[y] {

			if isCorner(point{x, y}) {
				grid[y][x] = on
				continue
			}

			ns := grid.Neighbours8(x, y)
			var numOn int
			for _, n := range ns {
				if grid[n.Y][n.X] == on {
					numOn++
				}
			}

			if grid[y][x] == on {
				if numOn != 2 && numOn != 3 {
					nextGrid[y][x] = off
				}
			} else {
				if numOn == 3 {
					nextGrid[y][x] = on
				}
			}

		}
	}

	return nextGrid
}

func countOnLights(g grid) int {
	count := 0
	for y := range g {
		for x := range g[y] {
			if g[y][x] == on {
				count++
			}
		}
	}
	return count
}

func animationLoop(grid grid, steps int) grid {
	for i := 0; i < steps; i++ {
		grid = update(grid)
	}
	return grid
}

func main() {
	lines := stringutil.InputLines(os.Args[1])
	grid := newGrid(lines)
	fmt.Println(grid)
	grid = animationLoop(grid, 100)
	fmt.Println(grid)
	fmt.Println(countOnLights(grid))
}
