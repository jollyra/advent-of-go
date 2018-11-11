package main

import (
	"fmt"
)

const lookupSize uint32 = 1<<20 - 1

func buildLookupV0() *[lookupSize]uint32 {
	lookup := [lookupSize]uint32{}
	var i uint32
	for i = 1; i < lookupSize; i++ {
		var j uint32
		for j = i; j < lookupSize; j += i {
			lookup[j] += i
		}
	}
	return &lookup
}

func buildLookupV1() *[lookupSize]uint32 {
	lookup := [lookupSize]uint32{}
	var i uint32
	for i = 1; i < lookupSize; i++ {
		var j uint32
		for j = i; j < i+i*50 && j < lookupSize; j += i {
			lookup[j] += i * 11
		}
	}
	return &lookup
}

func main() {
	part1 := buildLookupV0()
	for i, x := range part1 {
		if x >= 3600000 {
			fmt.Println("Part 1:", i)
			break
		}
	}

	part2 := buildLookupV1()
	for i, x := range part2 {
		if x >= 36000000 {
			fmt.Println("Part 2:", i)
			break
		}
	}
}
