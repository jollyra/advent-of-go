package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"strings"
)

func parseEdges(lines []string) map[string]map[string]int {
	E := make(map[string]map[string]int)
	for _, line := range lines {
		if strings.Contains(line, "no") {
			var a string
			var b string
			fmt.Sscanf(line, "%s %s contain no other bags.", &a, &b)
			v := fmt.Sprintf("%s %s", a, b)
			E[v] = make(map[string]int)
		} else {
			parts := strings.Split(line, "contain")
			var v0a string
			var v0b string
			fmt.Sscanf(parts[0], "%s %s bags", &v0a, &v0b)
			v0 := fmt.Sprintf("%s %s", v0a, v0b)
			E[v0] = make(map[string]int)

			rhs := strings.Split(parts[1], ",")
			for _, term := range rhs {
				var c int
				var vna string
				var vnb string
				fmt.Sscanf(term, "%d %s %s bags", &c, &vna, &vnb)
				vn := fmt.Sprintf("%s %s", vna, vnb)
				E[v0][vn] = c
			}
		}
	}
	return E
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	E := parseEdges(lines)

	nodes := make(map[string]bool)
	for v := range E {
		horizon := []string{v}
		for len(horizon) > 0 {
			n := len(horizon) - 1
			cur := horizon[n]
			if cur == "shiny gold" {
				nodes[v] = true
			}
			horizon = horizon[:n]
			for v := range E[cur] {
				horizon = append(horizon, v)
			}
		}
	}
	fmt.Println("part 1: ", len(nodes)-1)

	totalCost := 0
	horizon := []string{"shiny gold"}
	for len(horizon) > 0 {
		n := len(horizon) - 1
		cur := horizon[n]
		horizon = horizon[:n]
		for v, cost := range E[cur] {
			for i := 0; i < cost; i++ {
				horizon = append(horizon, v)
			}
			totalCost += cost
		}
	}
	fmt.Println("part 2: ", totalCost)
}
