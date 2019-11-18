package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	// "time"
)

const (
	SPRING  = '+'
	SAND    = '.'
	CLAY    = '#'
	WETSAND = '|'
	WATER   = '~'
)

// TODO remove next 2 lines
var topLeft Point
var botRight Point

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
	for _, line := range lines {
		var x0, x1, y0, y1 int
		ny, errx := fmt.Sscanf(line, "x=%d, y=%d..%d", &x0, &y0, &y1)
		nx, erry := fmt.Sscanf(line, "y=%d, x=%d..%d", &y0, &x0, &x1)
		if errx != nil && erry != nil {
			panic(errx.Error())
		}
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
				fmt.Printf("%c", c)
			} else {
				fmt.Print(".")
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
	topLeft = Point{2<<32 - 1, 2<<32 - 1}
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
	// fmt.Println("falling", src)
	cur := src
	below := Point{cur.X, cur.Y + 1}
	for {
		groundType := getGroundType(coords, below)
		if groundType == SAND || groundType == WETSAND {
			coords[cur] = WETSAND
			cur = below
			below = Point{below.X, below.Y + 1}
			if cur.Y > minY {
				return coords, Point{}, errors.New("Min y value exceeded")
			}
		}
		if groundType == CLAY || groundType == WATER {
			coords[cur] = WETSAND
			return coords, cur, nil
		}
	}
	return coords, Point{2<<32 - 1, 2<<32 - 1}, errors.New("wtf")
}

func floodAcross(coords Coords, src Point, mover func(Point) Point) (Coords, error) {
	var (
		isWallLeft  = false
		isWallRight = false
		leftWall    Point
		rightWall   Point
		cur         Point
		err         error
	)

	// fmt.Println("floodAcross", src)
	// fmt.Println("searching left")
	cur = src
	for {
		cur = moveLeft(cur)
		t := getGroundType(coords, cur)
		if t == CLAY {
			isWallLeft = true
			leftWall = cur
			break
		}
		t = getGroundType(coords, moveBelow(cur))
		if t == WETSAND || t == SAND {
			// fmt.Println("flowing over")
			coords[cur] = WETSAND
			coords, err = waterDFS(cur, coords, topLeft, botRight)
			if err != nil {
				return coords, err
			}
			break
		}
		coords[cur] = WETSAND
	}
	// render(coords, topLeft, botRight)

	// fmt.Println("searching right")
	cur = src
	for {
		cur = moveRight(cur)
		t := getGroundType(coords, cur)
		if t == CLAY {
			isWallRight = true
			rightWall = cur
			break
		}
		t = getGroundType(coords, moveBelow(cur))
		if t == WETSAND || t == SAND {
			// fmt.Println("flowing over")
			coords[cur] = WETSAND
			coords, err = waterDFS(cur, coords, topLeft, botRight)
			if err != nil {
				return coords, err
			}
			break
		}
		coords[cur] = WETSAND
	}
	// render(coords, topLeft, botRight)

	// fmt.Println("flooding level")
	if isWallRight && isWallLeft {
		for x := leftWall.X + 1; x < rightWall.X; x++ {
			coords[Point{x, leftWall.Y}] = WATER
		}
		return coords, nil
	}
	// render(coords, topLeft, botRight)

	return coords, nil
}

func moveLeft(p Point) Point {
	return Point{p.X - 1, p.Y}
}

func moveRight(p Point) Point {
	return Point{p.X + 1, p.Y}
}

func moveBelow(p Point) Point {
	return Point{p.X, p.Y + 1}
}

func waterDFS(src Point, coords Coords, topLeft Point, botRight Point) (Coords, error) {
	var cur Point
	var err error

	coords, cur, err = floodDown(coords, src, botRight.Y)
	if err != nil {
		fmt.Println(err)
		return coords, err
	}

	coords, err = floodAcross(coords, cur, moveLeft)
	if err != nil {
		fmt.Println(err)
		return coords, err
	}

	return coords, nil
}

func countWetTiles(coords Coords, topLeft Point, botRight Point) int {
	n := 0
	for y := topLeft.Y; y <= botRight.Y; y++ {
		for x := topLeft.X; x <= botRight.X; x++ {
			p := Point{x, y}
			c, ok := coords[p]
			if ok {
				if c == WETSAND || c == WATER {
					n++
				}
			}
		}
	}
	return n
}

func main() {
	fmt.Println("Day 17 part 1")
	// lines := inputLines("17_test1.in")
	// lines := inputLines("17_mitch.in")
	lines := inputLines("17.in")
	coords := parseCoords(lines)
	topLeft, botRight = findBoundaries(coords)
	topLeft.X--
	botRight.X++
	fmt.Printf("topLeft: %v, botRight: %v\n", topLeft, botRight)

	var err error
	src := Point{500, 0}
	count := 180
	for count > 0 {
		fmt.Println(count)
		coords, err = waterDFS(src, coords, topLeft, botRight)
		if err != nil {
			break
		}
		// render(coords, topLeft, botRight)
		// time.Sleep(300 * time.Millisecond)
		count--
	}
	render(coords, topLeft, botRight)

	n := countWetTiles(coords, topLeft, botRight)
	// last wrong anser:27530
	// 27,537
	// 27,536
	// mitch's input: 37,735
	fmt.Println("part 1:", n)
}
