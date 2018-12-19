package main

type Instruction struct {
	Opcode int
	A      int
	B      int
	Out    int
}

func Addr(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] + regs[ins.B]
	return regs
}

func Addi(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] + ins.B
	return regs
}

func Mulr(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] * regs[ins.B]
	return regs
}

func Muli(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] * ins.B
	return regs
}

func Banr(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] & regs[ins.B]
	return regs
}

func Bani(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] & ins.B
	return regs
}

func Borr(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] | regs[ins.B]
	return regs
}

func Bori(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A] | ins.B
	return regs
}

func Setr(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = regs[ins.A]
	return regs
}

func Seti(ins Instruction, regs [4]int) [4]int {
	regs[ins.Out] = ins.A
	return regs
}

func Gtir(ins Instruction, regs [4]int) [4]int {
	if ins.A > regs[ins.B] {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

func Gtri(ins Instruction, regs [4]int) [4]int {
	if regs[ins.A] > ins.B {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

func Gtrr(ins Instruction, regs [4]int) [4]int {
	if regs[ins.A] > regs[ins.B] {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

func Eqir(ins Instruction, regs [4]int) [4]int {
	if ins.A == regs[ins.B] {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

func Eqri(ins Instruction, regs [4]int) [4]int {
	if regs[ins.A] == ins.B {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

func Eqrr(ins Instruction, regs [4]int) [4]int {
	if regs[ins.A] == regs[ins.B] {
		regs[ins.Out] = 1
	} else {
		regs[ins.Out] = 0
	}
	return regs
}

var OpcodeFuncMap = map[int]func(Instruction, [4]int) [4]int{
	8:  Addr,
	5:  Bori,
	12: Addi,
	14: Mulr,
	10: Muli,
	11: Banr,
	1:  Bani,
	9:  Borr,
	7:  Setr,
	6:  Seti,
	3:  Gtir,
	0:  Gtri,
	15: Gtrr,
	4:  Eqir,
	13: Eqri,
	2:  Eqrr,
}
