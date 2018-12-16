package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Unit struct {
	id          int
	coord       Coordinate
	hitPoints   int
	ctype       rune
	attackPower int
	alive       bool
}

type Coordinate struct {
	x int
	y int
}

// In reading order
var neighbors = Coordinates{{x: 0, y: -1}, {x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}}

type Units []Unit
type Coordinates []Coordinate

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %d\n", part1(contents))
	fmt.Printf("Part 2: %d\n", part2(contents, 3))
}

func part1(contents *[]string) int {
	m, units := buildMap(contents, 3)
	// printMap(m, units, nil)

	rounds := 0

	for {
		rounds++
		sort.Sort(units)
		for i, src := range units {
			// If no targets remain, combat ends.
			aliveunits := aliveUnits(units)
			if !targetsRemaining(aliveunits) {
				rounds--
				break
			}
			if !src.alive {
				continue
			}
			// Move, if possible
			om := overlayedMap(m, aliveunits, nil)
			we, err := weakestEnemy(src, aliveunits)
			if err != nil {
				r := reachablePositions(inRangePositions(om, aliveunits, src), connectedComponents(om, src.coord))
				g := NewGraph(om, src.coord)
				units[i].coord = g.nextMove(src.coord, r)
			}

			// After moving (or if the unit began its turn in range of a target), the unit attacks.
			we, err = weakestEnemy(units[i], aliveunits)
			if err == nil {
				// After attacking, the unit's turn ends.
				// fmt.Printf("ATTACK: (%d,%d) -> (%d,%d)\n", src.coord.x, src.coord.y, we.coord.x, we.coord.y)
				for j := range units {
					if we.id == units[j].id {
						units[j].hitPoints -= src.attackPower
						if units[j].hitPoints <= 0 {
							// fmt.Printf("DIED: (%d, %d,%d)\n", units[j].id, units[j].coord.x, units[j].coord.y)
							units[j].alive = false
						}
					}
				}
			}
		}
		// printMap(m, aliveUnits(units), nil)

		units = aliveUnits(units)

		if !targetsRemaining(units) {
			break
		}
	}

	hpSum := 0
	for _, c := range units {
		hpSum += c.hitPoints
	}
	// fmt.Printf("Part 1: Rounds: %d HP: %d Product: %d\n", rounds, hpSum, hpSum*rounds)
	return hpSum * rounds
}

func part2(contents *[]string, elfAttackPower int) int {
	_, unitss := buildMap(contents, elfAttackPower)

	elfCount := unitsRemaining(unitss, 'E')
	rounds := 0
	hpSum := 0

	for {
		m, units := buildMap(contents, elfAttackPower)
		for {
			rounds++
			sort.Sort(units)
			for i, src := range units {
				// If no targets remain, combat ends.
				aliveunits := aliveUnits(units)
				if !targetsRemaining(aliveunits) {
					rounds--
					break
				}
				if !src.alive {
					continue
				}
				// Move, if possible
				om := overlayedMap(m, aliveunits, nil)
				we, err := weakestEnemy(src, aliveunits)
				if err != nil {
					r := reachablePositions(inRangePositions(om, aliveunits, src), connectedComponents(om, src.coord))
					g := NewGraph(om, src.coord)
					units[i].coord = g.nextMove(src.coord, r)
				}

				// After moving (or if the unit began its turn in range of a target), the unit attacks.
				we, err = weakestEnemy(units[i], aliveunits)
				if err == nil {
					// After attacking, the unit's turn ends.
					for j := range units {
						if we.id == units[j].id {
							units[j].hitPoints -= src.attackPower
							if units[j].hitPoints <= 0 {
								units[j].alive = false
							}
						}
					}
				}
			}
			// printMap(m, aliveUnits(units), nil)
			units = aliveUnits(units)

			if !targetsRemaining(units) {
				hps := 0
				for _, c := range units {
					hps += c.hitPoints
				}
				break
			}
		}

		if elfCount == unitsRemaining(units, 'E') {
			for _, c := range units {
				hpSum += c.hitPoints
			}
			break
		}
		rounds = 0
		elfAttackPower++
	}
	// fmt.Printf("Rounds: %d HP: %d AP: %d\n", rounds, hpSum, elfAttackPower)
	return rounds * hpSum
}

