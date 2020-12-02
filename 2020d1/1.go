package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func inputLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

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
	lines := inputLines("1.in")
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
