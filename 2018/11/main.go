package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var hundredsRE = regexp.MustCompile(`\d*(\d)\d\d$`)

func main() {
	fmt.Printf("Part 1: %s\n", part1(6548))
	fmt.Printf("Part 2: %s\n", part2(6548))
}

func part1(gridSerial int) string {
	cells := [300][300]int{}

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells); y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += gridSerial
			powerLevel *= rackId
			m := hundredsRE.FindStringSubmatch(strconv.Itoa(powerLevel))
			digit := 0
			if m != nil {
				digit, _ = strconv.Atoi(m[1])
			}
			powerLevel = digit
			powerLevel -= 5
			cells[x][y] = powerLevel
		}
	}
	maxFuel := 0
	maxCell := ""
	for x := 0; x < len(cells)-3; x++ {
		for y := 0; y < len(cells)-3; y++ {
			sum := 0
			for i := x; i < (x + 3); i++ {
				for j := y; j < (y + 3); j++ {
					sum += cells[i][j]
				}
			}
			if sum > maxFuel {
				maxFuel = sum
				maxCell = strconv.Itoa(x) + "," + strconv.Itoa(y)
			}
		}
	}
	return maxCell
}

func part2(gridSerial int) string {
	cells := [300][300]int{}

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells); y++ {
			rackId := x + 10
			powerLevel := rackId * y
			powerLevel += gridSerial
			powerLevel *= rackId
			m := hundredsRE.FindStringSubmatch(strconv.Itoa(powerLevel))
			digit := 0
			if m != nil {
				digit, _ = strconv.Atoi(m[1])
			}
			powerLevel = digit
			powerLevel -= 5
			cells[x][y] = powerLevel
		}
	}
	maxFuel := 0
	maxCell := ""
	maxSize := 0
	for m := 0; m < 300; m++ {
		for x := 0; x < 300-m; x++ {
			for y := 0; y < 300-m; y++ {
				sum := 0
				for i := x; i < (x + m); i++ {
					for j := y; j < (y + m); j++ {
						sum += cells[i][j]
					}
				}
				if sum > maxFuel {
					maxFuel = sum
					maxSize = m
					maxCell = strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(maxSize)
				}
			}
		}
	}
	return maxCell
}
