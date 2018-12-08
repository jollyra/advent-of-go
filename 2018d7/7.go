package main

import (
	"container/heap"
	"fmt"

	"github.com/jollyra/go-advent-util"
)

type runeHeap []rune

func (h runeHeap) Len() int           { return len(h) }
func (h runeHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h runeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *runeHeap) Push(x interface{}) {
	if contains(*h, x.(rune)) {
		return
	}

	*h = append(*h, x.(rune))
}

func (h *runeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseEdges(lines []string) (map[rune][]rune, map[rune][]rune, []rune) {
	edges := make(map[rune][]rune)
	for key := range edges {
		edges[key] = make([]rune, 0)
	}

	dependency := make(map[rune][]rune)
	for key := range dependency {
		dependency[key] = make([]rune, 0)
	}

	srcs := make([]rune, 0)
	dsts := make([]rune, 0)
	for _, line := range lines {
		edge := advent.Split(line)
		a := rune(edge[0][0])
		b := rune(edge[1][0])
		edges[a] = append(edges[a], b)
		dependency[b] = append(dependency[b], a)
		srcs = append(srcs, a)
		dsts = append(dsts, b)
	}

	starts := diffStartingNodes(srcs, dsts)

	return edges, dependency, starts
}

func diffStartingNodes(xs, ys []rune) []rune {
	starts := make([]rune, 0)
	for _, x := range xs {
		if !contains(ys, x) {
			starts = append(starts, x)
		}
	}
	return starts
}

func contains(xs []rune, y rune) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func dependenciesMet(dependency, complete []rune) bool {
	if len(dependency) == 0 {
		return true
	}

	for _, d := range dependency {
		if !contains(complete, d) {
			return false
		}
	}
	return true
}

func bfs(edges, dependency map[rune][]rune, starts []rune) []rune {
	complete := make([]rune, 0, len(edges))
	horizon := &runeHeap{}
	for _, start := range starts {
		heap.Push(horizon, start)
	}
	heap.Init(horizon)
	for horizon.Len() > 0 {
		fmt.Println(horizon)

		var cur rune
		blocked := make([]rune, 0)
		found := false
		for found == false {
			cur = heap.Pop(horizon).(rune)
			if dependenciesMet(dependency[cur], complete) {
				for _, r := range blocked {
					heap.Push(horizon, r)
				}
				found = true
			} else {
				blocked = append(blocked, cur)
			}
		}

		complete = append(complete, cur)
		for _, next := range edges[cur] {
			heap.Push(horizon, next)
		}
	}

	return complete
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	edges, dependency, starts := parseEdges(lines)

	for k, v := range edges {
		fmt.Printf("edge %c ->", k)
		for i := range v {
			fmt.Printf("%c", v[i])
		}
		fmt.Println()
	}

	for k, v := range dependency {
		fmt.Printf("dependency %c ->", k)
		for i := range v {
			fmt.Printf("%c", v[i])
		}
		fmt.Println()
	}

	path := bfs(edges, dependency, starts)
	for i := range path {
		fmt.Printf("%c", path[i])
	}
	fmt.Println()
}
