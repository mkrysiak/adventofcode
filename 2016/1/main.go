package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

var inputRE = regexp.MustCompile(`(R|L)(\d+)`)

const (
	north int = iota
	south
	east
	west
)

func main() {
	location := Coord{}
	direction := north
	visited := map[Coord]struct{}{}
	visitedTwice := Coord{}
	seen := false
	for _, v := range strings.Split(input(), ", ") {
		m := inputRE.FindStringSubmatch(v)
		if m != nil {
			dist, _ := strconv.Atoi(m[2])
			direction = nextDirection(direction, m[1])
			switch direction {
			case north:
				for i := 0; i < dist; i++ {
					location.y++
					if !seen {
						if _, ok := visited[location]; ok {
							visitedTwice = location
							seen = true
						}
						visited[location] = struct{}{}
					}
				}
			case south:
				for i := 0; i < dist; i++ {
					location.y--
					if !seen {
						if _, ok := visited[location]; ok {
							visitedTwice = location
							seen = true
						}
						visited[location] = struct{}{}
					}
				}
			case west:
				for i := 0; i < dist; i++ {
					location.x--
					if !seen {
						if _, ok := visited[location]; ok {
							visitedTwice = location
							seen = true
						}
						visited[location] = struct{}{}
					}
				}
			case east:
				for i := 0; i < dist; i++ {
					location.x++
					if !seen {
						if _, ok := visited[location]; ok {
							visitedTwice = location
							seen = true
						}
						visited[location] = struct{}{}
					}
				}
			}
		}
	}
	fmt.Println(location.manhattanDistance(Coord{}))
	fmt.Println(visitedTwice.manhattanDistance(Coord{}))
}

func (c Coord) manhattanDistance(n Coord) int {
	t := math.Abs(float64(c.x)-float64(n.x)) + math.Abs(float64(c.y)-float64(n.y))
	return int(t)
}

func nextDirection(curr int, turn string) int {
	switch curr {
	case north:
		if turn == "L" {
			return west
		}
		if turn == "R" {
			return east
		}
	case south:
		if turn == "L" {
			return east
		}
		if turn == "R" {
			return west
		}
	case east:
		if turn == "L" {
			return north
		}
		if turn == "R" {
			return south
		}
	case west:
		if turn == "L" {
			return south
		}
		if turn == "R" {
			return north
		}
	}
	return curr
}

func input() string {
	return `R4, R3, R5, L3, L5, R2, L2, R5, L2, R5, R5, R5, R1, R3, L2, L2, L1, R5, L3, R1, L2, R1, L3, L5, L1, R3, L4, R2, R4, L3, L1, R4, L4, R3, L5, L3, R188, R4, L1, R48, L5, R4, R71, R3, L2, R188, L3, R2, L3, R3, L5, L1, R1, L2, L4, L2, R5, L3, R3, R3, R4, L3, L4, R5, L4, L4, R3, R4, L4, R1, L3, L1, L1, R4, R1, L4, R1, L1, L3, R2, L2, R2, L1, R5, R3, R4, L5, R2, R5, L5, R1, R2, L1, L3, R3, R1, R3, L4, R4, L4, L1, R1, L2, L2, L4, R1, L3, R4, L2, R3, L1, L5, R4, R5, R2, R5, R1, R5, R1, R3, L3, L2, L2, L5, R2, L2, R5, R5, L2, R3, L5, R5, L2, R4, R2, L1, R3, L5, R3, R2, R5, L1, R3, L2, R2, R1`
}
