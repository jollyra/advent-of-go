package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jollyra/go-advent-util"
)

type action struct {
	action string
	value  float64
}

func parse(lines []string) (actions []action) {
	for _, line := range lines {
		x, err := strconv.ParseFloat(line[1:], 64)
		if err != nil {
			panic(err)
		}
		actions = append(actions, action{action: line[:1], value: x})
	}
	return actions
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	actions := parse(lines)
	pos := 0 + 0i
	waypoint := 10 + 1i
	for _, a := range actions {
		switch action := a.action; action {
		case "N":
			waypoint += complex(0, a.value)
		case "S":
			waypoint += complex(0, -a.value)
		case "E":
			waypoint += complex(a.value, 0)
		case "W":
			waypoint += complex(-a.value, 0)
		case "L":
			x := int(a.value / 90)
			for i := 0; i < x; i++ {
				waypoint *= 1i
			}
		case "R":
			x := int(a.value / 90)
			for i := 0; i < x; i++ {
				waypoint *= -1i
			}
		case "F":
			pos = pos + complex(a.value, 0)*waypoint
		}
	}
	fmt.Println("ship's manhattan distance from (0, 0):", math.Abs(real(pos))+math.Abs(imag(pos)))
}
