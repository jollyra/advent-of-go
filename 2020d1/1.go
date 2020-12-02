package main

import (
	"fmt"
	"strconv"

	"github.com/jollyra/go-advent-util"
)

func ints(ss []string) (xs []int) {
	for _, s := range ss {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		xs = append(xs, x)
	}
	return xs
}

func main() {
	lines := advent.InputLines("1.in")
	xs := ints(lines)
	for i := range xs {
		for j := range xs {
			for k := range xs {
				if xs[i]+xs[j]+xs[k] == 2020 {
					fmt.Println(xs[i] * xs[j] * xs[k])
					return
				}
			}
		}
	}
}
