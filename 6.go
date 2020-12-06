package main

import (
	"errors"
	"fmt"
	"github.com/jollyra/go-advent-util"
)

func parse(lines []string) (xss [][]string) {
	xs := make([]string, 0, 0)
	for _, line := range lines {
		if line == "" {
			xss = append(xss, xs)
			xs = make([]string, 0, 0)
			continue
		}
		xs = append(xs, line)
	}
	xss = append(xss, xs)
	return xss
}

type runeSet map[rune]bool

func newRuneSetFromString(s string) runeSet {
	set := make(runeSet)
	for _, r := range s {
		set[r] = true
	}
	return set
}

func newRuneSetFromStrings(ss []string) runeSet {
	set := make(runeSet)
	for _, s := range ss {
		for _, r := range s {
			set[r] = true
		}
	}
	return set
}

func intersect(s0, s1 runeSet) runeSet {
	sIntersection := make(runeSet)
	for k := range s0 {
		if s1[k] {
			sIntersection[k] = true
		}
	}
	return sIntersection
}

func intersects(ss ...runeSet) (runeSet, error) {
	if len(ss) < 2 {
		return nil, errors.New("Cannot take the intersect of less than 2 sets")
	}
	sIntersection := intersect(ss[0], ss[1])
	for _, sn := range ss[2:] {
		sIntersection = intersect(sIntersection, sn)
	}
	return sIntersection, nil
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	xss := parse(lines)
	fmt.Println(xss)

	acc := 0
	for _, xs := range xss {
		acc += len(newRuneSetFromStrings(xs))
	}
	fmt.Println("Part 1: ", acc)

	acc = 0
	for _, xs := range xss {
		if len(xs) == 1 {
			acc += len(xs[0])
			continue
		}
		var sets []runeSet
		for _, x := range xs {
			sets = append(sets, newRuneSetFromString(x))
		}
		is, err := intersects(sets...)
		if err != nil {
			panic(err)
		}
		acc += len(is)
	}
	fmt.Println("Part 2: ", acc)
}
