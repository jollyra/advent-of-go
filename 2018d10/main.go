package main

import (
	"fmt"
	// "github.com/jollyra/stringutil"
	// "github.com/jollyra/numutil"
	"github.com/jollyra/go-advent-util"
	"github.com/jollyra/go-advent-util/point"
	"math"
	// "time"
)

type Point = point.Point

var print = fmt.Println

func parse(lines []string) [][2]Point {
	stars := make([][2]Point, 0, len(lines))
	for _, line := range lines {
		p := Point{}
		v := Point{}
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>",
			&(p.X), &(p.Y), &(v.X), &(v.Y))
		print(line, p, v)
		stars = append(stars, [2]Point{p, v})
	}
	return stars
}

func show(sizeX, sizeY int, points []Point) {
	for y := 0; y < sizeX; y++ {
		for x := 0; x < sizeX; x++ {
			if point.Contains(points, Point{x, y}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		print()
	}
	print()
}

func showWindow(topLeft, bottomRight Point, points []Point) {
	for y := topLeft.Y; y < bottomRight.Y+10; y++ {
		for x := topLeft.X; x < bottomRight.X; x++ {
			if point.Contains(points, Point{x, y}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		print()
	}
	print()
}

func findBounds(stars []Point) (Point, Point) {
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxY := math.MinInt32
	maxX := math.MinInt32

	for _, star := range stars {
		if star.X < minX {
			minX = star.X
		}
		if star.X > maxX {
			maxX = star.X
		}
		if star.Y < minY {
			minY = star.Y
		}
		if star.Y > maxY {
			maxY = star.Y
		}
	}

	return Point{minX, minY}, Point{maxX, maxY}
}

func calcArea(topLeft, bottomRight Point) int {
	return (bottomRight.X - topLeft.X) * (bottomRight.Y - topLeft.Y)
}

func loop(stars [][2]Point) {
	i := 0
	// prevArea := math.MaxInt32
	for {
		ps := make([]Point, len(stars), len(stars))
		for i := range stars {
			ps = append(ps, stars[i][0])
		}

		topLeft, bottomRight := findBounds(ps)
		area := calcArea(topLeft, bottomRight)
		fmt.Printf("%10d %10d\n", i, area)

		// if i > 10383 && i < 10431 {
		if i > 10400 && i < 10431 {
			showWindow(topLeft, bottomRight, ps)
			// showWindow(Point{170, 0}, Point{300, 50}, ps)
			// time.Sleep(100 * time.Millisecond)
		}

		for i := range stars {
			pos, vel := stars[i][0], stars[i][1]
			stars[i] = [2]Point{Point{pos.X + vel.X, pos.Y + vel.Y}, vel}
		}

		i++
	}
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	stars := parse(lines)
	for _, s := range stars {
		print(s)
	}
	loop(stars)
}
