package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

var inputRE = regexp.MustCompile(`([\.#]+)\s+=>\s+([\.#])`)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents *[]string) int {
	initialState := []rune(".........." + (*contents)[0] + "..........")
	mapper := map[string]string{}
	for _, line := range (*contents)[1:] {
		m := inputRE.FindStringSubmatch(line)
		if m != nil {
			mapper[m[1]] = m[2]
		}
	}

	var nextState = make([]rune, len(initialState))
	copy(nextState, initialState)
	for i := 0; i < 20; i++ {
		for j := 2; j < len(initialState)-3; j++ {
			match := false
			for k, v := range mapper {
				if bytes.Equal([]byte(string(initialState[j-2:j+3])), []byte(k)) {
					nextState[j] = ([]rune(v))[0]
					match = true
					break
				}
			}
			if nextState[len(nextState)-5] == '#' {
				nextState = append(nextState, []rune("..........")...)
				initialState = append(initialState, []rune("..........")...)
			}
			if !match {
				nextState[j] = '.'
			}
		}
		copy(initialState, nextState)
	}

	sum := 0
	for k, v := range nextState {
		if v == '#' {
			sum += (k - 10)
		}
	}
	return sum
}

func part2(contents *[]string) int {
	initialState := []rune(".........." + (*contents)[0] + "..........")
	mapper := map[string]string{}
	for _, line := range (*contents)[1:] {
		m := inputRE.FindStringSubmatch(line)
		if m != nil {
			mapper[m[1]] = m[2]
		}
	}

	var nextState = make([]rune, len(initialState))
	copy(nextState, initialState)
	generations := 0
	for i := 0; i < 1000; i++ {
		for j := 2; j < len(initialState)-3; j++ {
			match := false
			for k, v := range mapper {
				if bytes.Equal([]byte(string(initialState[j-2:j+3])), []byte(k)) {
					nextState[j] = ([]rune(v))[0]
					match = true
					break
				}
			}
			if nextState[len(nextState)-5] == '#' {
				nextState = append(nextState, []rune("..........")...)
				initialState = append(initialState, []rune("..........")...)
			}
			if !match {
				nextState[j] = '.'
			}
		}
		//Reached a steady state
		if bytes.Equal([]byte(string(initialState[:len(initialState)-1])), []byte(string(nextState[1:]))) {
			generations = i - 1
			break
		}

		copy(initialState, nextState)
	}

	sum := 0
	for k, v := range nextState {
		if v == '#' {
			sum += (k - 10) + (50000000000 - generations)
		}
	}
	return sum
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
