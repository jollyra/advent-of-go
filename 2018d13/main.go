package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
	p "github.com/jollyra/go-advent-util/point"
	"time"
)

type point = p.Point

var print = fmt.Println

var down = point{0, 1}
var up = point{0, -1}
var left = point{-1, 0}
var right = point{1, 0}

type cart struct {
	pos   point
	dir   point
	turns int
}

func (c *cart) Show() rune {
	if c.dir.Equals(down) {
		return 'v'
	} else if c.dir.Equals(up) {
		return '^'
	} else if c.dir.Equals(left) {
		return '<'
	} else if c.dir.Equals(right) {
		return '>'
	}
	return '?'
}

func (c *cart) Update(grid map[point]rune) {
	c.pos = p.Add(c.pos, c.dir)
	if c.dir == down || c.dir == up {
		if grid[c.pos] == '\\' {
			c.dir = p.RotateL(c.dir)
		} else if grid[c.pos] == '/' {
			c.dir = p.RotateR(c.dir)
		}
	} else if c.dir == left || c.dir == right {
		if grid[c.pos] == '\\' {
			c.dir = p.RotateR(c.dir)
		} else if grid[c.pos] == '/' {
			c.dir = p.RotateL(c.dir)
		}
	}

	if grid[c.pos] == '+' {
		switch i := c.turns % 3; i {
		case 0:
			c.dir = p.RotateL(c.dir)
		case 1:
		case 2:
			c.dir = p.RotateR(c.dir)
		}
		c.turns++
	}
}

func parseTracks(lines []string) (map[point]rune, []*cart, int, int) {
	var dx, dy int
	carts := make([]*cart, 0)
	cl := make(map[point]rune)
	for y, line := range lines {
		for x, r := range line {
			pos := point{x, y}
			if r == 'v' {
				carts = append(carts, &cart{pos, down, 0})
				cl[pos] = '|'
			} else if r == '^' {
				carts = append(carts, &cart{pos, up, 0})
				cl[pos] = '|'
			} else if r == '<' {
				carts = append(carts, &cart{pos, left, 0})
				cl[pos] = '-'
			} else if r == '>' {
				carts = append(carts, &cart{pos, right, 0})
				cl[pos] = '-'
			} else {
				cl[pos] = r
			}

			if y > dy {
				dy = y
			}

			if x > dx {
				dx = x
			}
		}
	}
	return cl, carts, dx + 1, dy + 1
}

func findCarts(cl map[point]rune, carts []*cart, dx, dy int) []*cart {
	sorted := make([]*cart, 0)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			for _, cart := range carts {
				if cart.pos.Equals(point{x, y}) {
					sorted = append(sorted, cart)
				}
			}
		}
	}
	return sorted
}

func findCart(carts []*cart, q point) (*cart, error) {
	for _, cart := range carts {
		if cart.pos.Equals(q) {
			return cart, nil
		}
	}
	return &cart{}, fmt.Errorf("No cart at position %q", q)
}

func showTracks(cl map[point]rune, carts []*cart, dx, dy int) {
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pos := point{x, y}
			if cart, err := findCart(carts, pos); err == nil {
				fmt.Printf("%c", cart.Show())
			} else {
				r, ok := cl[pos]
				if ok {
					fmt.Printf("%c", r)
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func collision(carts []*cart) (point, error) {
	for i := range carts {
		for j := range carts {
			if i != j {
				if carts[i].pos.Equals(carts[j].pos) {
					return carts[i].pos, nil
				}
			}
		}
	}
	return point{}, fmt.Errorf("No collisions")
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	grid, carts, dx, dy := parseTracks(lines)
	for {
		sortedCarts := findCarts(grid, carts, dx, dy)
		for _, cart := range sortedCarts {
			showTracks(grid, carts, dx, dy)
			cart.Update(grid)
			if pos, err := collision(carts); err == nil {
				print("collision at", pos)
				showTracks(grid, carts, dx, dy)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
