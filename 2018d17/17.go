package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SPRING  = '+'
	SAND    = '.'
	CLAY    = '#'
	WETSAND = '|'
	WATER   = '~'
)

type Point struct{ X, Y int }
type Coords map[Point]rune

func inputLines(fn string) []string {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func parseCoords(lines []string) Coords {
	coords := make(map[Point]rune)
	coords[Point{X: 500, Y: 0}] = SPRING // Add the water source
	for _, line := range lines {
		var x0, x1, y0, y1 int
		ny, _ := fmt.Sscanf(line, "x=%d, y=%d..%d", &x0, &y0, &y1)
		nx, _ := fmt.Sscanf(line, "y=%d, x=%d..%d", &y0, &x0, &x1)
		if ny == 3 {
			for y := y0; y <= y1; y++ {
				p := Point{x0, y}
				coords[p] = CLAY
			}
		} else if nx == 3 {
			for x := x0; x <= x1; x++ {
				p := Point{x, y0}
				coords[p] = CLAY
			}
		} else {
			panic("Sscanf failed")
		}
	}
	return coords
}

func render(coords Coords, topLeft, botRight Point) {
	for y := topLeft.Y; y <= botRight.Y; y++ {
		for x := topLeft.X; x <= botRight.X; x++ {
			p := Point{x, y}
			c, ok := coords[p]
			if ok {
				fmt.Printf("%+q", c)
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Print("\n")
	}
}

func appendIfFalling(horizon []Point, p Point, coords Coords) []Point {
	if coords[p] == CLAY || coords[p] == WATER {
		return horizon
	}
	return append(horizon, p)
}

func getGroundType(coords Coords, p Point) rune {
	groundType, ok := coords[p]
	if !ok {
		groundType = SAND // Sand isn't in the coordinate list
	}
	return groundType
}

func waterDFS(src Point, coords Coords) Coords {
	// TODO remove next 2 lines
	topLeft := Point{X: 494, Y: 0}
	botRight := Point{X: 507, Y: 13}

	horizon := make([]Point, 0, 0)
	horizon = append(horizon, src)
	for len(horizon) > 0 {
		var cur Point
		cur, horizon = horizon[0], horizon[1:]
		groundType := getGroundType(coords, cur)
		pointBelow := Point{cur.X, cur.Y + 1}
		if groundType == SAND || groundType == WETSAND {
			if coords[pointBelow] == CLAY || coords[pointBelow] == WATER {
				horizon = append(horizon, Point{cur.X + 1, cur.Y})
				horizon = append(horizon, Point{cur.X - 1, cur.Y})
				coords[cur] = WATER
			} else {
				horizon = append(horizon, pointBelow)
				coords[cur] = WETSAND
			}
		} else if groundType == SPRING {
			horizon = appendIfFalling(horizon, pointBelow, coords)
		}

		render(coords, topLeft, botRight)
		time.Sleep(100 * time.Millisecond)
	}
	return coords
}

func main() {
	fmt.Println("Day 17 part 1")
	lines := inputLines("17_test.in")
	coords := parseCoords(lines)
	topLeft := Point{X: 494, Y: 0}
	botRight := Point{X: 507, Y: 13}
	render(coords, topLeft, botRight)
	for {
		coords = waterDFS(Point{500, 0}, coords)
	}
}
