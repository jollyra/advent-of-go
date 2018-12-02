package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
	"log"
)

type runeCounter map[rune]int

func (counter runeCounter) Contains(targetVal int) bool {
	for _, val := range counter {
		if val == targetVal {
			return true
		}
	}
	return false
}

func counterizeStrings(ss []string) []runeCounter {
	counters := make([]runeCounter, 0)
	for _, s := range ss {
		counter := make(map[rune]int)
		for _, r := range s {
			counter[r]++
		}
		counters = append(counters, counter)
	}
	return counters
}

func score(counters []runeCounter) int {
	var double, triple int
	for _, counter := range counters {
		if counter.Contains(2) {
			double++
		}
		if counter.Contains(3) {
			triple++
		}
	}
	return double * triple
}

func hammingDist(s1, s2 string) int {
	if len(s1) != len(s2) {
		log.Fatalf("Cannot calculate Hamming distance of unequal length strings: %s, %s", s1, s2)
	}
	dist := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			dist++
		}
	}
	return dist
}

func part2(ss []string) {
	for _, s1 := range ss {
		for _, s2 := range ss {
			if hammingDist(s1, s2) == 1 {
				fmt.Println(s1)
				fmt.Println(s2)
			}
		}
	}
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	counters := counterizeStrings(lines)
	checksum := score(counters)
	fmt.Println(checksum)
	part2(lines)
}
