package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	in := strings.Split(input(), "-")
	start, _ := strconv.Atoi(in[0])
	end, _ := strconv.Atoi(in[1])
	s := time.Now()
	count := 0
	for i := start; i <= end; i++ {
		if isIncreasingAndHasDouble(i) {
			count++
		}
	}

	fmt.Println(count)

	count = 0
	for i := start; i <= end; i++ {
		if isIncreasingAndHasDouble(i) && hasAtleastOneDouble(i) {
			count++
		}
	}

	elapsed := time.Since(s)
	fmt.Println(count, elapsed)
}

func isIncreasingAndHasDouble(in int) bool {
	prevVal := 10 // Any value over 9
	length := 0
	for in > 0 {
		currVal := in % 10
		in /= 10
		if prevVal == currVal {
			length++
		} else if prevVal < currVal {
			return false
		}
		prevVal = currVal
	}

	return length > 0
}

func hasAtleastOneDouble(in int) bool {
	inStr := strconv.Itoa(in)
	if len(inStr) < 2 {
		return false
	}

	found := false
	var start, end, count int
	for i := 1; i < len(inStr); i++ {
		if inStr[i] != inStr[i-1] {
			count = end - start
			if count == 1 {
				found = true
			}
			start, end = i, i
		} else if inStr[i] == inStr[i-1] {
			if i != len(inStr)-1 {
				end = i
			} else if i == len(inStr)-1 {
				end = i
				count = end - start
				if count == 1 {
					found = true
				}
			}
		}
	}

	return found
}

func input() string {
	return `235741-706948`
}
