package main

import (
	"fmt"
	"strconv"
)

func chomps(s string) []string {
	chomps := make([]string, 0)
	if len(s) == 1 {
		return []string{s}
	}
	start := 0
	for i := range s {
		if s[start] != s[i] {
			chomps = append(chomps, s[start:i])
			start = i
		}
	}
	chomps = append(chomps, s[start:])
	return chomps
}

func LookSay(s string) string {
	lookSay := ""
	for _, s := range chomps(s) {
		lookSay = lookSay + strconv.Itoa(len(s)) + string(s[0])
	}
	return lookSay
}

func main() {
	input := "3113322113"
	for i := 0; i < 40; i++ {
		fmt.Println(i, len(input))
		input = LookSay(input)
	}
	fmt.Println("Part 1:", len(input))
}
