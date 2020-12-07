package main

import (
	"fmt"
	"strings"

	"github.com/jollyra/go-advent-util"
)

func parseEdges(lines []string) map[string]map[string]int {
	E := make(map[string]map[string]int)
	for _, line := range lines {
		if strings.Contains(line, "no") {
			var a, b string
			fmt.Sscanf(line, "%s %s contain no other bags.", &a, &b)
			v := fmt.Sprintf("%s %s", a, b)
			E[v] = make(map[string]int)
		} else {
			parts := strings.Split(line, "contain")
			var v0a, v0b string
			fmt.Sscanf(parts[0], "%s %s bags", &v0a, &v0b)
			v0 := fmt.Sprintf("%s %s", v0a, v0b)
			E[v0] = make(map[string]int)
			rhs := strings.Split(parts[1], ",")
			for _, term := range rhs {
				var c int
				var vna, vnb string
				fmt.Sscanf(term, "%d %s %s bags", &c, &vna, &vnb)
				vn := fmt.Sprintf("%s %s", vna, vnb)
				E[v0][vn] = c
			}
		}
	}
	return E
}

type stack []string

func (s *stack) isEmpty() bool   { return len(*s) == 0 }
func (s *stack) push(str string) { *s = append(*s, str) }
func (s *stack) pop() string {
	n := len(*s) - 1
	str := (*s)[n]
	*s = (*s)[:n]
	return str
}

func partOneDFS(E map[string]map[string]int, target string) int {
	sources := make(map[string]bool)
	for v := range E {
		horizon := stack{v}
		for !horizon.isEmpty() {
			cur := horizon.pop()
			if cur == target {
				sources[v] = true
			}
			for v := range E[cur] {
				horizon.push(v)
			}
		}
	}
	return len(sources) - 1
}

func partTwoDFS(E map[string]map[string]int, source string) int {
	totalCost := 0
	horizon := stack{source}
	for !horizon.isEmpty() {
		cur := horizon.pop()
		for v, cost := range E[cur] {
			for i := 0; i < cost; i++ {
				horizon.push(v)
			}
			totalCost += cost
		}
	}
	return totalCost
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	E := parseEdges(lines)
	fmt.Println("part 1: ", partOneDFS(E, "shiny gold"))
	fmt.Println("part 2: ", partTwoDFS(E, "shiny gold"))
}