func targetsRemaining(units Units) bool {
	goblinCount, elfCount := 0, 0
	for _, c := range units {
		if c.ctype == 'G' {
			goblinCount++
		}
		if c.ctype == 'E' {
			elfCount++
		}
	}
	if goblinCount == 0 || elfCount == 0 {
		return false
	}
	return true
}

func unitsRemaining(units Units, unit rune) int {
	elfCount := 0
	for _, c := range units {
		if c.ctype == unit {
			elfCount++
		}
	}
	return elfCount
}

func aliveUnits(units Units) Units {
	u := Units{}
	for _, v := range units {
		if v.alive {
			u = append(u, v)
		}
	}
	return u
}

func printMap(m [][]rune, units Units, other Coordinates) {
	mm := overlayedMap(m, units, other)

	unitsByLine := map[int]string{}
	sort.Sort(units)
	for _, v := range units {
		if v.ctype == 'E' {
			unitsByLine[v.coord.y] += fmt.Sprintf("E(%d,%d,%d), ", v.hitPoints, v.coord.x, v.coord.y)
		} else {
			unitsByLine[v.coord.y] += fmt.Sprintf("G(%d,%d,%d), ", v.hitPoints, v.coord.x, v.coord.y)
		}
	}

	for y := range mm {
		fmt.Printf("%2d ", y)
		for x := range mm[y] {
			fmt.Printf("%s", string(mm[y][x]))
		}
		if _, ok := unitsByLine[y]; ok {
			fmt.Printf("  %s\n", unitsByLine[y])
		} else {
			fmt.Println()
		}
	}
}

func weakestEnemy(src Unit, units Units) (Unit, error) {

	var enemyType rune
	if src.ctype == 'E' {
		enemyType = 'G'
	} else {
		enemyType = 'E'
	}

	enemies := map[Coordinate]Unit{}
	for _, v := range units {
		if v.ctype == enemyType {
			enemies[v.coord] = v
		}
	}

	weakestEnemy := Unit{}
	healthOfWeakestEnemy := 201
	for _, n := range neighbors {
		if e, ok := enemies[Coordinate{x: (src.coord.x + n.x), y: (src.coord.y + n.y)}]; ok {
			if e.hitPoints < healthOfWeakestEnemy {
				healthOfWeakestEnemy = e.hitPoints
				weakestEnemy = e
			}
		}
	}
	if (weakestEnemy == Unit{}) {
		return weakestEnemy, errors.New("No adjacent enemy found")
	}
	return weakestEnemy, nil
}

func inRangePositions(m [][]rune, units Units, source Unit) Coordinates {
	openPositions := Coordinates{}
	for _, c := range units {
		if c.ctype != source.ctype {
			for _, v := range neighbors {
				if m[c.coord.y+v.y][c.coord.x+v.x] == '.' {
					openPositions = append(openPositions, Coordinate{x: c.coord.x + v.x, y: c.coord.y + v.y})
				}
			}
		}
	}
	return openPositions
}

func reachablePositions(inRange Coordinates, cc Coordinates) Coordinates {
	coordinates := Coordinates{}
	// Union
	ccMap := map[Coordinate]struct{}{}
	inRangeMap := map[Coordinate]struct{}{}
	for _, v := range cc {
		ccMap[v] = struct{}{}
	}
	for _, v := range inRange {
		inRangeMap[v] = struct{}{}
	}
	for k := range inRangeMap {
		if _, ok := ccMap[k]; ok {
			coordinates = append(coordinates, k)
		}
	}
	return coordinates
}

func overlayedMap(m [][]rune, units Units, other Coordinates) [][]rune {
	cMap := map[Coordinate]rune{}
	mm := [][]rune{}
	for y := range m {
		mm = append(mm, []rune{})
		for x := range m[y] {
			mm[y] = append(mm[y], m[y][x])
		}
	}
	for _, v := range units {
		cMap[v.coord] = v.ctype
	}
	if other != nil {
		for _, v := range other {
			cMap[v] = '!'
		}
	}
	for y := range mm {
		for x := range mm[y] {
			if v, ok := cMap[Coordinate{x: x, y: y}]; ok {
				mm[y][x] = v
			}
		}
	}
	return mm
}

func connectedComponents(m [][]rune, coord Coordinate) Coordinates {
	ccMap := map[Coordinate]struct{}{
		coord: struct{}{},
	}
	q := []Coordinate{coord}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		for _, v := range neighbors {
			if m[c.y+v.y][c.x+v.x] == '.' {
				cc := Coordinate{x: c.x + v.x, y: c.y + v.y}
				if _, ok := ccMap[cc]; !ok {
					q = append(q, cc)
					ccMap[cc] = struct{}{}
				}
			}
		}
	}
	delete(ccMap, coord)
	ccl := Coordinates{}
	for k := range ccMap {
		ccl = append(ccl, k)
	}
	return ccl
}

