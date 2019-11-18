package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
)

// Point is a 4-dimensional point in spacetime
type Point struct{ A, B, C, D int }

func parsePoint(s string) Point {
	var a, b, c, d int
	fmt.Sscanf(s, "%d,%d,%d,%d", &a, &b, &c, &d)
	return Point{a, b, c, d}
}

func parsePoints(lines []string) []Point {
	ps := make([]Point, 0, 0)
	for _, line := range lines {
		ps = append(ps, parsePoint(line))
	}
	return ps
}

func manhattanDistance(p, q Point) int {
	return abs(p.A-q.A) + abs(p.B-q.B) + abs(p.C-q.C) + abs(p.D-q.D)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printConstellations(cs [][]Point) {
	for _, c := range cs {
		for _, p := range c {
			fmt.Println(p)
		}
		fmt.Println()
	}
}

func isSameConstellation(c1, c2 []Point) bool {
	for _, p := range c1 {
		for _, q := range c2 {
			if manhattanDistance(p, q) <= 3 {
				return true
			}
		}
	}
	return false
}

func mergeConstellations(cs [][]Point, i int, j int) [][]Point {
	// Build csNew with all the elements of cs expect indices i and j.
	csNew := make([][]Point, 0, 0)
	for q, c := range cs {
		if q != i && q != j {
			csNew = append(csNew, c)
		}
	}

	// Merge the constellations at index i and j.
	cNew := make([]Point, 0, 0)
	for _, point := range cs[i] {
		cNew = append(cNew, point)
	}
	for _, point := range cs[j] {
		cNew = append(cNew, point)
	}

	// Add the new constellation to the master list
	return append(csNew, cNew)
}

func groupOne(constellations [][]Point) ([][]Point, bool) {
	for i := 0; i < len(constellations); i++ {
		for j := 0; j < len(constellations); j++ {
			if i == j {
				continue
			}

			if isSameConstellation(constellations[i], constellations[j]) {
				constellations = mergeConstellations(constellations, i, j)
				return constellations, true
			}
		}
	}
	return constellations, false
}

func groupMany(ps []Point) [][]Point {
	//Init constellations
	constellations := make([][]Point, 0, 0)
	for _, p := range ps {
		c := make([]Point, 0, 0)
		c = append(c, p)
		constellations = append(constellations, c)
	}

	//Combinations
	var didWeFindMore bool
	for {
		constellations, didWeFindMore = groupOne(constellations)
		if didWeFindMore == false {
			break
		}
	}

	return constellations
}

func countConstellations(fn string) {
	lines := advent.InputLines(fn)
	ps := parsePoints(lines)
	constellations := groupMany(ps)
	// printConstellations(constellations)
	fmt.Println(fn, len(constellations))
}

func main() {
	countConstellations("test0.in") // 2
	countConstellations("test1.in") // 4
	countConstellations("test2.in") // 3
	countConstellations("test3.in") // 8
	countConstellations("25.in")    // 331
}
