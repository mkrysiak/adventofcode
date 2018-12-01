package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var freq, loop int
	dup := make(map[int]struct{})

	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			if loop == 0 {
				fmt.Printf("Part 1: %d\n", freq)
			}
			loop++
			file.Seek(0, 0)
			continue
		}

		str = strings.TrimSpace(str)
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
		}

		freq += val

		if _, ok := dup[freq]; ok {
			fmt.Printf("Part 2: %d\n", freq)
			break
		} else {
			dup[freq] = struct{}{}
		}
	}
}
