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

type sue struct {
	name  string
	props map[string]int
}

func (s sue) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n", s.name)
	for k, v := range s.props {
		fmt.Fprintf(&b, " |  %s:%d\n", k, v)
	}
	fmt.Fprintf(&b, "\n")
	return b.String()
}

func newSue(name string) *sue {
	return &sue{
		name:  name,
		props: make(map[string]int),
	}
}

func parseLines(lines []string) []sue {
	sues := make([]sue, 0)
	for _, line := range lines {
		s := parseLine(line)
		sues = append(sues, *s)
	}
	return sues
}

func splitOnce(s, sep string) []string {
	return strings.SplitN(s, sep, 2)
}

func parseLine(line string) *sue {
	halves := splitOnce(line, ": ")
	name := halves[0]
	s := newSue(name)
	facts := strings.Split(halves[1], ", ")
	for _, fact := range facts {
		pair := strings.Split(fact, ": ")
		key, val := pair[0], pair[1]
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Panic("Failed to parse", fact, err)
		}
		s.props[key] = i
	}
	return s
}

func isMatch(full, partial sue) bool {
	for kP, vP := range partial.props {
		vF, ok := full.props[kP]
		if ok {
			if kP == "cats" || kP == "trees" {
				if vP <= vF {
					return false
				}
			} else if kP == "pomeranians" || kP == "goldfish" {
				if vP >= vF {
					return false
				}
			} else if vP != vF {
				return false
			}
		}
	}
	return true
}

func main() {
	lines := inputLines(os.Args[1])
	sues := parseLines(lines)

	lines = inputLines(os.Args[2])
	targetSue := parseLine(lines[0])
	fmt.Println("Target Sue:", targetSue)

	for _, s := range sues {
		match := isMatch(*targetSue, s)
		if match {
			fmt.Println("Match found!\n", s)
		}
	}
}
