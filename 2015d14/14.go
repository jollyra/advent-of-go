package main

import (
	"bufio"
	"fmt"
	"github.com/jollyra/numutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name      string
	speed     int
	endurance int
	rest      int
}

type result struct {
	name string
	dist int
}

func (r *reindeer) race(delta int, ch chan result) {
	dist := 0
	endurance := r.endurance
	rest := r.rest
	state := "flying"
	for s := 0; s < delta; s++ {
		if state == "flying" {
			dist += r.speed
			endurance--
			if endurance == 0 {
				state = "resting"
				endurance = r.endurance
			}
		} else if state == "resting" {
			rest--
			if rest == 0 {
				state = "flying"
				rest = r.rest
			}
		}
	}
	ch <- result{r.name, dist}
}

func inputLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func initReindeer(line string) reindeer {
	words := strings.Split(line, " ")
	speed, _ := strconv.Atoi(words[3])
	endurance, _ := strconv.Atoi(words[6])
	rest, _ := strconv.Atoi(words[13])
	return reindeer{
		name:      string(words[0]),
		speed:     speed,
		endurance: endurance,
		rest:      rest,
	}
}

func collectResults(ch chan result, size int) []result {
	results := make([]result, 0)
	for i := 0; i < size; i++ {
		results = append(results, <-ch)
	}
	return results
}

func max(rs []result) result {
	max := rs[0]
	for _, r := range rs {
		if r.dist > max.dist {
			max = r
		}
	}
	return max
}

func findAllWinners(rs []result, max int) []result {
	allWinners := make([]result, 0)
	for _, r := range rs {
		if r.dist == max {
			allWinners = append(allWinners, r)
		}
	}
	return allWinners
}

func updateScore(score *map[string]int, winners []result) {
	for _, w := range winners {
		(*score)[w.name]++
	}
}

func maxVal(m map[string]int) int {
	vals := make([]int, 0)
	for _, v := range m {
		vals = append(vals, v)
	}
	return numutil.Max(vals...)
}

func main() {
	lines := inputLines("/Users/nrahkola/go/src/github.com/jollyra/" +
		"advent-of-go/14/14.in")

	rs := make([]reindeer, 0)
	for _, line := range lines {
		rs = append(rs, initReindeer(line))
	}

	ch := make(chan result)
	raceLengthSeconds := 2503
	score := make(map[string]int)
	for i := 1; i <= raceLengthSeconds; i++ {
		for _, r := range rs {
			go func(r reindeer) {
				r.race(i, ch)
			}(r)
		}
		results := collectResults(ch, len(rs))
		winner := max(results)
		winners := findAllWinners(results, winner.dist)
		updateScore(&score, winners)
	}

	fmt.Println(maxVal(score))
}
