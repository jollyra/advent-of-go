package main

import (
	"fmt"
	// "github.com/jollyra/stringutil"
	// "github.com/jollyra/numutil"
	"github.com/jollyra/go-advent-util"
	"strings"
)

var print = fmt.Println

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func parse(lines []string) (string, map[string]string) {
	var initialState string
	fmt.Sscanf(lines[0], "initial state: %s", &initialState)

	transform := make(map[string]string)
	for _, line := range lines[2:] {
		var in, out string
		fmt.Sscanf(line, "%s => %s", &in, &out)
		transform[in] = out
	}
	return initialState, transform
}

func sumPotsWithPlants(s string, offset int) int {
	acc := 0
	for i := range s {
		if s[i] == '#' {
			acc += i - offset
		}
	}
	return acc
}

func nextGen(transform map[string]string, offset int, gen string) (int, string) {
	if gen[0] == '#' {
		gen = ".." + gen
		offset += 2
	} else if gen[:1] == ".#" {
		gen = "." + gen
		offset++
	} else if gen[len(gen)-1] == '#' {
		gen = gen + ".."
	} else if gen[len(gen)-2:] == "#." {
		gen = gen + "."
	}

	var b strings.Builder
	for i := 0; i < len(gen); i++ {
		plant := "."
		var pre string
		if i == 0 {
			pre = fmt.Sprintf("..%s", gen[:3])
		} else if i == 1 {
			pre = fmt.Sprintf(".%s", gen[:4])
		} else if i == len(gen)-1 {
			pre = fmt.Sprintf("%s..", gen[i-2:])
		} else if i == len(gen)-2 {
			pre = fmt.Sprintf("%s.", gen[i-2:])
		} else {
			pre = gen[i-2 : i+3]
		}

		if v, ok := transform[pre]; ok {
			plant = v
		}
		fmt.Fprintf(&b, "%s", plant)
	}

	return offset, b.String()
}

func main() {
	assert(sumPotsWithPlants("#....##....#####...#######....#.#..##.", 2) == 325)

	lines := advent.InputLines(advent.MustGetArg(1))
	gen, transform := parse(lines)
	for k, v := range transform {
		print(k, v)
	}

	offset := 0
	for i := 0; i < 50000000000; i++ {
		if i%10000 == 0 {
			print(i)
		}
		offset, gen = nextGen(transform, offset, gen)
		// fmt.Printf("%2d: %s\n", i, gen)
	}

	print("Part 1", sumPotsWithPlants(gen, offset))
}
