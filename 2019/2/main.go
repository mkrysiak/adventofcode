package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(stoi()))
	fmt.Println(part2(stoi()))
}

func part1(input []int) int {
	in := input
	ptr := 0
	for in[ptr] != 99 && ptr < len(in) {
		if in[ptr] == 1 { // Add
			in[in[ptr+3]] = in[in[ptr+1]] + in[in[ptr+2]]
		} else if in[ptr] == 2 { // Multiply
			in[in[ptr+3]] = in[in[ptr+1]] * in[in[ptr+2]]
		} else {
			fmt.Println("Shouldn't be here")
		}
		ptr += 4
	}
	return in[0]
}

func part2(originalInput []int) int {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			//input := originalInput
			input := append(make([]int, 0, len(originalInput)), originalInput...)
			input[1] = i
			input[2] = j
			if part1(input) == 19690720 {
				return 100*i + j
			}
		}
	}
	return 0
}

func stoi() []int {
	inputSlice := strings.Split(input(), ",")
	in := make([]int, len(inputSlice))
	for k, v := range inputSlice {
		o, _ := strconv.Atoi(v)
		in[k] = o
	}
	in[1] = 12
	in[2] = 2
	return in
}

func input() string {
	return `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,5,23,2,13,23,27,1,10,27,31,2,6,31,35,1,9,35,39,2,10,39,43,1,43,9,47,1,47,9,51,2,10,51,55,1,55,9,59,1,59,5,63,1,63,6,67,2,6,67,71,2,10,71,75,1,75,5,79,1,9,79,83,2,83,10,87,1,87,6,91,1,13,91,95,2,10,95,99,1,99,6,103,2,13,103,107,1,107,2,111,1,111,9,0,99,2,14,0,0`
}
