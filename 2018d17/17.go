package main

import (
	"bufio"
	"errors"
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

func getGroundType(coords Coords, p Point) rune {
	groundType, ok := coords[p]
	if !ok {
		groundType = SAND // Sand isn't in the coordinate list
	}
	return groundType
}

func findBoundaries(coords Coords) (topLeft Point, botRight Point) {
	for p := range coords {
		if p.X < topLeft.X {
			topLeft.X = p.X
		} else if p.X > botRight.X {
			botRight.X = p.X
		}
		if p.Y < topLeft.Y {
			topLeft.Y = p.Y
		} else if p.Y > botRight.Y {
			botRight.Y = p.Y
		}
	}
	return topLeft, botRight
}

func floodDown(coords Coords, src Point, minY int) (Coords, Point, error) {
	cur := src
	var dst Point
	below := Point{cur.X, cur.Y + 1}
	for {
		// fmt.Println("cur: ", cur)
		fmt.Println("falling")
		groundType := getGroundType(coords, below)
		if groundType == SAND || groundType == WETSAND {
			coords[cur] = WETSAND
			cur = below
			below = Point{below.X, below.Y + 1}
			if below.Y > minY {
				return coords, Point{}, errors.New("Min y value exceeded")
			}
		}
		if groundType == CLAY || groundType == WATER {
			coords[cur] = WETSAND
			dst = cur
			break
		}
	}
	return coords, dst, nil
}

func floodAcross(coords Coords, src Point, mover func(Point) Point) (Coords, []Point) {
	cur := src
	next := mover(cur)
	for {
		fmt.Println("flooding across")
		groundType := getGroundType(coords, next)
		if groundType == SAND || groundType == WETSAND {
			if getGroundType(coords, Point{cur.X, cur.Y + 1}) == SAND {
				coords[cur] = SPRING
				return coords, []Point{cur}
			}
			coords[cur] = WATER
			cur = mover(cur)
			next = mover(cur)
		} else if groundType == CLAY {
			coords[cur] = WATER
			break
		}
	}
	return coords, []Point{}
}

func moveLeft(p Point) Point {
	return Point{p.X - 1, p.Y}
}

func moveRight(p Point) Point {
	return Point{p.X + 1, p.Y}
}

func waterDFS(src Point, coords Coords, topLeft Point, botRight Point) (Coords, []Point) {
	var cur Point
	var err error

	coords, cur, err = floodDown(coords, src, botRight.Y)
	if err != nil {
		return coords, []Point{}
	}
	render(coords, topLeft, botRight)

	var newSrcs []Point
	coords, newSrcs = floodAcross(coords, cur, moveLeft)
	if len(newSrcs) > 0 {
		return coords, newSrcs
	}
	render(coords, topLeft, botRight)

	coords, newSrcs = floodAcross(coords, cur, moveRight)
	if len(newSrcs) > 0 {
		return coords, newSrcs
	}
	render(coords, topLeft, botRight)

	return coords, []Point{src}
}

func main() {
	fmt.Println("Day 17 part 1")
	lines := inputLines("17_test.in")
	coords := parseCoords(lines)
	topLeft, botRight := findBoundaries(coords)
	fmt.Printf("topLeft: %v, botRight: %v\n", topLeft, botRight)

	// TODO remove next 2 lines
	topLeft = Point{X: 494, Y: 0}
	botRight = Point{X: 507, Y: 13}

	src := Point{500, 0}
	horizon := make([]Point, 0, 0)
	horizon = append(horizon, src)
	for len(horizon) > 0 {
		src, horizon = horizon[0], horizon[1:]
		var srcs []Point
		coords, srcs = waterDFS(src, coords, topLeft, botRight)
		for _, s := range srcs {
			horizon = append(horizon, s)
		}
		time.Sleep(300 * time.Millisecond)
	}
	render(coords, topLeft, botRight)
}
