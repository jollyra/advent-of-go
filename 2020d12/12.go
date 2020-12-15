package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jollyra/go-advent-util"
	// "github.com/stretchr/testify/assert"
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
		a := action{
			action: line[:1],
			value:  x,
		}
		actions = append(actions, a)
	}
	return actions
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	actions := parse(lines)
	fmt.Println(actions)

	pos := 0 + 0i
	heading := 1 + 0i
	for _, a := range actions {
		switch action := a.action; action {
		case "N":
			pos += complex(0, a.value)
		case "S":
			pos += complex(0, -a.value)
		case "E":
			pos += complex(a.value, 0)
		case "W":
			pos += complex(-a.value, 0)
		case "L":
			x := int(a.value / 90)
			for i := 0; i < x; i++ {
				heading *= 1i
			}
		case "R":
			x := int(a.value / 90)
			for i := 0; i < x; i++ {
				heading *= -1i
			}
		case "F":
			pos = pos + complex(a.value, 0)*heading
		}
		fmt.Println(pos, heading)
	}
	fmt.Println(math.Abs(real(pos)) + math.Abs(imag(pos)))
}
