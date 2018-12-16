package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Operation struct {
	before [4]int
	after  [4]int
	op     [4]int
}

type Operations []Operation

var registers = [4]int{}
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

var beforeRE = regexp.MustCompile(`^Before:\s+\[(\d+),\s+(\d+),\s+(\d+),\s+(\d+)\]`)
var afterRE = regexp.MustCompile(`^After:\s+\[(\d+),\s+(\d+),\s+(\d+),\s+(\d+)\]`)
var inputRE = regexp.MustCompile(`^(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents *[]string) int {
	ops := parseTestCases(contents)

	gtrThreeCount := 0
	for _, o := range ops {
		count := 0
		for fname := range opCodes {
			if testFuncEquality(fname, o) {
				count++
			}
		}
		if count >= 3 {
			gtrThreeCount++
		}
	}
	// for _, v := range ops {
	// 	fmt.Printf("Before: %v Input: %v After: %v\n", v.before, v.op, v.after)
	// }

	return gtrThreeCount
}

func part2(contents *[]string) int {
	ops := parseTestCases(contents)
	opCodeMap := map[string]int{}

	for len(opCodeMap) != 16 {
		for _, o := range ops {
			count := 0
			matches := map[string]int{}
			for fname := range opCodes {
				if testFuncEquality(fname, o) {
					if _, ok := opCodeMap[fname]; !ok {
						matches[fname] = o.op[0]
						count++
					}
				}
			}

			if len(matches) == 1 {
				for k, v := range matches {
					opCodeMap[k] = v
				}
			}
		}
	}

	inverseOpCodeMap := map[int]string{}
	for k, v := range opCodeMap {
		inverseOpCodeMap[v] = k
	}

	for _, op := range parseTestProgram(contents) {
		opName := inverseOpCodeMap[op[0]]
		opCodes[opName].(func([4]int, *[4]int))(op, &registers)
	}

	return registers[0]
}

func testFuncEquality(fname string, op Operation) bool {
	register := op.before
	opCodes[fname].(func([4]int, *[4]int))(op.op, &register)
	return testSliceEquality(op.after, register)
}

func testSliceEquality(a, b [4]int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func addr(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] + (*r)[i[2]]
}

func addi(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] + i[2]
}

func mulr(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] * (*r)[i[2]]
}

func muli(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] * i[2]
}

func banr(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] & (*r)[i[2]]
}

func bani(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] & i[2]
}

func borr(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] | (*r)[i[2]]
}

func bori(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]] | i[2]
}

func setr(i [4]int, r *[4]int) {
	(*r)[i[3]] = (*r)[i[1]]
}

func seti(i [4]int, r *[4]int) {
	(*r)[i[3]] = i[1]
}

func gtir(i [4]int, r *[4]int) {
	if i[1] > (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func gtri(i [4]int, r *[4]int) {
	if (*r)[i[1]] > i[2] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func gtrr(i [4]int, r *[4]int) {
	if (*r)[i[1]] > (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqir(i [4]int, r *[4]int) {
	if i[1] == (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqri(i [4]int, r *[4]int) {
	if (*r)[i[1]] == i[2] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func eqrr(i [4]int, r *[4]int) {
	if (*r)[i[1]] == (*r)[i[2]] {
		(*r)[i[3]] = 1
		return
	}
	(*r)[i[3]] = 0
}

func parseTestCases(contents *[]string) Operations {
	ops := Operations{}

	chunk := false
	op := Operation{}
	for _, v := range *contents {
		b := beforeRE.FindStringSubmatch(v)
		if chunk == false && b != nil {
			chunk = true
			in := [4]int{}
			for j := 1; j < len(b); j++ {
				in[j-1], _ = strconv.Atoi(b[j])
			}
			op.before = in
		}
		i := inputRE.FindStringSubmatch(v)
		if chunk != false && i != nil {
			in := [4]int{}
			for j := 1; j < len(i); j++ {
				in[j-1], _ = strconv.Atoi(i[j])
			}
			op.op = in
		}
		a := afterRE.FindStringSubmatch(v)
		if chunk != false && a != nil {
			chunk = false
			in := [4]int{}
			for j := 1; j < len(a); j++ {
				in[j-1], _ = strconv.Atoi(a[j])
			}
			op.after = in
			ops = append(ops, op)
		}
	}
	return ops
}

func parseTestProgram(contents *[]string) [][4]int {
	ops := [][4]int{}

	chunk := false
	for _, v := range *contents {
		b := beforeRE.FindStringSubmatch(v)
		if chunk == false && b != nil {
			chunk = true
		}
		i := inputRE.FindStringSubmatch(v)
		if chunk == false && i != nil {
			in := [4]int{}
			for j := 1; j < len(i); j++ {
				in[j-1], _ = strconv.Atoi(i[j])
			}
			ops = append(ops, in)
		}
		a := afterRE.FindStringSubmatch(v)
		if chunk != false && a != nil {
			chunk = false
		}
	}
	return ops
}

func readInputFile(infile string) *[]string {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}
	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return &contents
}
