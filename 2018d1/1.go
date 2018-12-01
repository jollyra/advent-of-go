package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := advent.InputLines(os.Args[1])

	var sum int
	for _, line := range lines {
		i, _ := strconv.Atoi(strings.TrimSpace(line))
		sum += i
	}
	fmt.Println("part 1:", sum)

	seen := make(map[int]int)
	for {
		for _, line := range lines {
			seen[sum]++
			if seen[sum] == 2 {
				fmt.Println("part 2:", sum)
				return
			}
			i, _ := strconv.Atoi(strings.TrimSpace(line))
			sum += i
		}
	}
}
