package main

import (
	"fmt"
	"github.com/jollyra/numutil"
	"github.com/jollyra/stringutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseContainers(lines []string) []int {
	xs := make([]int, 0, len(lines))
	for _, line := range lines {
		i, _ := strconv.Atoi(strings.TrimSpace(line))
		xs = append(xs, i)
	}
	sort.Ints(xs)
	numutil.Reverse(xs)
	return xs
}

func pack(xs []int, max int) int {
	if len(xs) == 0 {
		return 0
	}

	x := xs[0]
	rest := xs[1:]
	if x == max {
		return 1 + pack(rest, max)
	} else if x > max {
		return pack(rest, max)
	} else {
		return pack(rest, max) + pack(rest, max-x)
	}
}

func boundedPack(xs []int, max, bound int) int {
	if len(xs) == 0 {
		return 0
	}

	if bound == 0 {
		return 0
	}

	x := xs[0]
	rest := xs[1:]
	if x == max {
		return 1 + boundedPack(rest, max, bound)
	} else if x > max {
		return boundedPack(rest, max, bound)
	} else {
		return boundedPack(rest, max, bound) + boundedPack(rest, max-x, bound-1)
	}
}

func main() {
	filename := os.Args[1]
	knapsackSize, _ := strconv.Atoi(os.Args[2])

	lines := stringutil.InputLines(filename)
	containers := parseContainers(lines)
	fmt.Println(containers)

	ans := pack(containers, knapsackSize)
	fmt.Printf("There are %d ways to pack a full knapsack\n", ans)

	fmt.Printf("The number of ways to pack a min knapsack is %d\n",
		boundedPack(containers, knapsackSize, 4))
}
