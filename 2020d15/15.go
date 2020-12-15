package main

import (
	"fmt"
)

func memoryGame(startingNumbers []int, rounds int) int {
	memory := make(map[int]int)
	var xs []int

	var i int
	for i = 0; i < len(startingNumbers)-1; i++ {
		xs = append(xs, startingNumbers[i])
		memory[startingNumbers[i]] = i + 1
	}
	xs = append(xs, startingNumbers[i])

	for i := len(startingNumbers) - 1; i < rounds-1; i++ {
		prev := xs[len(xs)-1]
		prevIdx, ok := memory[prev]
		if ok {
			next := i + 1 - prevIdx
			xs = append(xs, next)
			memory[prev] = i + 1
		} else {
			xs = append(xs, 0)
			memory[prev] = i + 1
		}
	}
	return xs[len(xs)-1]
}

func main() {
	fmt.Println(memoryGame([]int{0, 3, 6}, 2020))
	fmt.Println(memoryGame([]int{6, 4, 12, 1, 20, 0, 16}, 2020))
	fmt.Println(memoryGame([]int{6, 4, 12, 1, 20, 0, 16}, 30000000))
}
