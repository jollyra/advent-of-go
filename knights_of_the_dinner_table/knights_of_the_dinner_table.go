package main

import (
	"bufio"
	"fmt"
	"github.com/jollyra/numutil"
	"github.com/jollyra/stringutil"
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

func uniqueNames(lines []string) []string {
	names := make([]string, 0)
	for _, line := range lines {
		words := strings.Split(line, " ")
		n0, n1 := string(words[0]), string(words[2])
		names = append(names, n0, n1)
	}
	return stringutil.Unique(names)
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

func main() {
	lines := inputLines("/Users/nrahkola/go/src/github.com/jollyra/" +
		"advent-of-go/knights_of_the_dinner_table/input.txt")
	names := uniqueNames(lines)
	relationships := parseRelationships(lines)
	ps := make([][]string, 0)
	stringutil.Permutations(&ps, names, 0)

	tables := make([]int, len(ps))
	for _, p := range ps {
		tables = append(tables, totalHappiness(relationships, p))
	}
	fmt.Println("The max happiness seating arrangement it", numutil.Max(tables...))
}
