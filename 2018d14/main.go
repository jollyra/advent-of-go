package main

import (
	"fmt"
	// "github.com/jollyra/stringutil"
	"github.com/jollyra/numutil"
	// "github.com/jollyra/go-advent-util"
	"container/list"
	"strings"
)

var print = fmt.Println

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func digits(x int) []int {
	if x == 0 {
		return []int{0}
	}

	ds := make([]int, 0)
	for x > 0 {
		ds = append(ds, x%10)
		x = x / 10
	}
	numutil.Reverse(ds)
	return ds
}

func walk(l *list.List, e *list.Element, i int) *list.Element {
	for i := e.Value.(int) + 1; i > 0; i-- {
		if e.Next() == nil {
			e = l.Front()
		} else {
			e = e.Next()
		}
	}
	return e
}

func recipesAfter(n int) string {
	l := list.New()
	elf0 := l.PushBack(3)
	elf1 := l.PushBack(7)

	for l.Len() < n+10 {
		sum := elf0.Value.(int) + elf1.Value.(int)
		ds := digits(sum)
		for _, d := range ds {
			l.PushBack(d)
		}

		elf0 = walk(l, elf0, elf0.Value.(int))
		elf1 = walk(l, elf1, elf1.Value.(int))
	}

	var b strings.Builder
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i > 5405610 {
			fmt.Printf("%d", e.Value)
		}
		if i >= n && i < n+10 {
			fmt.Fprintf(&b, "%d", e.Value.(int))
		}
		i++
	}

	return b.String()
}

func part2(s string) int {
	l := list.New()
	elf0 := l.PushBack(3)
	elf1 := l.PushBack(7)

	window := list.New()
	i := 2
	for {
		sum := elf0.Value.(int) + elf1.Value.(int)
		ds := digits(sum)
		for _, d := range ds {
			i++
			l.PushBack(d)
			window.PushBack(d)

			for window.Len() > len(s) {
				window.Remove(window.Front())
			}

			var b strings.Builder
			for e := window.Front(); e != nil; e = e.Next() {
				fmt.Fprintf(&b, "%d", e.Value)
			}

			if b.String() == s {
				return i - len(s)
			}
		}
		elf0 = walk(l, elf0, elf0.Value.(int))
		elf1 = walk(l, elf1, elf1.Value.(int))
	}
}

func main() {
	assert(recipesAfter(9) == "5158916779")
	assert(recipesAfter(5) == "0124515891")
	assert(recipesAfter(18) == "9251071085")
	assert(recipesAfter(2018) == "5941429882")
	print("Part 1 pass")
	print("Part 1", recipesAfter(54056100))

	assert(part2("51589") == 9)
	assert(part2("01245") == 5)
	assert(part2("92510") == 18)
	assert(part2("59414") == 2018)
	print("Part 2 pass")
	print("Part 2", part2("540561"))
}
