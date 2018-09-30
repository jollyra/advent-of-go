package main

import (
	"bufio"
	"encoding/json"
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

func sumJSONWithIgnore(data []byte) int {
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		log.Fatalln("Failed unmarshalling json", err)
	}
	sum := 0
	traverseSum(f, &sum)
	return int(sum)
}

func traverseSum(json interface{}, sum *int) {
	switch vv := json.(type) {
	case string:
	case float64:
		*sum = *sum + int(vv)
	case []interface{}:
		for _, v := range vv {
			traverseSum(v, sum)
		}
	case map[string]interface{}:
		ignore := false
		for _, v := range vv {
			switch t := v.(type) {
			case string:
				if t == "red" {
					ignore = true
				}
			}
		}
		if ignore == true {
			return
		}
		for _, v := range vv {
			traverseSum(v, sum)
		}
	default:
		fmt.Println(vv, "is of a type I don't know how to handle")
	}
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
	input := inputLine()
	fmt.Println("Part 1:", sum(input))
	fmt.Println("Part 2:", sumJSONWithIgnore([]byte(input)))
}
