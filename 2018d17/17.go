package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func inputLines(fn string) []string {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}

func parseLines(lines []string) {
	for _, line := range lines {
		var x, y0, y1 int
		fmt.Sscanf(line, "x=%d, y=%d..%d", &x, &y0, &y1)
	}
}

func main() {
	fmt.Println("Day 17 part 1")
	lines := inputLines("17_test.in")
	fmt.Println(lines)
}
