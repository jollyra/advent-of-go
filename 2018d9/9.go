package main

import (
	"container/ring"
	"fmt"

	"github.com/jollyra/go-advent-util"
)

var print = fmt.Println

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func play(players, lastMarble int) map[int]int {
	scores := make(map[int]int)
	r := ring.New(1)
	r.Value = 0
	for marble := 1; marble <= lastMarble; marble++ {
		player := marble % players
		if marble%23 == 0 {
			scores[player] += marble
			for i := 0; i <= 7; i++ {
				r = r.Prev()
			}
			s := r.Unlink(1)
			scores[player] += s.Value.(int)
			r = r.Next()
		} else {
			r = r.Next()
			s := ring.New(1)
			s.Value = marble
			r.Link(s)
			r = r.Next()
		}
	}
	return scores
}

func main() {
	_, v := advent.MaxVal(play(9, 25))
	assert(v == 32)
	_, v = advent.MaxVal(play(10, 1618))
	assert(v == 8317)
	_, v = advent.MaxVal(play(13, 7999))
	assert(v == 146373)
	_, v = advent.MaxVal(play(17, 1104))
	assert(v == 2764)
	_, v = advent.MaxVal(play(21, 6111))
	assert(v == 54718)
	_, v = advent.MaxVal(play(30, 5807))
	assert(v == 37305)
	_, v = advent.MaxVal(play(432, 71019))
	assert(v == 400493)
	_, v = advent.MaxVal(play(432, 7101900))
	print(v)
}
