package main

import (
	"container/heap"
	"fmt"

	"github.com/mkrysiak/adventofcode/2018/22/lib"
)

type rtype int

const (
	rocky rtype = iota
	wet
	narrow
)

type tool int

// Align the region type index with the invalid tool index
// for simplified logic
const (
	neither tool = iota // rocky
	torch               // wet
	gear                // narrow
)

var neighbors = lib.Coordinates{{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}}

var caves = [100][12000]rune{}
var cavesErosionLevel = [100][12000]int{}

// INPUT:
// depth: 11109
// target: 9,731

func main() {
	depth := 11109
	target := lib.Coordinate{9, 731}
	fillCavesErosionLevel(depth, target)
	fmt.Printf("Part 1: %d\n", part1(target))
	fmt.Printf("Part 2: %d\n", part2(target))
}

func part1(target lib.Coordinate) int {
	return totalRiskLevel(target)
}

// Dijkstra's
func part2(target lib.Coordinate) int {
	distTo := map[lib.Item]int{}
	marked := map[lib.CoordinateWithTool]struct{}{}

	pq := lib.PriorityQueue{}
	heap.Init(&pq)

	item := &lib.Item{
		Coord:    lib.Coordinate{0, 0},
		Tool:     int(torch),
		Priority: 0, // Minutes
	}
	heap.Push(&pq, item)

	marked[lib.CoordinateWithTool{lib.Coordinate{0, 0}, int(torch)}] = struct{}{}

	for pq.Len() > 0 {
		item = heap.Pop(&pq).(*lib.Item)
		// fmt.Printf("Pop: %v %d %d\n", item.Coord, item.Priority, item.Tool)
		if _, ok := distTo[*item]; ok {
			if distTo[*item] <= item.Priority {
				continue
			}
		}
		distTo[*item] = item.Priority

		// Return early if the target is reached
		if item.Coord == target {
			// The target is always in a rocky region, and a torch is needed to find him.
			if item.Tool == int(torch) {
				return item.Priority
			}
		}
		// Add a node the priority queue for each valid tool at the current coordinate
		for t := 0; t < 3; t++ {
			if tool(t) != tool(item.Tool) && tool(t) != tool(regionType(item.Coord)) {
				newItem := &lib.Item{
					Coord:    item.Coord,
					Tool:     t,
					Priority: item.Priority + 7,
				}
				// fmt.Printf("Push: %v %d %d\n", newItem.Coord, newItem.Priority, newItem.Tool)
				heap.Push(&pq, newItem)
			}
		}

		for _, n := range neighbors {
			next := lib.Coordinate{item.Coord.X + n.X, item.Coord.Y + n.Y}
			nextWithTool := lib.CoordinateWithTool{next, item.Tool}
			if next.X < 0 || next.Y < 0 || next.X > len(caves)-1 || next.Y > len(caves[next.X])-1 {
				continue
			}
			// The current tool is invalid tool for the next region
			if int(regionType(next)) == int(item.Tool) {
				continue
			}
			if _, ok := marked[nextWithTool]; ok {
				continue
			}
			newItem := &lib.Item{
				Coord:    next,
				Tool:     int(item.Tool),
				Priority: item.Priority + 1,
			}
			// fmt.Printf("Push: %v %d %d\n", newItem.Coord, newItem.Priority, newItem.Tool)
			marked[nextWithTool] = struct{}{}
			heap.Push(&pq, newItem)
		}

	}

	return 0
}

func totalRiskLevel(target lib.Coordinate) int {
	total := 0
	for x := 0; x <= target.X; x++ {
		for y := 0; y <= target.Y; y++ {
			total += int(regionType(lib.Coordinate{x, y}))
		}
	}
	return total
}

func fillCavesErosionLevel(depth int, target lib.Coordinate) {
	for x := 0; x < len(cavesErosionLevel); x++ {
		for y := 0; y < len(cavesErosionLevel[x]); y++ {
			cavesErosionLevel[x][y] = erosionLevel(depth, lib.Coordinate{x, y}, target)
		}
	}
}

func regionType(c lib.Coordinate) rtype {
	switch cavesErosionLevel[c.X][c.Y] % 3 {
	case 0:
		return rocky
	case 1:
		return wet
	}
	return narrow
}

func erosionLevel(depth int, c, target lib.Coordinate) int {
	return (geologicIndex(depth, c, target) + depth) % 20183
}

func geologicIndex(depth int, c, target lib.Coordinate) int {
	if c == (lib.Coordinate{0, 0}) {
		return 0
	}
	if c == target {
		return 0
	}
	if c.Y == 0 {
		return c.X * 16807
	}
	if c.X == 0 {
		return c.Y * 48271
	}
	return cavesErosionLevel[c.X-1][c.Y] * cavesErosionLevel[c.X][c.Y-1]
}
