package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"strings"
)

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	V := make(map[string]bool)
	E := make(map[string][]string)
	for _, line := range lines {
		if strings.Contains(line, "no") {
			var a string
			var b string
			fmt.Sscanf(line, "%s %s contain no other bags.", &a, &b)
			v := fmt.Sprintf("%s %s", a, b)
			E[v] = []string{}
			V[v] = true
		} else {
			parts := strings.Split(line, "contain")
			var v0a string
			var v0b string
			fmt.Sscanf(parts[0], "%s %s bags", &v0a, &v0b)
			v0 := fmt.Sprintf("%s %s", v0a, v0b)
			E[v0] = make([]string, 0)
			V[v0] = true

			rhs := strings.Split(parts[1], ",")
			for _, term := range rhs {
				var c int
				var vna string
				var vnb string
				fmt.Sscanf(term, "%d %s %s bags", &c, &vna, &vnb)
				vn := fmt.Sprintf("%s %s", vna, vnb)
				E[v0] = append(E[v0], vn)
			}
		}
	}

	nodes := make(map[string]bool)
	for v := range V {
		// fmt.Println(v)
		horizon := []string{v}
		for len(horizon) > 0 {
			n := len(horizon) - 1
			cur := horizon[n]
			if cur == "shiny gold" {
				fmt.Println(v)
				nodes[v] = true
			}
			horizon = horizon[:n]
			horizon = append(horizon, E[cur]...)
		}
	}
	fmt.Println(len(nodes))
}
