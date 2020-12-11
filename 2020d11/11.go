package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
	"github.com/jollyra/go-advent-util/point"
)

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, b := range row {
			fmt.Printf("%c", b)
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func parse(lines []string) (grid [][]byte) {
	for y := range lines {
		row := make([]byte, len(lines))
		for x := range lines[y] {
			row = append(row, lines[y][x])
		}
		grid = append(grid, row)
	}
	return grid
}

func inBounds(grid [][]byte, p point.Point) bool {
	valid := true
	dy := len(grid)
	dx := len(grid[0])
	if p.X < 0 || p.X >= dx {
		valid = false
	}
	if p.Y < 0 || p.Y >= dy {
		valid = false
	}
	return valid
}

func occupiedSeatsInLineOfSight(grid [][]byte, p point.Point) (occupied int) {
	slopes := []point.Point{
		point.Point{-1, -1},
		point.Point{0, -1},
		point.Point{1, -1},
		point.Point{-1, 0},
		point.Point{1, 0},
		point.Point{-1, 1},
		point.Point{0, 1},
		point.Point{1, 1},
	}
outer:
	for _, m := range slopes {
		q := point.Add(p, m)
		for inBounds(grid, q) {
			tile := grid[q.Y][q.X]
			if tile == '#' {
				occupied++
				continue outer
			}
			q = point.Add(q, m)
		}
	}
	return occupied
}

func update(grid [][]byte) (updatedGrid [][]byte, updated bool) {
	for y := range grid {
		updatedRow := make([]byte, len(grid[y]))
		copy(updatedRow, grid[y])
		updatedGrid = append(updatedGrid, updatedRow)

		for x := range grid[y] {
			p := point.Point{X: x, Y: y}
			occupied := occupiedSeatsInLineOfSight(grid, p)
			switch tile := grid[y][x]; tile {
			case 'L':
				if occupied == 0 {
					updatedGrid[y][x] = '#'
					updated = true
				}
			case '#':
				if occupied >= 5 {
					updatedGrid[y][x] = 'L'
					updated = true
				}
			case '.':
			}
		}
	}
	return updatedGrid, updated
}

func countTile(grid [][]byte, target byte) (count int) {
	for _, row := range grid {
		for _, tile := range row {
			if tile == target {
				count++
			}
		}
	}
	return count
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	grid := parse(lines)

	i := 1
	for {
		fmt.Println(i)
		var updated bool
		grid, updated = update(grid)
		if !updated {
			fmt.Println(i)
			fmt.Println("occupied seats:", countTile(grid, '#'))
			break
		}
		i++
	}
}
