package main

import (
	"container/list"
	"fmt"
	"github.com/jollyra/go-advent-util"
	util "github.com/jollyra/go-advent-util/point"
	"strings"
)

var print = fmt.Println

type point = util.Point

type stageType struct {
	Grid   [][]*actor
	Dx, Dy int
}

type actor struct {
	Type        rune
	Pos         point
	Health      int
	AttackPower int
}

// Update returns true if the actor made a move, false otherwise.
func (a *actor) Update(stage *stageType) bool {
	var target rune
	if a.Type == 'E' {
		target = 'G'
	} else {
		target = 'E'
	}

	path, err := readingOrderBFS(stage, a.Pos, target)
	if err != nil {
		return false
	}
	if len(path) >= 2 {
		nxt := path[len(path)-2]
		a.Pos = nxt
		return true
	}
	return false
}

func walk(ps map[point]point, q point) []point {
	path := make([]point, 0)
	if p, ok := ps[q]; ok {
		path = append(path, p)
		return append(path, walk(ps, p)...)
	}
	return path
}

// readingOrderBFS performs a reading ordered Breadth First Search to target
// and records the path taken. Returns the first step along the path.
func readingOrderBFS(stage *stageType, src point, target rune) ([]point, error) {
	prev := make(map[point]point)
	seen := make([]point, 0)
	seen = append(seen, src)
	queue := list.New()
	queue.PushBack(src)
	for queue.Len() > 0 {
		// print()
		// print("queue")
		// for e := queue.Front(); e != nil; e = e.Next() {
		// 	print(e.Value.(point))
		// }

		cur := queue.Remove(queue.Front()).(point)
		// print("cur", cur)
		for _, nxt := range cur.Neighbours4() {
			cell := stage.Grid[nxt.Y][nxt.X]
			if cell.Type == target {
				prev[nxt] = cur
				path := walk(prev, nxt)
				return path, nil
			} else if cell.Type == '.' {
				if !util.Contains(seen, nxt) {
					seen = append(seen, nxt)
					queue.PushBack(nxt)
					prev[nxt] = cur
				}
			}
		}
	}
	return []point{}, fmt.Errorf("No available moves")
}

func (stage *stageType) Show(xs []point, r rune) {
	var b strings.Builder
	fmt.Fprint(&b, "   ")
	for x := 0; x < stage.Dx; x++ {
		fmt.Fprintf(&b, "%d", x%10) // Show x values across the top
	}
	fmt.Fprint(&b, "\n")
	for y := range stage.Grid {
		fmt.Fprintf(&b, "%2d ", y) // Show y values down the left
		for _, x := range stage.Grid[y] {
			fmt.Fprintf(&b, "%c", x.Type)
		}
		fmt.Fprint(&b, "\n")
	}
	print(b.String())
}

func parseStage(lines []string) *stageType {
	grid := make([][]*actor, 0)
	var dx, dy int
	for y := range lines {
		row := make([]*actor, 0)
		for x := range lines[y] {
			pos := point{x, y}
			switch lines[y][x] {
			case 'E':
				e := &actor{'E', pos, 200, 3}
				row = append(row, e)
			case 'G':
				g := &actor{'G', pos, 200, 3}
				row = append(row, g)
			case '#':
				w := &actor{'#', pos, -1, -1}
				row = append(row, w)
			case '.':
				w := &actor{'.', pos, -1, -1}
				row = append(row, w)
			}

			if x > dx {
				dx = x
			}
			if y > dy {
				dy = y
			}
		}
		grid = append(grid, row)
	}
	return &stageType{grid, dx + 1, dy + 1}
}

func getActorsReadingOrder(stage *stageType) []*actor {
	actors := make([]*actor, 0)
	for y := range stage.Grid {
		for _, cell := range stage.Grid[y] {
			if cell.Type == 'E' || cell.Type == 'G' {
				actors = append(actors, cell)
			}
		}
	}
	return actors
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	stage := parseStage(lines)

	stage.Show([]point{}, '+')
	for i := 0; i < 4; i++ {
		for _, a := range getActorsReadingOrder(stage) {
			prev := a.Pos
			if a.Update(stage) == true {
				stage.Grid[prev.Y][prev.X] = &actor{'.', point{}, 0, 0}
				stage.Grid[a.Pos.Y][a.Pos.X] = a
			}
		}
		stage.Show([]point{}, '+')
	}
}
