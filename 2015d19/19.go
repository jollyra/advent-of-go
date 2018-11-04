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

func reverseMutation(m mutation) mutation {
	return mutation{m.After, m.Before}
}

func reverseMutations(ms []mutation) []mutation {
	reversed := make([]mutation, 0)
	for _, m := range ms {
		reversed = append(reversed, reverseMutation(m))
	}
	return reversed
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

func reverseEngineer(mutations []mutation, molecule string) int {
	mutations = reverseMutations(mutations)
	fmt.Printf("Reverse Engineering %s with %d mutations\n",
		molecule, len(mutations))
	gen := []string{molecule}
	for i := 0; i < 100; i++ {
		fmt.Println("step", i)
		nextGen := make([]string, 0)
		gen = stringutil.Unique(gen)
		for _, mol := range gen {
			if mol == "e" {
				return i
			}
			nextGen = append(nextGen, mutator(mutations, mol)...)
		}
		gen = nextGen
	}
	return -1
}

func main() {
	var molecule = "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl"
	lines := stringutil.InputLines(os.Args[1])
	mutations := parseMutations(lines)

	num := calibrate(mutations, molecule)
	fmt.Println("Calibrating... # of num molecules is", num)

	steps := reverseEngineer(mutations, molecule)
	fmt.Printf("%d steps to reverse engineer molecule\n", steps)
}
