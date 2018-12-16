package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents [4]int) int {
	var sum, last int
	for i, v := range contents {
		if i == 0 {
			last = contents[len(contents)-1]
		}
		if v == last {
			sum += v
		}
		last = v
	}
	return sum
}

func part2(contents [4]int) int {
	var sum, last int
	step := len(contents) / 2
	for i, v := range contents {
		if i == 0 {
			last = contents[len(contents)-1]
		}
		if contents[(i+step-1)%len(contents)] == last {
			sum += last
		}
		last = v
	}
	return sum
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func readInputFile(infile string) [4]int {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}
	var contents [4]int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		contents = append(contents, toInt(scanner.Bytes()))
	}

	return contents
}
