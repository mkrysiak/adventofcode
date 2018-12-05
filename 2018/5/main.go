package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents[0]))
}

func part1(contents string) int {
	runes := []rune(contents)
	var i int
	for i < len(runes)-1 {
		diff := runes[i] - runes[i+1]
		if diff == 32 || diff == -32 {
			// fmt.Printf("%s %s\n", string(runes[i]), string(runes[i+1]))
			if i == 0 {
				runes = runes[2:]
				i = 0
				continue
			}
			runes = append(runes[:i], runes[i+2:]...)
			i = i - 1
			continue
		}
		i++

	}
	return len(runes)
}

func readInputFile(infile string) []string {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}
	var contents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
