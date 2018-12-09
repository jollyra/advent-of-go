package main

import (
	"fmt"
	// "github.com/jollyra/stringutil"
	// "github.com/jollyra/numutil"
	// "github.com/jollyra/go-advent-util"
)

var print = fmt.Println

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func showGame(circle []int, cur, marble int) {
	fmt.Printf("[%d]  ", marble)
	for _, x := range circle {
		if x == cur {
			fmt.Printf("(%d) ", x)
		} else {
			fmt.Printf("%3d ", x)
		}
	}
	print()
}

// Insert inserts integer x after xs[index].
func Insert(xs []int, x, i int) []int {
	return append(xs[:i+1], append([]int{x}, xs[i+1:]...)...)
}

// Remove removes the integer after xs[index].
func Remove(xs []int, i int) []int {
	return append(xs[:i], xs[i+1:]...)
}

// MaxVal return the key and value of the item in the map with the max value.
func MaxVal(m map[int]int) (int, int) {
	maxVal := -1 << 31
	maxKey := -1
	for k, v := range m {
		if v > maxVal {
			maxVal = v
			maxKey = k
		}
	}
	return maxKey, maxVal
}

func play(players, lastMarble int) map[int]int {
	scores := make(map[int]int)
	circle := make([]int, 0)
	circle = append(circle, 0)
	cur := 0
	for marble := 1; marble <= lastMarble; marble++ {
		// if marble%100 == 0 {
		// print(marble)
		// }
		player := marble % players
		if marble%23 == 0 {
			scores[player] += marble
			removeIdx := (cur + len(circle) - 7) % len(circle)
			// print("remove marble", circle[removeIdx])
			scores[player] += circle[removeIdx]
			circle = Remove(circle, removeIdx)
			cur = removeIdx
		} else {
			insertIdx := (cur + 1) % len(circle)
			circle = Insert(circle, marble, insertIdx)
			cur = insertIdx + 1
		}
		// showGame(circle, cur, player)
	}
	// print(scores)
	return scores
}

func main() {
	_, v := MaxVal(play(9, 25))
	assert(v == 32)
	_, v = MaxVal(play(10, 1618))
	assert(v == 8317)
	_, v = MaxVal(play(13, 7999))
	assert(v == 146373)
	_, v = MaxVal(play(17, 1104))
	assert(v == 2764)
	_, v = MaxVal(play(21, 6111))
	assert(v == 54718)
	_, v = MaxVal(play(30, 5807))
	assert(v == 37305)
	_, v = MaxVal(play(432, 71019))
	assert(v == 400493)
	_, v = MaxVal(play(432, 7101900))
	print(v)
}
