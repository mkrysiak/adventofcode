package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	freq := 0
// 	loop := false
// 	dup := map[int]struct{}{}

// 	file, _ := os.Open("input")
// 	defer file.Close()
// 	reader := bufio.NewReader(file)

// 	for {
// 		str, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			if !loop {
// 				fmt.Printf("Part 1: %d\n", freq)
// 			}
// 			loop = true
// 			file.Seek(0, 0)
// 			continue
// 		}

// 		str = strings.TrimSpace(str)
// 		val, err := strconv.Atoi(str)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		freq += val

// 		if _, ok := dup[freq]; ok {
// 			fmt.Printf("Part 2: %d\n", freq)
// 			break
// 		}

// 		dup[freq] = struct{}{}
// 	}
// }

func main() {
	contents := readInputFile(os.Args[1])
	frequency := part1(contents)
	duplicate := part2(contents)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", frequency, duplicate)
}

func part1(contents *[]string) int {
	freq := 0
	for _, line := range *contents {

		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to convert string to int: %s\n", err)
		}
		freq += val
	}
	return freq

}

func part2(contents *[]string) int {

	freq := 0
	dup := map[int]struct{}{}

	dup[0] = struct{}{}

	for {
		for _, line := range *contents {

			val, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to convert string to int: %s\n", err)
			}

			freq += val

			if _, ok := dup[freq]; ok {
				return freq
			}
			dup[freq] = struct{}{}
		}
	}
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
