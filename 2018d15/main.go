package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
	p "github.com/jollyra/go-advent-util/point"
	"strings"
)

var print = fmt.Println

type point = p.Point

type actor struct {
	Pos         point
	Health      int
	AttackPower int
}

type stageType struct {
	Elves   map[point]*actor
	Goblins map[point]*actor
	Walls   map[point]rune
	Dx, Dy  int
}

func (stage *stageType) String() string {
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
			if _, ok := stage.Elves[pos]; ok {
				fmt.Fprint(&b, "E")
			} else if _, ok := stage.Goblins[pos]; ok {
				fmt.Fprint(&b, "G")
			} else if _, ok := stage.Walls[pos]; ok {
				fmt.Fprint(&b, "#")
			} else {
				fmt.Fprint(&b, ".")
			}
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
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
				elves[pos] = &actor{pos, 200, 3}
			case 'G':
				goblins[pos] = &actor{pos, 200, 3}
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
	print(stage)
}
