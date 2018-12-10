package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Pair struct {
	positionX int
	positionY int
	velocityX int
	velocityY int
}

type Pairs []Pair

var lineRE = regexp.MustCompile(`position=<\s?(-?[0-9]+),\s+(-?\d+)>\s+velocity=<\s?(-?\d+),\s+(-?\d+)>`)

func main() {
	contents := readInputFile("input")
	pairs, cycles := part1(contents)
	fmt.Println("Part 1:")
	printPairs(pairs)
	fmt.Printf("Part 2: %d\n", cycles)
}

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	if p[i].positionX < p[j].positionX {
		return true
	}
	if p[i].positionX == p[j].positionX {
		if p[i].positionY < p[j].positionY {
			return true
		}
	}
	return false
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func printPairs(p Pairs) {
	sort.Sort(p)
	minX := p[0].positionX
	maxX := p[len(p)-1].positionX

	minY := 999999999
	maxY := 0
	for _, v := range p {
		if v.positionY < minY {
			minY = v.positionY
		}
		if v.positionY > maxY {
			maxY = v.positionY
		}
	}

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			match := false
			for _, v := range p {
				if v.positionY == i && v.positionX == j {
					match = true
					break
				}
			}
			if match {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func part1(contents *[]string) (Pairs, int) {

	pairs := Pairs{}

	uniqY := map[int]struct{}{}
	for _, v := range *contents {
		m := lineRE.FindStringSubmatch(v)
		if m != nil {
			posX, _ := strconv.Atoi(m[1])
			posY, _ := strconv.Atoi(m[2])
			velX, _ := strconv.Atoi(m[3])
			velY, _ := strconv.Atoi(m[4])

			pairs = append(pairs, Pair{
				positionX: posX,
				positionY: posY,
				velocityX: velX,
				velocityY: velY,
			})
			uniqY[posY] = struct{}{}
		}
	}

	uniqYLen := len(uniqY)
	cycles := 0
	sort.Sort(pairs)
	for {
		u := map[int]struct{}{}
		for i := range pairs {
			pairs[i].positionX += pairs[i].velocityX
			pairs[i].positionY += pairs[i].velocityY
			u[pairs[i].positionY] = struct{}{}
		}
		cycles++
		if float64(uniqYLen)*.12 > float64(len(u)) {
			return pairs, cycles
		}
	}
}

func area(pairs Pairs) int {
	return (pairs[len(pairs)-1].positionX - pairs[0].positionX) * (pairs[len(pairs)-1].positionY - pairs[0].positionY)
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
