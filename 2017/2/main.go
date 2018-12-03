package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	contents := readInputFile("input")
	fmt.Printf("%d\n", part1(contents))
}

func part1(contents *[]string) int {
	var sum int
	for _, v := range *contents {
		strFields := strings.Fields(v)
		var intFields []int
		for _, vv := range strFields {
			f, _ := strconv.Atoi(vv)
			intFields = append(intFields, f)
		}
		sort.Ints(intFields)
		sum += (intFields[len(intFields)-1] - intFields[0])
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
