package main

import (
	"errors"
	"fmt"

	"github.com/jollyra/go-advent-util"
)

var big = 1<<63 - 1

var minX = big
var minY = big
var maxX = 0
var maxY = 0

type coordlist map[point]int

func newCoordList() coordlist { return make(map[point]int) }

type point struct{ X, Y int }

func (p point) String() string { return fmt.Sprintf("point{x=%d, y=%d}", p.X, p.Y) }

func (p point) Equals(q point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

// neighbours4 returns all neighbour of point (x, y) including diagonals.
func neighbours4(p point) []point {
	points := []point{
		point{p.X + 1, p.Y},
		point{p.X, p.Y + 1},
		point{p.X - 1, p.Y},
		point{p.X, p.Y - 1},
	}
	return points
}

func manhattanDistance(a, b point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Return the index of the smallest int in xs. Return error if there's a tie.
func minIntSlice(xs []int) (int, error) {
	minIndex := 0
	for i := range xs {
		if xs[i] < xs[minIndex] {
			minIndex = i
		}
	}

	isTie := false
	for i := range xs {
		if i != minIndex && xs[i] == xs[minIndex] {
			isTie = true
		}
	}

	if isTie {
		return 0, errors.New("There was a tie")
	}

	return minIndex, nil
}

func maxInts(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func contains(ps []point, q point) bool {
	for _, p := range ps {
		if p.Equals(q) {
			return true
		}
	}
	return false
}

func onBorder(p point) bool {
	if p.X == minX || p.X == maxX || p.Y == minY || p.Y == maxY {
		return true
	}
	return false
}

func parsePoints(lines []string) []point {
	points := make([]point, 0, len(lines))
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		points = append(points, point{x, y})

		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)
	}
	return points
}

func buildGrid(coords coordlist, locations []point) coordlist {
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := point{x, y}
			// Check how far point p is from each location
			dists := make([]int, 0)
			for _, loc := range locations {
				dists = append(dists, manhattanDistance(loc, p))
			}
			closestLocIndex, err := minIntSlice(dists)
			if err != nil {
				// Ties are no man's land
				coords[p] = -1
			} else {
				coords[p] = closestLocIndex
			}
		}
	}
	return coords
}

func floodfillArea(cl coordlist, loc point) (int, error) {
	horizon := make([]point, 0)
	horizon = append(horizon, loc)
	visited := make([]point, 0)
	area := 0
	for len(horizon) > 0 {
		var p point
		p, horizon = horizon[0], horizon[1:]

		if contains(visited, p) {
			continue
		}

		if cl[loc] == cl[p] {
			if onBorder(p) {
				return -1, errors.New("infinite area")
			}

			area++
			visited = append(visited, p)

			for _, q := range neighbours4(p) {
				if !contains(visited, q) && !contains(horizon, q) {
					horizon = append(horizon, q)
				}
			}
		}
	}

	return area, nil
}

func wellConnectedAreaSize(locs []point) int {
	max := 10000
	area := 0
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			dist := 0
			for _, q := range locs {
				dist += manhattanDistance(point{x, y}, q)
			}
			if dist < max {
				area++
			}
		}
	}
	return area
}

func findLargestFiniteArea(coords coordlist, locations []point) int {
	areas := make([]int, 0)
	for _, loc := range locations {
		area, err := floodfillArea(coords, loc)
		if err == nil {
			areas = append(areas, area)
		}
	}
	return maxInts(areas...)
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	locations := parsePoints(lines)
	coords := newCoordList()
	coords = buildGrid(coords, locations)
	fmt.Println("Part 1:", findLargestFiniteArea(coords, locations))
	fmt.Println("Part 2:", wellConnectedAreaSize(locations))
}
