package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
	util "github.com/jollyra/go-advent-util/point"
	"strings"
)

var print = fmt.Println

type point = util.Point

type actor struct {
	Type        rune
	Pos         point
	Health      int
	AttackPower int
}

// Update returns true if the actor made a move, false otherwise.
func (a *actor) Update(stage *stageType) bool {
	inRange := a.inRange(stage)
	stage.Show(inRange, '?')
	if len(inRange) == 0 {
		return false
	}

	return true
}

func (a *actor) inRange(stage *stageType) []point {
	inRange := make([]point, 0)
	next := stage.ReadingOrderGenerator()
	for p, err := next(); err == nil; p, err = next() {
		for _, q := range p.Neighbours4() {
			if a.Type == 'E' {
				if _, ok := stage.Goblins[q]; ok {
					inRange = append(inRange, p)
				}
			} else {
				if _, ok := stage.Elves[q]; ok {
					inRange = append(inRange, p)
				}
			}
		}
	}

	return inRange
}

type stageType struct {
	Elves   map[point]*actor
	Goblins map[point]*actor
	Walls   map[point]rune
	Dx, Dy  int
}

func (stage *stageType) ReadingOrderGenerator() func() (point, error) {
	var x0, y0 int
	return func() (point, error) {
		x := x0
		for y := y0; y < stage.Dy; y++ {
			for x < stage.Dx {
				nxt := point{x, y}
				if _, ok := stage.Walls[nxt]; !ok {
					x0 = x + 1
					y0 = y
					return nxt, nil
				}
				x++
			}
			x = 0
		}
		return point{}, fmt.Errorf("No more points")
	}
}

func (stage *stageType) Show(xs []point, r rune) {
	all := make(map[point]rune)
	for k := range stage.Elves {
		all[k] = 'E'
	}
	for k := range stage.Goblins {
		all[k] = 'G'
	}
	for k := range stage.Walls {
		all[k] = '#'
	}

	for _, p := range xs {
		all[p] = r
	}

	var b strings.Builder
	fmt.Fprint(&b, "   ")
	for x := 0; x < stage.Dx; x++ {
		fmt.Fprintf(&b, "%d", x%10) // Show x values across the top
	}
	fmt.Fprint(&b, "\n")
	for y := 0; y < stage.Dy; y++ {
		fmt.Fprintf(&b, "%2d ", y) // Show y values down the left
		for x := 0; x < stage.Dx; x++ {
			pos := point{x, y}
			if r, ok := all[pos]; ok {
				fmt.Fprintf(&b, "%c", r)
			} else {
				fmt.Fprint(&b, ".")
			}
		}
		fmt.Fprint(&b, "\n")
	}
	print(b.String())
}

func parseStage(lines []string) *stageType {
	walls := make(map[point]rune)
	elves := make(map[point]*actor)
	goblins := make(map[point]*actor)
	var dx, dy int
	for y := range lines {
		for x := range lines[y] {
			pos := point{x, y}
			switch lines[y][x] {
			case 'E':
				elves[pos] = &actor{'E', pos, 200, 3}
			case 'G':
				goblins[pos] = &actor{'G', pos, 200, 3}
			case '#':
				walls[pos] = '#'
			}

			if x > dx {
				dx = x
			}
			if y > dy {
				dy = y
			}
		}
	}
	return &stageType{elves, goblins, walls, dx + 1, dy + 1}
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	stage := parseStage(lines)
	for _, elf := range stage.Elves {
		elf.Update(stage)
		return
	}
}
