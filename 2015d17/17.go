package main

import (
	"fmt"
	"github.com/jollyra/stringutil"
	// "github.com/jollyra/numutil"
	"os"
	"strconv"
	"strings"
)

func parseContainers(lines []string) []int {
	xs := make([]int, 0, len(lines))
	for _, line := range lines {
		i, _ := strconv.Atoi(strings.TrimSpace(line))
		xs = append(xs, i)
	}
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
		return pack(rest, max) + pack(rest, max-x)
	} else {
		return pack(rest, max) + pack(rest, max-x)
	}
}

func main() {
	filename := os.Args[1]
	knapsackSize, _ := strconv.Atoi(os.Args[2])

	lines := stringutil.InputLines(filename)
	containers := parseContainers(lines)

	ans := pack(containers, knapsackSize)
	fmt.Printf("There are %d ways to pack a full knapsack\n", ans)
}
