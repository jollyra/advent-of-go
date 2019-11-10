package main

import (
	"container/ring"
	"fmt"
	// "github.com/jollyra/go-advent-util"
)

var print = fmt.Println

func play(n int) int {
	elves := make([]int, n, n)
	for i := range elves {
		elves[i] = 1
	}

	eliminated := n
	i := 0
	for eliminated > 1 {
		j := i + 1
		for elves[j%n] == 0 {
			j++
		}
		elves[i%n] += elves[j%n]
		elves[j%n] = 0
		i = j
		for elves[i%n] == 0 {
			i++
		}
		eliminated--
	}

	winner := -1
	for i := range elves {
		if elves[i] > 0 {
			print(elves[i], i)
			winner = i
		}
	}

	return winner
}

func playV2(n int) int {
	r := ring.New(n)
	for r.Len() > 1 {
		r.Unlink(1)
		r = r.Next()
		print(r.Len())
	}

	return r.Len()
}

func main() {
	// fmt.Println(play(5) + 1)
	// fmt.Println(play(3018458) + 1)
	fmt.Println(playV2(5))
	// fmt.Println(playV2(3018458))
}
