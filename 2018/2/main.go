package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	contents := readInputFile(os.Args[1])
	checksum := part1(contents)
	fmt.Printf("Part 1: %d\n", checksum)
	fmt.Printf("Part 2: %s\n", part2(contents))
	fmt.Printf("Part 2a: %s\n", part2Alternative(contents))
}

func part1(contents *[]string) int {
	var countTwos, countThrees int
	for _, v := range *contents {
		twos, threes := hasTwosOrThrees(v)
		if twos {
			countTwos++
		}
		if threes {
			countThrees++
		}
	}
	return countTwos * countThrees
}

func hasTwosOrThrees(s string) (bool, bool) {
	var twos, threes bool
	charCounter := map[rune]int{}
	for _, c := range s {
		if _, ok := charCounter[c]; ok {
			charCounter[c]++
		} else {
			charCounter[c] = 1
		}
	}
	for _, c := range charCounter {
		if c == 2 {
			twos = true
		} else if c == 3 {
			threes = true
		}
	}
	return twos, threes
}

func part2Alternative(contents *[]string) string {

	seen := map[string]struct{}{}

	for _, s := range *contents {
		r := []rune(s)
		for i := range r {
			pair := string(r[:i]) + "-" + string(r[(i+1):])
			if _, ok := seen[pair]; ok {
				return strings.Replace(pair, "-", "", 1)
			}
			seen[pair] = struct{}{}
		}
	}
	return ""
}

func part2(contents *[]string) string {
	for _, s := range *contents {
		for _, t := range *contents {
			index := charMismatchIndex(s, t)
			if index >= 0 {
				// Make a copy of the string, less the mismatching character
				currentRune := []rune(s)
				newRune := make([]rune, len(currentRune)-1)
				j := 0
				for i, v := range currentRune {
					if i == index {
						continue
					}
					newRune[j] = v
					j++
				}
				return string(newRune)
			}
		}
	}
	return ""
}

func charMismatchIndex(s1, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)
	var diff, index int
	for i := 0; i < len(s1); i++ {
		if r1[i] != r2[i] {
			diff++
			index = i
		}
		// Return early if there is more than one character mismatch
		if diff > 1 {
			return -1
		}
	}
	// Return false if there is an exact match
	if diff == 0 {
		return -1
	}
	return index
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
