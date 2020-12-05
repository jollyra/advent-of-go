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

type seat struct {
	row int
	col int
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	max := 0
	seats := make(map[seat]bool)
	for _, line := range lines {
		row, col := decode(line)
		seatID := row*8 + col
		if seatID > max {
			max = seatID
		}
		seats[seat{row: row, col: col}] = true
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			// fmt.Println(i*8 + j)
			exists := seats[seat{row: i, col: j}]
			if !exists {
				seatID := i*8 + j
				fmt.Println(i, j, seatID)
			}
		}
	}

	fmt.Println("Max seat ID: ", max)
}
