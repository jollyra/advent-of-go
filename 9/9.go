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

func makeKey(a, b string) string {
	return fmt.Sprintf("%s-%s", a, b)
}

func parseEdges(lines []string) map[string]int {
	edges := make(map[string]int)
	for _, line := range lines {
		words := strings.Split(line, " ")
		weight, err := strconv.Atoi(words[2])
		if err != nil {
			log.Fatalf("Failed to convert string %s to int: %v",
				words[1], err)
		}
		key := makeKey(string(words[0]), string(words[1]))
		edges[key] = weight

		keyReverse := makeKey(string(words[1]), string(words[0]))
		edges[keyReverse] = weight
	}
	return edges
}

func uniqueNames(lines []string) []string {
	names := make([]string, 0)
	for _, line := range lines {
		words := strings.Split(line, " ")
		n0, n1 := string(words[0]), string(words[1])
		names = append(names, n0, n1)
	}
	return stringutil.Unique(names)
}

func pathCost(edges map[string]int, path []string) int {
	total := 0
	for i := 0; i < len(path)-1; i++ {
		key := makeKey(path[i], path[i+1])
		cost, ok := edges[key]
		if !ok {
			log.Fatalf("Edge %s not found in map", key)
		}
		total += cost
	}
	return total
}

func main() {
	lines := inputLines("/Users/nrahkola/go/src/github.com/jollyra/" +
		"advent-of-go/9/9.in")
	edges := parseEdges(lines)
	names := uniqueNames(lines)
	possiblePaths := make([][]string, 0)
	stringutil.Permutations(&possiblePaths, names, 0)
	pathCosts := make([]int, 0)
	for _, path := range possiblePaths {
		cost := pathCost(edges, path)
		pathCosts = append(pathCosts, cost)
	}
	shortestRouteCost := numutil.Max(pathCosts...)
	fmt.Println("The distance of Santa's shortest route is", shortestRouteCost)
}
