package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x int
	y int
}

type Coordinates []Coordinate

var data [50][50]rune

var adj = Coordinates{
	Coordinate{0, -1},
	Coordinate{1, -1},
	Coordinate{1, 0},
	Coordinate{1, 1},
	Coordinate{0, 1},
	Coordinate{-1, 1},
	Coordinate{-1, 0},
	Coordinate{-1, -1},
}

func main() {
	readInputFile("input")
	process()
}

func process() {
	newState := [50][50]rune{}
	part1State, part2State := 0, 0
	minutes := 0
	for minutes < 1000000000 {
		minutes++
		for y := range data {
			for x := range data[y] {
				// fmt.Println(string(data[x][y]))
				adjAcres := map[rune]int{
					'.': 0,
					'|': 0,
					'#': 0,
				}
				for _, a := range adj {
					xx := x + a.x
					yy := y + a.y
					if xx >= 0 && xx <= len(data)-1 &&
						yy >= 0 && yy <= len(data)-1 {
						adjAcres[data[xx][yy]]++
					}
				}
				switch data[x][y] {
				case '.':
					if adjAcres['|'] >= 3 {
						newState[x][y] = '|'
					} else {
						newState[x][y] = data[x][y]
					}
				case '|':
					if adjAcres['#'] >= 3 {
						newState[x][y] = '#'
					} else {
						newState[x][y] = data[x][y]
					}
				case '#':
					if adjAcres['#'] >= 1 && adjAcres['|'] >= 1 {
						newState[x][y] = '#'
					} else {
						newState[x][y] = '.'
					}
				}
			}
		}
		data = newState

		// Pattern repeates every 7000 iterations
		if minutes == 9 {
			metadata := getMetaData(data)
			part1State = metadata['|'] * metadata['#']
		}
		if minutes%1000 == 0 && minutes%7 == 6 {
			metadata := getMetaData(data)
			part2State = metadata['|'] * metadata['#']
			break
		}
	}
	fmt.Printf("Part 1: %d\n", part1State)
	fmt.Printf("Part 2: %d\n", part2State)
}

func printData() {
	for y := range data {
		for x := range data[y] {
			fmt.Printf("%s", string(data[x][y]))
		}
		fmt.Println()
	}
}

func getMetaData(data [50][50]rune) map[rune]int {
	metadata := map[rune]int{}
	for y := range data {
		for x := range data[y] {
			switch data[y][x] {
			case '.':
				metadata['.']++
			case '|':
				metadata['|']++
			case '#':
				metadata['#']++
			}
		}
	}
	return metadata
}

func readInputFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		b := []rune(scanner.Text())
		for i, v := range b {
			data[i][row] = v
		}
		row++
	}
}

// stateChange := false
// for y := range newState {
// 	for x := range newState[y] {
// 		if newState[x][y] != data[x][y] {
// 			stateChange = true
// 			break
// 		}
// 	}
// 	if stateChange {
// 		break
// 	}
// }
// if stateChange {
// 	data = newState
// } else {
// 	break
// }
