package main

import (
	"fmt"
	"strings"

	"github.com/jollyra/go-advent-util"
)

func sameLetter(r1, r2 byte) bool {
	if r1 == r2+32 || r2 == r1+32 {
		return true
	}
	return false
}

func reduce(polymer string) string {
	var b strings.Builder
	i := 0
	for i < len(polymer) {
		if i < len(polymer)-1 {
			if sameLetter(polymer[i], polymer[i+1]) {
				i += 2
			} else {
				fmt.Fprintf(&b, "%c", polymer[i])
				i++
			}
		} else {
			fmt.Fprintf(&b, "%c", polymer[i])
			i++
		}
	}

	reduced := b.String()
	if len(reduced) == len(polymer) {
		return polymer
	}

	return reduce(reduced)
}

func improve(polymer string) string {
	shortest := polymer
	for r := 97; r <= 122; r++ {
		candidate := strings.Replace(polymer, string(r), "", -1)
		candidate = strings.Replace(candidate, string(r-32), "", -1)
		reducedCandidate := reduce(candidate)
		if len(reducedCandidate) < len(shortest) {
			shortest = reducedCandidate
		}
	}
	return shortest
}

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func main() {
	assert(sameLetter("a"[0], "A"[0]) == true)
	assert(sameLetter("Z"[0], "z"[0]) == true)
	assert(reduce("aA") == "")
	assert(reduce("abBA") == "")
	assert(reduce("abAB") == "abAB")
	assert(reduce("aabAAB") == "aabAAB")
	assert(reduce("dabAcCaCBAcCcaDA") == "dabCBAcaDA")
	assert(improve("dabAcCaCBAcCcaDA") == "daDA")
	fmt.Println("pass")

	polymer := advent.InputLines(advent.MustGetArg(1))[0]
	reduced := reduce(polymer)
	fmt.Println("Part 1:", len(reduced))

	improved := improve(polymer)
	fmt.Println("Part 2:", len(improved))
}
