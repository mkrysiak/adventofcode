package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// math.Abs() takes a float
type Coordinate struct {
	x float64
	y float64
}

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %f\n", part1(stringToFloatCoordinates(contents)))
	fmt.Printf("Part 2: %f\n", part2(stringToFloatCoordinates(contents)))
}

func part1(contents *[]Coordinate) float64 {
	var board [357][357]Coordinate
	area := map[Coordinate]float64{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			board[i][j] = closestCoordinate(contents, Coordinate{float64(i), float64(j)})
			area[board[i][j]]++
		}
	}

	for k, _ := range edges(board) {
		delete(area, k)
	}

	var largestArea float64
	for _, v := range area {
		if v > largestArea {
			largestArea = v
		}
	}
	return largestArea
}

func part2(contents *[]Coordinate) float64 {
	boardSize := 357
	var area float64

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if sumOfManhattanDistances(contents, Coordinate{float64(i), float64(j)}) < 10000 {
				area++
			}
		}
	}

	return area
}

func sumOfManhattanDistances(contents *[]Coordinate, src Coordinate) float64 {
	var sum float64
	for _, v := range *contents {
		sum += manhattanDistance(src, v)
	}
	return sum
}

func closestCoordinate(contents *[]Coordinate, src Coordinate) Coordinate {
	c := map[float64]Coordinate{}
	distance := []float64{}
	for _, dst := range *contents {
		m := manhattanDistance(src, dst)
		c[m] = dst
		distance = append(distance, m)
	}
	sort.Float64s(distance)
	if distance[0] == distance[1] {
		return Coordinate{-1, -1}
	}

	return c[distance[0]]
}

func edges(board [357][357]Coordinate) map[Coordinate]struct{} {

	e := map[Coordinate]struct{}{}
	for j := 0; j < len(board); j++ {
		e[board[0][j]] = struct{}{}
		e[board[len(board)-1][j]] = struct{}{}
	}

	for i := 0; i < len(board); i++ {
		e[board[i][0]] = struct{}{}
		e[board[i][len(board)-1]] = struct{}{}
	}
	return e
}

func manhattanDistance(src Coordinate, dst Coordinate) float64 {
	return math.Abs(src.x-dst.x) + math.Abs(src.y-dst.y)
}

func stringToFloatCoordinates(contents *[]string) *[]Coordinate {
	c := []Coordinate{}
	for _, v := range *contents {
		s := strings.Split(v, ", ")
		x, _ := strconv.ParseFloat(s[0], 64)
		y, _ := strconv.ParseFloat(s[1], 64)
		c = append(c, Coordinate{x, y})
	}
	return &c
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
