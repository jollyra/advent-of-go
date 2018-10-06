package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func indexOf(strings []string, target string) int {
	for i, s := range strings {
		if s == target {
			return i
		}
	}
	return -1
}

func uniqueNames(lines []string) []string {
	names := make([]string, 0)
	for _, line := range lines {
		words := strings.Split(line, " ")
		n0, n1 := string(words[0]), string(words[2])
		if indexOf(names, n0) == -1 {
			names = append(names, n0)
		}
		if indexOf(names, n1) == -1 {
			names = append(names, n1)
		}
	}
	return names
}

func parseRelationships(lines []string) map[string]int {
	relationships := make(map[string]int)
	for _, line := range lines {
		words := strings.Split(line, " ")
		rating, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalf("Failed to convert string %s to int: %v",
				words[1], err)
		}
		key := fmt.Sprintf("%s-%s", string(words[0]), string(words[2]))
		relationships[key] = rating
	}
	return relationships
}

func copyStringSlice(xs []string) []string {
	ys := make([]string, len(xs))
	for i := range xs {
		ys[i] = xs[i]
	}
	return ys
}

func permutations(dest *[][]string, xs []string, unfixed int) {
	if len(xs)-1 == unfixed {
		*dest = append(*dest, xs)
	}

	for i := unfixed; i < len(xs); i++ {
		ys := copyStringSlice(xs)
		ys[unfixed], ys[i] = ys[i], ys[unfixed]
		permutations(dest, ys, unfixed+1)
	}
}

func totalHappiness(relationships map[string]int, seatingArrangement []string) int {
	sum := 0
	for i, name := range seatingArrangement {
		iLeft := i - 1
		if i == 0 {
			iLeft = len(seatingArrangement) - 1
		}
		nameLeft := seatingArrangement[iLeft%len(seatingArrangement)]
		nameRight := seatingArrangement[(i+1)%len(seatingArrangement)]
		keyLeft := fmt.Sprintf("%s-%s", name, nameLeft)
		keyRight := fmt.Sprintf("%s-%s", name, nameRight)
		ratingLeft, ok := relationships[keyLeft]
		if ok {
			sum += ratingLeft
		}
		ratingRight, ok := relationships[keyRight]
		if ok {
			sum += ratingRight
		}
	}
	return sum
}

func max(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func main() {
	lines := inputLines("/Users/nrahkola/go/src/github.com/jollyra/" +
		"advent-of-go/knights_of_the_dinner_table/input.txt")
	names := uniqueNames(lines)
	fmt.Println(names)
	relationships := parseRelationships(lines)
	fmt.Println(relationships)
	ps := make([][]string, 0)
	permutations(&ps, names, 0)

	tables := make([]int, len(ps))
	for _, p := range ps {
		tables = append(tables, totalHappiness(relationships, p))
	}
	fmt.Println(max(tables...))
}
