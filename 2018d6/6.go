package main

import (
	"fmt"
	// "github.com/jollyra/stringutil"
	// "github.com/jollyra/numutil"
	"errors"
	"github.com/jollyra/go-advent-util"
)

var size = 500

type point struct{ X, Y int }

func (p point) String() string { return fmt.Sprintf("point{x=%d, y=%d}", p.X, p.Y) }

type coordlist map[point]int

func show(cl coordlist) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%2d", cl[point{x, y}])
		}
		fmt.Println("")
	}
}

func parsePoints(lines []string) []point {
	points := make([]point, 0, len(lines))
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		points = append(points, point{x, y})
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(a, b point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

// min returns the index of the smallest int in xs.
func min(xs []int) (int, error) {
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

func walk(coords coordlist, locations []point) coordlist {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			p := point{x, y}
			// check how far point is from each location
			dists := make([]int, 0)
			for _, loc := range locations {
				dists = append(dists, manhattanDistance(loc, p))
			}
			closestLocIndex, err := min(dists)
			if err != nil {
				// there was a tie
				coords[p] = -1
			} else {
				coords[p] = closestLocIndex
			}
		}
	}
	return coords
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

func (p point) Equals(q point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
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
	if p.X == 0 || p.X == size-1 {
		return true
	}
	if p.Y == 0 || p.Y == size-1 {
		return true
	}
	return false
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

func part2(locs []point) int {
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

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	locations := parsePoints(lines)
	fmt.Println(locations)
	coords := make(map[point]int)
	coords = walk(coords, locations)
	// show(coords)
	for _, loc := range locations {
		area, _ := floodfillArea(coords, loc)
		fmt.Println(loc, area)
	}

	fmt.Println("part 2:", part2(locations))
}
