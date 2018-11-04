package main

import (
	"fmt"
	"github.com/jollyra/stringutil"
	"os"
	"regexp"
	"strings"
)

type mutation struct {
	Before, After string
}

func (m mutation) String() string {
	return fmt.Sprintf("%s=>%s", m.Before, m.After)
}

func mutator(mutations []mutation, str string) []string {
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

func calibrate(mutations []mutation, molecule string) int {
	results := mutator(mutations, molecule)
	unique := stringutil.Unique(results)
	return len(unique)
}

func main() {
	var molecule = "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	lines := stringutil.InputLines(os.Args[1])
	mutations := parseMutations(lines)

	num := calibrate(mutations, molecule)
	fmt.Println("Calibrating... # of num molecules is", num)

	// Part 2 solved by hand!
}
