package main

import (
	"fmt"
	"sort"

	"github.com/jollyra/go-advent-util"
)

func differences(xs []int) map[int]int {
	counter := make(map[int]int)
	sort.Ints(xs)
	for i := 1; i < len(xs); i++ {
		diff := xs[i] - xs[i-1]
		if diff > 3 || diff < 1 {
			panic("invalid diff")
		}
		counter[diff]++
	}
	return counter
}

func part1(xs []int) int {
	outletJoltage := 0
	deviceOutputJoltage := advent.MaxInts(xs...) + 3
	ys := make([]int, 0)
	ys = append(ys, outletJoltage)
	ys = append(ys, xs...)
	ys = append(ys, deviceOutputJoltage)
	counter := differences(ys)
	ans := counter[1] * counter[3]
	return ans
}

type edges map[int][]int

func parseEdges(xs []int) edges {
	E := make(map[int][]int)
	sort.Ints(xs)
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[j]-xs[i] > 0 && xs[j]-xs[i] <= 3 {
				E[xs[i]] = append(E[xs[i]], xs[j])
			}
		}
	}
	return E
}

type stack []int

func (s *stack) isEmpty() bool { return len(*s) == 0 }
func (s *stack) push(x int)    { *s = append(*s, x) }
func (s *stack) pop() int {
	n := len(*s) - 1
	x := (*s)[n]
	*s = (*s)[:n]
	return x
}

var memoizer map[int]int

func dfs(E edges, src, dst int) int {
	v, ok := memoizer[src]
	if ok {
		return v
	}

	if src == dst {
		return 1
	}

	acc := 0
	for _, v := range E[src] {
		acc += dfs(E, v, dst)
	}
	memoizer[src] = acc
	return acc
}

func part2(xs []int) int {
	outletJoltage := 0
	deviceOutputJoltage := advent.MaxInts(xs...) + 3
	ys := make([]int, 0)
	ys = append(ys, outletJoltage)
	ys = append(ys, xs...)
	ys = append(ys, deviceOutputJoltage)
	E := parseEdges(ys)
	sort.Ints(ys)
	memoizer = make(map[int]int)
	return dfs(E, outletJoltage, deviceOutputJoltage)
}

func main() {
	advent.Assert(35 == part1(advent.StringsToInts(advent.InputLines("10_test.in"))))
	advent.Assert(220 == part1(advent.StringsToInts(advent.InputLines("10_test2.in"))))
	advent.Assert(1885 == part1(advent.StringsToInts(advent.InputLines("10.in"))))

	advent.Assert(8 == part2(advent.StringsToInts(advent.InputLines("10_test.in"))))
	advent.Assert(19208 == part2(advent.StringsToInts(advent.InputLines("10_test2.in"))))
	advent.Assert(2024782584832 == part2(advent.StringsToInts(advent.InputLines("10.in"))))
	fmt.Println("pass")
}
