package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	"eqri": eqri,
	"eqrr": eqrr,
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 1a: %d\n", part1Alt())
	fmt.Printf("Part 2: %d\n", part2()) // Brute force
	fmt.Printf("Part 2a: %d\n", part2Alt())
}

func part1() int {
	registers := [6]int{0, 0, 0, 0, 0, 0}
	lines := strings.Split(input(), "\n")
	instrReg, _ := strconv.Atoi(strings.Fields(lines[0])[1])
	lines = lines[1:]
	for registers[instrReg] < len(lines) {
		f := strings.Fields(string(lines[registers[instrReg]]))
		op := f[0]
		in1, _ := strconv.Atoi(f[1])
		in2, _ := strconv.Atoi(f[2])
		out, _ := strconv.Atoi(f[3])

		if op == "eqrr" {
			return registers[4]
		}
		// fmt.Printf("ip=%d [%d, %d, %d, %d, %d, %d] ", registers[instrReg], registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		opCodes[op].(func([4]int, *[6]int))([4]int{0, in1, in2, out}, &registers)
		// fmt.Printf("%s %d %d %d [%d, %d, %d, %d, %d, %d]\n", op, in1, in2, out, registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		registers[instrReg]++
	}
	return 0
}

func part2() int {
	registers := [6]int{0, 0, 0, 0, 0, 0}
	lines := strings.Split(input(), "\n")
	instrReg, _ := strconv.Atoi(strings.Fields(lines[0])[1])
	lines = lines[1:]
	seen := map[int]struct{}{}
	prev := 0
	for registers[instrReg] < len(lines) {
		f := strings.Fields(string(lines[registers[instrReg]]))
		op := f[0]
		in1, _ := strconv.Atoi(f[1])
		in2, _ := strconv.Atoi(f[2])
		out, _ := strconv.Atoi(f[3])

		if op == "eqrr" {
			if _, ok := seen[registers[4]]; !ok {
				seen[registers[4]] = struct{}{}
				prev = registers[4]
			} else {
				return prev
			}
		}
		// fmt.Printf("ip=%d [%d, %d, %d, %d, %d, %d] ", registers[instrReg], registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		opCodes[op].(func([4]int, *[6]int))([4]int{0, in1, in2, out}, &registers)
		// fmt.Printf("%s %d %d %d [%d, %d, %d, %d, %d, %d]\n", op, in1, in2, out, registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
		registers[instrReg]++
	}
	return 0
}

func part1Alt() int {
	r3, r4, r5 := 0, 0, 0
	for {
		r3 = r4 | 65536
		r4 = 10552971
		for {
			r5 = r3 & 255 // Lower 8
			r4 = r5 + r4
			r4 = r4 & 16777215 // Lower 24
			r4 = r4 * 65899
			r4 = r4 & 16777215
			if r3 < 256 {
				return r4
			}
			r3 = (r3 / 256)
			// else {
			// r5 = 0
			// for {
			// 	r5++
			// 	r2 = r5 * 256
			// 	if r2 > r3 {
			// 		r3 = r5
			// 		break
			// 	}
			// }
		}
	}
}

func part2Alt() int {
	r3, r4, r5 := 0, 0, 0
	seen := map[int]struct{}{}
	prev := 0
	for {
		r3 = r4 | 65536
		r4 = 10552971
		for {
			r5 = r3 & 255 // Lower 8
			r4 = r5 + r4
			r4 = r4 & 16777215 // Lower 24
			r4 = r4 * 65899
			r4 = r4 & 16777215
			if r3 < 256 {
				if _, ok := seen[r4]; !ok {
					seen[r4] = struct{}{}
					prev = r4
					break
				} else {
					return prev
				}
			}
			r3 = (r3 / 256)
		}
	}
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

func input() string {
	return `#ip 1
seti 123 0 4
bani 4 456 4
eqri 4 72 4
addr 4 1 1
seti 0 0 1
seti 0 2 4
bori 4 65536 3
seti 10552971 1 4
bani 3 255 5
addr 4 5 4
bani 4 16777215 4
muli 4 65899 4
bani 4 16777215 4
gtir 256 3 5
addr 5 1 1
addi 1 1 1
seti 27 7 1
seti 0 1 5
addi 5 1 2
muli 2 256 2
gtrr 2 3 2
addr 2 1 1
addi 1 1 1
seti 25 0 1
addi 5 1 5
seti 17 2 1
setr 5 7 3
seti 7 8 1
eqrr 4 0 5
addr 5 1 1
seti 5 0 1`
}
