package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var print = fmt.Println

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

type sample struct {
	Before [4]int
	Ins    Instruction
	After  [4]int
}

func parseSamples(lines []string) []sample {
	samples := make([]sample, 0)
	s := &sample{}
	for i, line := range lines {
		var a, b, c, d int
		switch i % 4 {
		case 0:
			fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)
			s.Before = [4]int{a, b, c, d}
		case 1:
			fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)
			ins := Instruction{a, b, c, d}
			s.Ins = ins
		case 2:
			fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &a, &b, &c, &d)
			s.After = [4]int{a, b, c, d}
		case 3:
			samples = append(samples, *s)
			s = &sample{}
		}
	}
	return samples
}

func equal(m1, m2 [4]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}

var funcs = map[string]func(Instruction, [4]int) [4]int{
	"addr": Addr,
	"addi": Addi,
	"mulr": Mulr,
	"muli": Muli,
	"banr": Banr,
	"bani": Bani,
	"borr": Borr,
	"bori": Bori,
	"setr": Setr,
	"seti": Seti,
	"gtir": Gtir,
	"gtri": Gtri,
	"gtrr": Gtrr,
	"eqir": Eqir,
	"eqri": Eqri,
	"eqrr": Eqrr,
}

type counter map[string]int

func newCounter() counter { return make(map[string]int) }

func testSamples() {
	lines := inputLines(os.Args[1])
	samples := parseSamples(lines)
	for _, s := range samples {
		matches := 0
		var b strings.Builder
		fmt.Fprintf(&b, "before %v ", s.Before)
		fmt.Fprintf(&b, "Ins %v ", s.Ins)
		fmt.Fprintf(&b, "after %v", s.After)
		for name, fn := range funcs {
			_, inDecided := OpcodeFuncMap[s.Ins.Opcode]
			if !inDecided {
				got := fn(s.Ins, s.Before)
				if equal(got, s.After) {
					matches++
					fmt.Fprintf(&b, " %s ", name)
				}
			}
		}
		fmt.Println(b.String())

		if matches == 1 {
			print("Single match above")
		}
	}
}

func main() {
	testSamples()
}
