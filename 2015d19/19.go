package main

import (
	"bufio"
	"fmt"
	"github.com/jollyra/stringutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type mutation struct {
	Before, After string
}

func (m mutation) String() string {
	return fmt.Sprintf("%s => %s", m.Before, m.After)
}

var mutations = []mutation{
	mutation{"H", "HO"},
	mutation{"H", "OH"},
	mutation{"O", "HH"},
}

func mutator(mutations []mutation, str string, n int) []string {
	results := make([]string, 0)
	bytes := []byte(str)
	for _, m := range mutations {
		pattern := regexp.MustCompile(m.Before)
		locs := pattern.FindAllIndex(bytes, -1)
		for _, loc := range locs {
			sNew := fmt.Sprintf("%s%s%s",
				str[:loc[0]], m.After, str[loc[1]:])
			results = append(results, sNew)
		}
	}
	return results
}

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

func parseMutations(lines []string) []mutation {
	mutations := make([]mutation, 0)
	for _, line := range lines {
		words := strings.Split(line, "=>")
		before := strings.Trim(words[0], " \n")
		after := strings.Trim(words[1], " \n")
		mutations = append(mutations, mutation{before, after})
	}
	return mutations
}

func main() {
	lines := inputLines(os.Args[1])
	mutations := parseMutations(lines)
	for _, m := range mutations {
		fmt.Println(m)
	}
	molecule := "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	results := mutator(mutations, molecule, 1)
	unique := stringutil.Unique(results)
	fmt.Println(unique)
	fmt.Println("# of unique molecules is", len(unique))
}
