package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jollyra/go-advent-util"
)

var errInfiniteLoop = errors.New("infinite loop")
var errUnrecognizedInstruction = fmt.Errorf("unrecognized instruction")
var errProcUnexpectedlyEnded = errors.New("proc ended unexpectedly")

type instruction struct {
	op   string
	args []int
}

type process struct {
	program []instruction
	acc     int
	ip      int
}

func (proc *process) print() {
	for _, ins := range proc.program {
		fmt.Println(ins)
	}
}

func (proc *process) step() error {
	if proc.ip >= len(proc.program) {
		return nil
	}

	ins := proc.program[proc.ip]
	switch ins.op {
	case "acc":
		proc.acc += ins.args[0]
		proc.ip++
	case "jmp":
		proc.ip += ins.args[0]
	case "nop":
		proc.ip++
	default:
		return errUnrecognizedInstruction
	}
	return errProcUnexpectedlyEnded
}

func (proc *process) kill() {
	newProc := process{program: proc.program}
	proc = &newProc
}

func (proc *process) run() error {
	proc.kill()
	visited := make([]int, len(proc.program)+1)
	for {
		visited[proc.ip]++
		if visited[proc.ip] > 1 {
			return errInfiniteLoop
		}
		err := proc.step()
		if err == nil {
			return nil
		}
	}
}

func parse(lines []string) (prog []instruction, err error) {
	for _, line := range lines {
		ins := instruction{}
		xs := strings.Split(line, " ")
		ins.op = xs[0]
		for _, x := range xs[1:] {
			i, err := strconv.Atoi(x)
			if err != nil {
				return prog, err
			}
			ins.args = append(ins.args, i)
		}
		prog = append(prog, ins)
	}
	return prog, nil
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	program, err := parse(lines)
	proc := process{program: program}
	err = proc.run()
	if err != nil {
		if err == errInfiniteLoop {
			fmt.Println("Part 1:", proc.acc)
		} else {
			panic(err)
		}
	}

	for i := range program {
		mutatedProg := make([]instruction, len(program))
		copy(mutatedProg, program)

		switch ins := program[i].op; ins {
		case "nop":
			mutatedProg[i].op = "jmp"
		case "jmp":
			mutatedProg[i].op = "nop"
		default:
			continue
		}

		proc := process{program: mutatedProg}
		err = proc.run()
		if err != nil {
			continue
		}
		fmt.Println("Part 2:", proc.acc)
		break
	}
}
