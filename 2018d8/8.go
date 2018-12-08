package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
)

type node struct {
	children []*node
	data     []int
}

func parseTree(xs []int, i int) (int, *node) {
	cs := xs[i]
	ms := xs[i+1]
	i += 2

	children := make([]*node, 0)
	for j := 0; j < cs; j++ {
		var child *node
		i, child = parseTree(xs, i)
		children = append(children, child)
	}
	self := &node{children: children}

	for k := i; k < i+ms; k++ {
		self.data = append(self.data, xs[k])
	}

	return i + ms, self
}

func valueV1(root *node) int {
	val := advent.SumInts(root.data...)
	for _, child := range root.children {
		val += valueV1(child)
	}
	return val
}

func valueV2(root *node) int {
	if len(root.children) == 0 {
		val := advent.SumInts(root.data...)
		return val
	}
	val := 0
	for _, i := range root.data {
		if i <= len(root.children) {
			val += valueV2(root.children[i-1])
		}
	}
	return val
}

func main() {
	line := advent.InputLines(advent.MustGetArg(1))[0]
	words := advent.Split(line)
	xs := advent.StringsToInts(words)
	_, root := parseTree(xs, 0)
	fmt.Println("Part 1:", valueV1(root))
	fmt.Println("Part 2:", valueV2(root))
}
