package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	coords[Point{X: 500, Y: 0}] = '+' // Add the water source
	for _, line := range lines {
		var x0, x1, y0, y1 int
		ny, _ := fmt.Sscanf(line, "x=%d, y=%d..%d", &x0, &y0, &y1)
		nx, _ := fmt.Sscanf(line, "y=%d, x=%d..%d", &y0, &x0, &x1)
		if ny == 3 {
			for y := y0; y <= y1; y++ {
				p := Point{x0, y}
				coords[p] = '#'
			}
		} else if nx == 3 {
			for x := x0; x <= x1; x++ {
				p := Point{x, y0}
				coords[p] = '#'
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

func main() {
	fmt.Println("Day 17 part 1")
	lines := inputLines("17_test.in")
	coords := parseCoords(lines)
	topLeft := Point{X: 494, Y: 0}
	botRight := Point{X: 507, Y: 13}
	render(coords, topLeft, botRight)
}
