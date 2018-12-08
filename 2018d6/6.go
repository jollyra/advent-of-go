package main

import (
	"errors"
	"fmt"

	"github.com/jollyra/go-advent-util"
	. "github.com/jollyra/go-advent-util/point"
)

var minX = advent.BigInt
var minY = advent.BigInt
var maxX = 0
var maxY = 0

type coordlist map[Point]int

func newCoordList() coordlist { return make(map[Point]int) }

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

func onBorder(p Point) bool {
	if p.X == minX || p.X == maxX || p.Y == minY || p.Y == maxY {
		return true
	}
	return false
}

func parsePoints(lines []string) []Point {
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		points = append(points, Point{x, y})

		minX = advent.Min(minX, x)
		maxX = advent.Max(maxX, x)
		minY = advent.Min(minY, y)
		maxY = advent.Max(maxY, y)
	}
	return points
}

func buildGrid(coords coordlist, locations []Point) coordlist {
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := Point{x, y}
			// Check how far Point p is from each location
			dists := make([]int, 0)
			for _, loc := range locations {
				dists = append(dists, ManhattanDistance(loc, p))
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

func floodfillArea(cl coordlist, loc Point) (int, error) {
	horizon := make([]Point, 0)
	horizon = append(horizon, loc)
	visited := make([]Point, 0)
	area := 0
	for len(horizon) > 0 {
		var p Point
		p, horizon = horizon[0], horizon[1:]

		if Contains(visited, p) {
			continue
		}

		if cl[loc] == cl[p] {
			if onBorder(p) {
				return -1, errors.New("infinite area")
			}

			area++
			visited = append(visited, p)

			for _, q := range p.Neighbours4() {
				if !Contains(visited, q) && !Contains(horizon, q) {
					horizon = append(horizon, q)
				}
			}
		}
	}

	return area, nil
}

func wellConnectedAreaSize(locs []Point) int {
	max := 10000
	area := 0
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			dist := 0
			for _, q := range locs {
				dist += ManhattanDistance(Point{x, y}, q)
			}
			if dist < max {
				area++
			}
		}
	}
	return area
}

func findLargestFiniteArea(coords coordlist, locations []Point) int {
	areas := make([]int, 0)
	for _, loc := range locations {
		area, err := floodfillArea(coords, loc)
		if err == nil {
			areas = append(areas, area)
		}
	}
	return advent.MaxInts(areas...)
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	locations := parsePoints(lines)
	coords := newCoordList()
	coords = buildGrid(coords, locations)
	fmt.Println("Part 1:", findLargestFiniteArea(coords, locations))
	fmt.Println("Part 2:", wellConnectedAreaSize(locations))
}
