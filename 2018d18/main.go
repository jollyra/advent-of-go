package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"github.com/jollyra/go-advent-util/point"
	// "time"
)

type Point = point.Point

type RuneGrid struct {
	yxGrid [][]rune
}

func (g *RuneGrid) Get(x, y int) rune { return g.yxGrid[y][x] }

func (g *RuneGrid) Set(x int, y int, r rune) { g.yxGrid[y][x] = r }

func (g *RuneGrid) Size() (sizeY, sizeX int) { return len((*g).yxGrid), len((*g).yxGrid[0]) }

func (g RuneGrid) Render() {
	for y := range g.yxGrid {
		for x := range g.yxGrid[y] {
			fmt.Printf("%c", g.yxGrid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g RuneGrid) Copy() RuneGrid {
	gp := make([][]rune, 0)
	for y := range g.yxGrid {
		yp := make([]rune, len(g.yxGrid[y]))
		copy(yp, g.yxGrid[y])
		gp = append(gp, yp)
	}
	return RuneGrid{yxGrid: gp}
}

func (g *RuneGrid) IsInBounds(x, y int) bool {
	dy, dx := g.Size()
	if 0 <= x && x < dx && 0 <= y && y < dy {
		return true
	}
	return false
}

func (g *RuneGrid) Neighbours8(x, y int) []Point {
	points := []Point{
		Point{x + 1, y},
		Point{x + 1, y + 1},
		Point{x, y + 1},
		Point{x - 1, y + 1},
		Point{x - 1, y},
		Point{x - 1, y - 1},
		Point{x, y - 1},
		Point{x + 1, y - 1},
	}

	pointsInBounds := make([]Point, 0)
	for _, p := range points {
		if g.IsInBounds(p.X, p.Y) {
			pointsInBounds = append(pointsInBounds, p)
		}
	}

	return pointsInBounds
}

func parseRuneGrid(lines []string) RuneGrid {
	ys := make([][]rune, 0, 0)
	for _, line := range lines {
		xs := make([]rune, 0, 0)
		for _, b := range line {
			xs = append(xs, b)
		}
		ys = append(ys, xs)
	}
	return RuneGrid{yxGrid: ys}
}

const (
	GROUND = '.'
	TREE   = '|'
	LUMBER = '#'
)

func update(grid RuneGrid) RuneGrid {
	grid2 := grid.Copy()
	ySize, xSize := grid.Size()
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			ns := grid.Neighbours8(x, y)
			// Create counter map
			counts := make(map[rune]int)
			for _, n := range ns {
				counts[grid.Get(n.X, n.Y)]++
			}
			switch tile := grid.Get(x, y); tile {
			case GROUND:
				if counts[TREE] >= 3 {
					grid2.Set(x, y, TREE)
				}
			case TREE:
				if counts[LUMBER] >= 3 {
					grid2.Set(x, y, LUMBER)
				}
			case LUMBER:
				if counts[LUMBER] >= 1 && counts[TREE] >= 1 {
					grid2.Set(x, y, LUMBER)
				} else {
					grid2.Set(x, y, GROUND)
				}
			}
		}
	}

	return grid2
}

func value(grid RuneGrid) int {
	ySize, xSize := grid.Size()
	counts := make(map[rune]int)
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			counts[grid.Get(x, y)]++
		}
	}
	return counts[TREE] * counts[LUMBER]
}

func part1() {
	lines := advent.InputLines("18.in")
	grid := parseRuneGrid(lines)
	grid.Render()
	minutes := 10
	for minutes > 0 {
		grid2 := update(grid)
		grid = grid2

		minutes--
		grid2.Render()
	}
	v := value(grid)
	fmt.Println("Part 1: The total resource value is", v)
}

/*
Since a pattern arises and repeats itself every 28 iterations we can caluculate what the grid
will look like on the nth iteration if n is large enough. I copied a couple thousand iteration
of (value, minutes) pairs into a spreadsheet and found the 28 iteration cycle. I then applied
the formula: "num_iterations + minutes * 28 = 1B" to each row in the sheet. One of them will
evenly divide into 1B and the answer will be at that iteration.
*/
func part2() {
	lines := advent.InputLines("18.in")
	grid := parseRuneGrid(lines)
	minutes := 1
	for {
		grid2 := update(grid)
		grid = grid2

		v := value(grid)
		fmt.Println(minutes, v)
		minutes++
	}
	v := value(grid)
	fmt.Println("Part 2: The total resource value is", v)
}

func main() {
	part1()
	// part2()
}
