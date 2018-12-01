package main

import (
	"fmt"
	"os"

	"github.com/jollyra/go-advent-util"
)

func freq(freqs []int) int {
	freq := 0
	for _, i := range freqs {
		freq += i
	}
	return freq
}

func firstRepeat(freqs []int) int {
	seen := make(map[int]int)
	freq := 0
	for {
		for _, f := range freqs {
			seen[freq]++
			if seen[freq] == 2 {
				return freq
			}
			freq += f
		}
	}
}

func main() {
	lines := advent.InputLines(os.Args[1])
	freqs := advent.LinesToInts(lines)
	fmt.Println("Part 1: Ended on frequency", freq(freqs))
	fmt.Println("Part 2: First repeated freqency", firstRepeat(freqs))
}
