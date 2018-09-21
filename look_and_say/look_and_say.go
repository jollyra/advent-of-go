package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func inputLine() string {
	file, err := os.Open("/Users/nrahkola/go/src/github.com/jollyra/advent-of-go/not_quite_lisp/1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func WhatFloor(instructions string) int {
	sum := 0
	for _, ins := range instructions {
		if ins == '(' {
			sum += 1
		} else if ins == ')' {
			sum -= 1
		}
	}
	return sum
}

func DistanceToFloor(instructions string, floor int) int {
	sum := 0
	for i, ins := range instructions {
		if ins == '(' {
			sum += 1
		} else if ins == ')' {
			sum -= 1
		}

		if sum == floor {
			return i + 1
		}
	}
	return -1
}

func main() {
	line := inputLine()
	fmt.Println("Part 1:", WhatFloor(line))
	fmt.Println("Part 2:", DistanceToFloor(line, -1))
}
