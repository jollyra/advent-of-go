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

func getGroundType(coords Coords, p Point) rune {
	groundType, ok := coords[p]
	if !ok {
		groundType = SAND // Sand isn't in the coordinate list
	}
	return groundType
}

func waterDFS(src Point, coords Coords) (Coords, Point) {
	var (
		falling       = true
		floodingLeft  = false
		floodingRight = false
	)

	// TODO remove next 2 lines
	topLeft := Point{X: 494, Y: 0}
	botRight := Point{X: 507, Y: 13}

	cur := src
	below := Point{cur.X, cur.Y + 1}
	for falling {
		groundType := getGroundType(coords, below)
		if groundType == SAND || groundType == WETSAND {
			coords[cur] = WETSAND
			cur = below
			below = Point{below.X, below.Y + 1}
		}
		if groundType == CLAY || groundType == WATER {
			coords[cur] = WETSAND
			falling = false
			floodingLeft = true
		}
	}
	render(coords, topLeft, botRight)

	left := Point{cur.X - 1, cur.Y}
	right := Point{cur.X + 1, cur.Y}

	for floodingLeft {
		groundType := getGroundType(coords, left)
		if groundType == SAND || groundType == WETSAND {
			if getGroundType(coords, Point{cur.X, cur.Y + 1}) == SAND {
				coords[cur] = SPRING
				falling = true
				floodingLeft = false
				return coords, cur
			} else {
				coords[cur] = WATER
				cur = left
				left = Point{left.X - 1, left.Y}
			}
		} else if groundType == CLAY {
			coords[cur] = WATER
			floodingLeft = false
			floodingRight = true
		}
	}
	render(coords, topLeft, botRight)

	for floodingRight {
		groundType := getGroundType(coords, right)
		if groundType == SAND || groundType == WETSAND {
			if getGroundType(coords, Point{cur.X, cur.Y + 1}) == SAND {
				coords[cur] = SPRING
				falling = true
				floodingRight = false
				return coords, cur
			} else {
				coords[cur] = WATER
				cur = right
				right = Point{right.X + 1, right.Y}
			}
		} else if groundType == CLAY {
			coords[cur] = WATER
			floodingRight = false
		}
	}
	render(coords, topLeft, botRight)

	return coords, src
}

func main() {
	fmt.Println("Day 17 part 1")
	lines := inputLines("17_test.in")
	coords := parseCoords(lines)
	topLeft := Point{X: 494, Y: 0}
	botRight := Point{X: 507, Y: 13}
	render(coords, topLeft, botRight)
	src := Point{499, 0}
	for {
		coords, src = waterDFS(src, coords)
		time.Sleep(500 * time.Millisecond)
	}
}