func buildMap(contents *[]string, elfAttackPower int) ([][]rune, Units) {
	var m = make([][]rune, len(*contents))
	units := Units{}
	for i := range m {
		m[i] = make([]rune, len((*contents)[0]))
	}
	id := 0
	for y, line := range *contents {
		r := []rune(line)
		for x, c := range r {
			switch c {
			case 'E', 'G':
				attackPower := 3
				if c == 'E' {
					attackPower = elfAttackPower
				}
				units = append(units, Unit{
					id:        id,
					hitPoints: 200,
					coord: Coordinate{
						x: x,
						y: y,
					},
					ctype:       c,
					attackPower: attackPower,
					alive:       true,
				})
				c = '.'
				id++
			default:
			}
			m[y][x] = c
		}
	}
	return m, units
}

func (c Coordinates) Len() int      { return len(c) }
func (c Coordinates) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Coordinates) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	}
	if c[i].y == c[j].y {
		return c[i].x < c[j].x
	}
	return false
}

func (u Units) Len() int      { return len(u) }
func (u Units) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u Units) Less(i, j int) bool {
	if u[i].coord.y < u[j].coord.y {
		return true
	}
	if u[i].coord.y == u[j].coord.y {
		return u[i].coord.x < u[j].coord.x
	}
	return false
}

type Graph struct {
	adj map[Coordinate][]Coordinate
}

func NewGraph(m [][]rune, source Coordinate) Graph {
	graph := Graph{
		adj: map[Coordinate][]Coordinate{},
	}
	for i := range graph.adj {
		graph.adj[i] = []Coordinate{}
	}
	seen := map[Coordinate]struct{}{
		source: struct{}{},
	}
	q := []Coordinate{source}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		for _, v := range neighbors {
			if m[c.y+v.y][c.x+v.x] == '.' {
				cc := Coordinate{x: c.x + v.x, y: c.y + v.y}
				if _, ok := seen[cc]; !ok {
					q = append(q, cc)
					seen[cc] = struct{}{}
					graph.adj[c] = append(graph.adj[c], cc)
					graph.adj[cc] = append(graph.adj[cc], c)
				}
			}
		}
	}
	return graph
}

func (g Graph) nextMove(src Coordinate, reachablePositions Coordinates) Coordinate {
	// Find shortest paths
	if len(reachablePositions) == 0 {
		return src
	}
	paths := map[Coordinate]Coordinates{}
	count := 0
	// Find the path for each reachable destination
	for _, pos := range reachablePositions {
		q := Coordinates{
			src,
		}
		seen := map[Coordinate]struct{}{
			src: struct{}{},
		}
		edgeTo := map[Coordinate]Coordinate{}
		for len(q) > 0 {
			c := q[0]
			q = q[1:]
			for _, v := range g.adj[c] {
				if _, ok := seen[v]; !ok {
					edgeTo[v] = c
					seen[v] = struct{}{}
					q = append(q, v)
				}
			}
		}
		nextHop := pos
		for nextHop != src {
			paths[pos] = append(paths[pos], nextHop)
			nextHop = edgeTo[nextHop]
			count++
		}
	}
	// First, find all paths that are shortest
	minDistance := count
	for _, pos := range reachablePositions {
		if len(paths[pos]) < minDistance {
			minDistance = len(paths[pos])
		}
	}
	targetPositions := Coordinates{}
	for _, pos := range reachablePositions {
		if len(paths[pos]) == minDistance {
			targetPositions = append(targetPositions, pos)
		}
	}

	// the unit chooses the step which is first in reading order
	sort.Sort(targetPositions)
	// fmt.Printf("DST: %v STEP: %v\n", targetPositions[0], paths[targetPositions[0]][len(paths[targetPositions[0]])-1])
	return paths[targetPositions[0]][len(paths[targetPositions[0]])-1]
}

func (g Graph) print() {
	for k, v := range g.adj {
		fmt.Printf("Graph (%d,%d): ", k.x, k.y)
		for _, vv := range v {
			fmt.Printf("(%d,%d) ", vv.x, vv.y)
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
