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
	Elves   map[point]*elf
	Goblins map[point]*goblin
	Walls   map[point]*wall
	Dx, Dy  int
}

type entity interface {
	Rune() rune
}

type actor interface {
	Update()
}

type elf struct {
	Pos         point
	Health      int
	AttackPower int
}

type goblin struct {
	Pos         point
	Health      int
	AttackPower int
}

type wall struct{ r rune }

func (e *elf) Rune() rune    { return 'E' }
func (g *goblin) Rune() rune { return 'G' }
func (w *wall) Rune() rune   { return w.r }

// Update returns true if the actor made a move, false otherwise.
func (e *goblin) Update(stage *stageType) bool {
	path, err := readingOrderBFS(stage, e.Pos, 'E')
	if err != nil {
		return false
	}
	if len(path) >= 2 {
		nxt := path[len(path)-2]
		e.Pos = nxt
		return true
	}
	return false
}

// Update returns true if the actor made a move, false otherwise.
func (e *elf) Update(stage *stageType) bool {
	path, err := readingOrderBFS(stage, e.Pos, 'G')
	if err != nil {
		return false
	}
	if len(path) >= 2 {
		nxt := path[len(path)-2]
		e.Pos = nxt
		return true
	}
	return false
}

func (stage *stageType) GetEntity(p point) (entity, error) {
	if v, ok := stage.Elves[p]; ok {
		return v, nil
	}
	if v, ok := stage.Goblins[p]; ok {
		return v, nil
	}
	if v, ok := stage.Walls[p]; ok {
		return v, nil
	}
	return &elf{}, fmt.Errorf("No entity at %v", p)
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
			if entity, err := stage.GetEntity(nxt); err == nil {
				if entity.Rune() == target {
					prev[nxt] = cur
					path := walk(prev, nxt)
					return path, nil
				}
			} else {
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
	walls := make(map[point]*wall)
	elves := make(map[point]*elf)
	goblins := make(map[point]*goblin)
	var dx, dy int
	for y := range lines {
		for x := range lines[y] {
			pos := point{x, y}
			switch lines[y][x] {
			case 'E':
				elves[pos] = &elf{pos, 200, 3}
			case 'G':
				goblins[pos] = &goblin{pos, 200, 3}
			case '#':
				walls[pos] = &wall{'#'}
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

	stage.Show([]point{}, '+')
	for i := 0; i < 4; i++ {
		next := stage.ReadingOrderGenerator()
		for p, err := next(); err == nil; p, err = next() {
			if e, err := stage.GetEntity(p); err == nil {
				switch v := e.(type) {
				case *elf:
					if v.Update(stage) == true {
						delete(stage.Elves, p)
						stage.Elves[v.Pos] = v
					}
				case *goblin:
					if v.Update(stage) == true {
						delete(stage.Goblins, p)
						stage.Goblins[v.Pos] = v
					}
				}
			}
		}
		stage.Show([]point{}, '+')
	}
}
