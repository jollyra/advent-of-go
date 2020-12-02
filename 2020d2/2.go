package main

import (
	"fmt"

	"github.com/jollyra/go-advent-util"
)

func parse(s string) (a, b int, letter, password string) {
	n, err := fmt.Sscanf(s, "%d-%d %s : %s", &a, &b, &letter, &password)
	if err != nil {
		fmt.Println(n, err, a, b, letter, password)
		panic(err)
	}
	return
}

func validatePart1(a, b int, letter, password string) bool {
	var count int
	for _, c := range password {
		if string(c) == letter {
			count++
		}
	}
	if count >= a && count <= b {
		return true
	}
	return false
}

func validatePart2(a, b int, letter, password string) bool {
	if string(password[a-1]) == letter && string(password[b-1]) == letter {
		return false
	} else if string(password[a-1]) == letter || string(password[b-1]) == letter {
		return true
	}
	return false
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	var part1, part2 int
	for _, line := range lines {
		a, b, letter, password := parse(line)
		if validatePart1(a, b, letter, password) {
			part1++
		}
		if validatePart2(a, b, letter, password) {
			part2++
		}
	}
	fmt.Printf("part 1: %d, part 2: %d\n", part1, part2)
}
