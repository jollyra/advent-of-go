package main

import (
	"fmt"
	"github.com/jollyra/go-advent-util"
)

func decode(s string) (int, int) {
	var rows []int
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}

	var cols []int
	for i := 0; i < 8; i++ {
		cols = append(cols, i)
	}

	for _, r := range s {
		c := string(r)
		if c == "F" {
			rows = rows[:len(rows)/2]
		} else if c == "B" {
			rows = rows[len(rows)/2:]
		} else if c == "L" {
			cols = cols[:len(cols)/2]
		} else if c == "R" {
			cols = cols[len(cols)/2:]
		}
	}
	return rows[0], cols[0]
}

func seatID(row, col int) int {
	return row*8 + col
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	max := 0
	seats := make(map[int]bool)
	for _, line := range lines {
		row, col := decode(line)
		sid := seatID(row, col)
		seats[sid] = true
		if sid > max {
			max = sid
		}
	}
	fmt.Println("part 1: ", max)

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			sid := seatID(i, j)
			if !seats[sid] {
				if seats[sid+1] && seats[sid-1] {
					fmt.Println("part 2: ", sid)
				}
			}
		}
	}
}
