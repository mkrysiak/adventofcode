package main

import (
	"fmt"
	"strconv"
	"strings"
)

var registers = [6]int{}
var opCodes = map[string]interface{}{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"ewri": eqri,
	"eqrr": eqrr,
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(input()))
	registers = [6]int{1, 0, 0, 0, 0, 0}
	fmt.Printf("Part 2: %d\n", part2(input()))
}

func part1(in string) int {
	lines := strings.Split(in, "\n")
	instrReg, _ := strconv.Atoi(strings.Fields(lines[0])[1])
	lines = lines[1:]

	for registers[instrReg] < len(lines) {
		f := strings.Fields(string(lines[registers[instrReg]]))
		op := f[0]
		in1, _ := strconv.Atoi(f[1])
		in2, _ := strconv.Atoi(f[2])
		out, _ := strconv.Atoi(f[3])

		// fmt.Printf("ip=%d [%d, %d, %d, %d, %d, %d] ", registers[instrReg], registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		opCodes[op].(func([4]int, *[6]int))([4]int{0, in1, in2, out}, &registers)
		// fmt.Printf("%s %d %d %d [%d, %d, %d, %d, %d, %d]\n", op, in1, in2, out, registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		registers[instrReg]++
	}
	return registers[0]
}

func part2(in string) int {
	lines := strings.Split(in, "\n")
	instrReg, _ := strconv.Atoi(strings.Fields(lines[0])[1])
	lines = lines[1:]

	accumulator := 0
	for registers[instrReg] < len(lines) {
		f := strings.Fields(string(lines[registers[instrReg]]))
		op := f[0]
		in1, _ := strconv.Atoi(f[1])
		in2, _ := strconv.Atoi(f[2])
		out, _ := strconv.Atoi(f[3])

		// fmt.Printf("ip=%d [%d, %d, %d, %d, %d, %d] ", registers[instrReg], registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		opCodes[op].(func([4]int, *[6]int))([4]int{0, in1, in2, out}, &registers)
		// fmt.Printf("%s %d%d %d [%d, %d, %d, %d, %d, %d]\n", op, in1, in2, out, registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])

		// Study the program and output, to identify the pattern.  See the "output" file.
		// Run through the first few operations to initalize register[4], and once the
		// program-counter == 1, register[4] has been initalized.
		if registers[instrReg] == 1 {
			for x := 1; x <= registers[4]; x++ {
				if registers[4]%x == 0 {
					accumulator += x
				}
			}
			break
		}
		registers[instrReg]++
	}
	return accumulator
}

func addr(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] + (*r)[i[2]]
}

func addi(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] + i[2]
}

func mulr(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] * (*r)[i[2]]
}

func muli(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] * i[2]
}

func banr(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] & (*r)[i[2]]
}

func bani(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] & i[2]
}

func borr(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] | (*r)[i[2]]
}

func bori(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]] | i[2]
}

func setr(i [4]int, r *[6]int) {
	(*r)[i[3]] = (*r)[i[1]]
}

func seti(i [4]int, r *[6]int) {
	(*r)[i[3]] = i[1]
}

func gtir(i [4]int, r *[6]int) {
	if i[1] > (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func gtri(i [4]int, r *[6]int) {
	if (*r)[i[1]] > i[2] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func gtrr(i [4]int, r *[6]int) {
	if (*r)[i[1]] > (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqir(i [4]int, r *[6]int) {
	if i[1] == (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqri(i [4]int, r *[6]int) {
	if (*r)[i[1]] == i[2] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqrr(i [4]int, r *[6]int) {
	if (*r)[i[1]] == (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func testInput() string {
	return `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`
}

func input() string {
	return `#ip 1
addi 1 16 1
seti 1 1 3
seti 1 9 5
mulr 3 5 2
eqrr 2 4 2
addr 2 1 1
addi 1 1 1
addr 3 0 0
addi 5 1 5
gtrr 5 4 2
addr 1 2 1
seti 2 6 1
addi 3 1 3
gtrr 3 4 2
addr 2 1 1
seti 1 6 1
mulr 1 1 1
addi 4 2 4
mulr 4 4 4
mulr 1 4 4
muli 4 11 4
addi 2 6 2
mulr 2 1 2
addi 2 2 2
addr 4 2 4
addr 1 0 1
seti 0 3 1
setr 1 4 2
mulr 2 1 2
addr 1 2 2
mulr 1 2 2
muli 2 14 2
mulr 2 1 2
addr 4 2 4
seti 0 0 0
seti 0 4 1`
}
