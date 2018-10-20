package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Bot struct {
	High, Low int
}

func inputLines(absPath string) []string {
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines = append(lines, line)
	}
	return lines
}

func main() {
	fmt.Println("Running", os.Args[0])

	absFilePath := os.Args[1]
	fmt.Println("Reading file", absFilePath)
	lines := inputLines(absFilePath)

	bots := make(map[int]Bot)
	for _, line := range lines {
		fmt.Println(line)
		cols := strings.Split(line, " ")

	}
}
