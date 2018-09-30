package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func sum(json string) int {
	re := regexp.MustCompile("-?[0-9]+")
	match := re.FindAllString(json, -1)
	if match == nil {
		return 0
	}
	count := 0
	for _, x := range match {
		i, _ := strconv.Atoi(x)
		count += i
	}
	return count
}

func inputLine() string {
	file, err := os.Open("/Users/nrahkola/go/src/github.com/jollyra/" +
		"advent-of-go/js_abacus_framework/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println(sum(inputLine()))
}
