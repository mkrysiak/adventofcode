package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lineRE = regexp.MustCompile(`^#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)$`)

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 1a: %d\n", part1Alt(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents *[]string) int {
	rows := map[int]map[int]struct{}{}
	overlap := map[int]map[int]struct{}{}

	for _, v := range *contents {
		m := lineRE.FindStringSubmatch(v)
		if m != nil {
			col, _ := strconv.Atoi(m[2])
			row, _ := strconv.Atoi(m[3])
			width, _ := strconv.Atoi(m[4])
			height, _ := strconv.Atoi(m[5])

			for r := row; r < row+height; r++ {
				if _, ok := rows[r]; !ok {
					rows[r] = map[int]struct{}{}
				}
				for i := col; i < col+width; i++ {
					if _, ok := rows[r][i]; ok {
						if _, ok := overlap[r]; !ok {
							overlap[r] = map[int]struct{}{}
						}
						overlap[r][i] = struct{}{}
						continue
					}
					rows[r][i] = struct{}{}
				}
			}
		}
	}

	var sum int
	for _, v := range overlap {
		sum += len(v)
	}

	return sum
}

func part1Alt(contents *[]string) int {
	// Should calculate the max width and height based on input
	table := [1100][1100]int{}
	for _, v := range *contents {
		m := lineRE.FindStringSubmatch(v)
		if m != nil {
			col, _ := strconv.Atoi(m[2])
			row, _ := strconv.Atoi(m[3])
			width, _ := strconv.Atoi(m[4])
			height, _ := strconv.Atoi(m[5])

			for r := row; r < row+height; r++ {
				for c := col; c < col+width; c++ {
					table[r][c]++
				}
			}
		}
	}

	var sum int
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i][j] > 1 {
				sum++
			}
		}
	}

	return sum
}

func part2(contents *[]string) int {
	rows := map[int]map[int]int{}
	overlap := map[int]map[int]int{}
	overlappingIds := map[int]bool{}

	for _, v := range *contents {
		m := lineRE.FindStringSubmatch(v)
		if m != nil {
			id, _ := strconv.Atoi(m[1])
			col, _ := strconv.Atoi(m[2])
			row, _ := strconv.Atoi(m[3])
			width, _ := strconv.Atoi(m[4])
			height, _ := strconv.Atoi(m[5])

			overlappingIds[id] = false

			for r := row; r < row+height; r++ {
				if _, ok := rows[r]; !ok {
					rows[r] = map[int]int{}
				}
				for i := col; i < col+width; i++ {
					if v, ok := rows[r][i]; ok {
						if _, ok := overlap[r]; !ok {
							overlap[r] = map[int]int{}
						}
						overlappingIds[v] = true
						overlappingIds[id] = true
						overlap[r][i] = id
						continue
					}
					rows[r][i] = id
				}
			}
		}
	}

	nonoverlappingId := -1
	for k, v := range overlappingIds {
		if !v {
			nonoverlappingId = k
		}
	}
	return nonoverlappingId
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
