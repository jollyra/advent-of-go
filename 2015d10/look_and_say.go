package main

import (
	"fmt"
	"strings"
)

func lookSay(s string) string {
	lookSay := ""
	var b strings.Builder
	start := 0
	for i := range s {
		if s[start] != s[i] {
			fmt.Fprintf(&b, "%s%d%s", lookSay, i-start, string(s[start]))
			start = i
		}
	}
	fmt.Fprintf(&b, "%s%d%s", lookSay, len(s)-start, string(s[len(s)-1]))
	return b.String()
}

func main() {
	input := "3113322113"
	for i := 0; i < 50; i++ {
		fmt.Println(i, len(input))
		input = lookSay(input)
	}
	fmt.Println("ans:", len(input))
}
