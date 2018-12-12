package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1(6548))
	fmt.Printf("Part 2a: %s\n", part2Alt(6548))
	fmt.Printf("Part 2: %s\n", part2(6548))
}

func part1(gridSerial int) string {
	cells := [300][300]int{}

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells); y++ {
			rackId := x + 10
			powerLevel := rackId*y + gridSerial
			powerLevel *= rackId
			powerLevel = powerLevel/100%10 - 5
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

// https://en.wikipedia.org/wiki/Summed-area_table
func part2Alt(gridSerial int) string {

	cells := [301][301]int{}

	for x := 1; x < 301; x++ {
		for y := 1; y < 301; y++ {
			rackId := x + 10
			powerLevel := rackId*y + gridSerial
			powerLevel *= rackId
			powerLevel = powerLevel/100%10 - 5
			cells[x][y] = powerLevel + cells[x][y-1] + cells[x-1][y] - cells[x-1][y-1]
		}
	}

	maxFuel := 0
	maxSize := 0
	maxX, maxY := 0, 0
	for m := 1; m < 301; m++ {
		for x := m; x < 301; x++ {
			for y := m; y < 301; y++ {
				sum := cells[x][y] - cells[x][y-m] - cells[x-m][y] + cells[x-m][y-m]
				if sum > maxFuel {
					maxFuel = sum
					maxSize = m
					maxX, maxY = x-m+1, y-m+1
				}
			}
		}
	}
	return strconv.Itoa(maxX) + "," + strconv.Itoa(maxY) + "," + strconv.Itoa(maxSize)
}

func part2(gridSerial int) string {
	cells := [300][300]int{}

	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells); y++ {
			rackId := x + 10
			powerLevel := rackId*y + gridSerial
			powerLevel *= rackId
			powerLevel = powerLevel/100%10 - 5
			cells[x][y] = powerLevel
		}
	}
	maxFuel := 0
	maxSize := 0
	maxX, maxY := 0, 0
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
					maxX, maxY = x, y
				}
			}
		}
	}
	return strconv.Itoa(maxX) + "," + strconv.Itoa(maxY) + "," + strconv.Itoa(maxSize)
}
