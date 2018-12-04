package main

import (
	"errors"
	"fmt"
	"github.com/jollyra/go-advent-util"
)

type grid [1000][1000]int

type rectangle struct{ ID, Left, Top, Right, Bot int }

func (r *rectangle) String() string {
	return fmt.Sprintf("rect{ID %d, Left %d, Top %d, Right %d, Bot %d}", r.ID, r.Left, r.Top, r.Right, r.Bot)
}

func parseRectangle(line []int) *rectangle {
	return &rectangle{
		ID:    line[0],
		Left:  line[1],
		Top:   line[2],
		Right: line[3] + line[1] - 1,
		Bot:   line[4] + line[2] - 1,
	}
}

func parse(lines []string) []*rectangle {
	parsed := make([]*rectangle, 0, len(lines))
	for _, line := range lines {
		line = advent.StripNonDigits(line)
		words := advent.Split(line)
		ints := advent.StringsToInts(words)
		parsed = append(parsed, parseRectangle(ints))
	}
	return parsed
}

func drawRect(g *grid, rect *rectangle) {
	for x := rect.Left; x <= rect.Right; x++ {
		for y := rect.Top; y <= rect.Bot; y++ {
			g[y][x]++
		}
	}
}

func drawRects(g *grid, rects []*rectangle) {
	for _, r := range rects {
		drawRect(g, r)
	}
}

func countOverlap(g *grid) int {
	count := 0
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			if g[y][x] > 1 {
				count++
			}
		}
	}
	return count
}

func findNonOverlappingRectangle(g *grid, rs []*rectangle) (*rectangle, error) {
	for _, rect := range rs {
		overlap := false
		for x := rect.Left; x <= rect.Right; x++ {
			for y := rect.Top; y <= rect.Bot; y++ {
				if g[y][x] > 1 {
					overlap = true
				}
			}
		}

		if overlap == false {
			return rect, nil
		}
	}

	return &rectangle{}, errors.New("All rectangles overlap")
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	rectangles := parse(lines)
	g := &grid{}
	drawRects(g, rectangles)
	fmt.Println("Part 1:", countOverlap(g))

	rect, err := findNonOverlappingRectangle(g, rectangles)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part 2:", rect)
}
