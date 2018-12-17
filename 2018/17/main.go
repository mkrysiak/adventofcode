package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var inputXRE = regexp.MustCompile(`^x=(\d+),\s+y=(\d+)\.\.(\d+)`)
var inputYRE = regexp.MustCompile(`^y=(\d+),\s+x=(\d+)\.\.(\d+)`)

type Grid struct {
	g    [][]rune
	xmin int
	xmax int
	ymin int
	ymax int
}

type Coordinate struct {
	x int
	y int
}

type Coordinates []Coordinate

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents))
}

func part1(contents *[]string) int {
	grid := NewGrid(contents)
	grid.fill(0, 500)
	// grid.print()

	// fmt.Printf("Xmin: %d Xmax: %d Ymin: %d Ymax: %d\n", grid.xmin, grid.xmax, grid.ymin, grid.ymax)
	waterCount := 0
	for y := grid.ymin; y <= grid.ymax; y++ {
		for x := grid.xmin - 1; x <= grid.xmax+1; x++ {
			if grid.g[y][x] == '|' || grid.g[y][x] == '~' {
				waterCount++
			}
		}
	}
	return waterCount
}

func part2(contents *[]string) int {
	grid := NewGrid(contents)

	grid.fill(0, 500)

	waterCount := 0
	for y := grid.ymin; y <= grid.ymax; y++ {
		for x := grid.xmin - 1; x <= grid.xmax+1; x++ {
			if grid.g[y][x] == '~' {
				waterCount++
			}
		}
	}
	return waterCount
}

func (g Grid) fill(y, x int) {
	// fmt.Printf("Y: %d\n", y)
	if y > g.ymax {
		return
	} else if g.notEmpty(y, x) {
		return
	}
	// If our current position is just above the bottom of a "bucket"
	if g.notEmpty(y+1, x) {
		// Go right of current position, and fill with water if
		// water cannot go down
		xRight := x + 1
		for !g.notEmpty(y, xRight) && g.notEmpty(y+1, xRight) {
			g.g[y][xRight] = '|'
			xRight++
		}
		// Go left of current position, and fill with water if
		// water cannot go down
		xLeft := x
		for !g.notEmpty(y, xLeft) && g.notEmpty(y+1, xLeft) {
			g.g[y][xLeft] = '|'
			xLeft--
		}

		if !g.notEmpty(y+1, xLeft) || !g.notEmpty(y+1, xRight) {
			// If the water has spilt over both sides, start a new branch
			// on each
			g.fill(y, xLeft)
			g.fill(y, xRight)

		} else if g.g[y][xRight] == '#' && g.g[y][xLeft] == '#' {
			// Otherwise, we're still within the walls of a bucket, and
			// convert to standing water ('|' -> '~')
			for xx := xLeft + 1; xx < xRight; xx++ {
				g.g[y][xx] = '~'
			}
		}
	} else if g.g[y][x] == '.' {
		g.g[y][x] = '|'
		g.fill(y+1, x)
		if g.g[y+1][x] == '~' {
			g.fill(y, x)
		}
	}
}

func (g Grid) notEmpty(y, x int) bool {
	return g.g[y][x] == '#' || g.g[y][x] == '~'
}

func NewGrid(contents *[]string) Grid {
	ymin, ymax, xmin, xmax := 100000, 0, 1000000, 0
	xlines, ylines := [][3]int{}, [][3]int{}
	for _, line := range *contents {
		m := inputXRE.FindStringSubmatch(line)
		if m != nil {
			x, _ := strconv.Atoi(m[1])
			yStart, _ := strconv.Atoi(m[2])
			yEnd, _ := strconv.Atoi(m[3])
			xlines = append(xlines, [3]int{x, yStart, yEnd})
			if yStart < ymin {
				ymin = yStart
			}
			if yEnd > ymax {
				ymax = yEnd
			}
		}
		m = inputYRE.FindStringSubmatch(line)
		if m != nil {
			y, _ := strconv.Atoi(m[1])
			xStart, _ := strconv.Atoi(m[2])
			xEnd, _ := strconv.Atoi(m[3])
			ylines = append(ylines, [3]int{y, xStart, xEnd})
			if xStart < xmin {
				xmin = xStart
			}
			if xEnd > xmax {
				xmax = xEnd
			}
		}
	}
	var g = make([][]rune, ymax+2)
	for i := range g {
		g[i] = make([]rune, xmax+2)
		for j := range g[i] {
			g[i][j] = '.'
		}
	}
	for _, v := range xlines {
		for j := v[1]; j <= v[2]; j++ {
			g[j][v[0]] = '#'
		}
	}
	for _, v := range ylines {
		for j := v[1]; j <= v[2]; j++ {
			g[v[0]][j] = '#'
		}
	}

	return Grid{
		g:    g,
		xmin: xmin,
		xmax: xmax,
		ymin: ymin,
		ymax: ymax,
	}
}

func (g Grid) print() {
	fmt.Printf("Xmin: %d Xmax: %d\n", g.xmin-10, g.xmax)
	for y := range g.g {
		fmt.Printf("%5d ", y)
		for x := g.xmin - 10; x <= g.xmax; x++ {
			fmt.Printf("%s", string(g.g[y][x]))
		}
		fmt.Println()
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
