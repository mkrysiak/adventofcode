package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Character struct {
	id          int
	coord       Coordinate
	hitPoints   int
	ctype       string
	attackPower int
}

type Coordinate struct {
	x int
	y int
}

// In reading order
var neighbors = Coordinates{{x: 0, y: -1}, {x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}}

type Graph struct {
	adj map[Coordinate][]Coordinate
}

type Characters []Character
type Coordinates []Coordinate

func main() {
	contents := readInputFile("input")
	part1(contents)
	part2(contents)
}

func part1(contents *[]string) {
	m, chars := buildMap(contents, 3)
	// printMap(m, chars, nil)

	destroyedCharacters := map[int]struct{}{}
	rounds := 0

	for {
		rounds++
		// fmt.Printf("Round: %d\n", rounds)
		sort.Sort(chars)
		for i, src := range chars {
			// If no targets remain, combat ends.
			if !targetsRemaining(lessDestroyedChars(destroyedCharacters, chars)) {
				rounds--
				break
			}
			if _, ok := destroyedCharacters[src.id]; ok {
				continue
			}
			// fmt.Printf("Character: (%d,%d)\n", src.coord.x, src.coord.y)

			// After moving (or if the unit began its turn in range of a target), the unit attacks.
			ldc := lessDestroyedChars(destroyedCharacters, chars)
			// fmt.Printf("LDC: %v\n", ldc)
			we := weakestEnemy(src, ldc)
			if (we == Character{}) {
				// else move closer, and attack
				a := inRangePositions(overlayedMap(m, ldc, nil), ldc, src)
				// fmt.Printf("InRange: %v\n", a)
				cc := connectedComponents(overlayedMap(m, ldc, nil), src.coord)
				// fmt.Printf("CC: %v\n", cc)
				r := reachablePositions(a, cc)
				// fmt.Printf("Reachable: %v\n", r)
				// printMap(m, ldc, r)
				g := NewGraph(overlayedMap(m, ldc, nil), src.coord)
				chars[i].coord = g.nextMove(src.coord, r)
			}
			we = weakestEnemy(chars[i], ldc)
			if (we != Character{}) {
				// After attacking, the unit's turn ends.
				// fmt.Printf("ATTACK: (%d,%d) -> (%d,%d)\n", src.coord.x, src.coord.y, we.coord.x, we.coord.y)
				for j := range chars {
					if we.id == chars[j].id {
						chars[j].hitPoints -= src.attackPower
						if chars[j].hitPoints <= 0 {
							// fmt.Printf("DIED: (%d, %d,%d)\n", chars[j].id, chars[j].coord.x, chars[j].coord.y)
							destroyedCharacters[chars[j].id] = struct{}{}
						}
					}
				}
			}
		}
		// printMap(m, lessDestroyedChars(destroyedCharacters, chars), nil)

		chars = lessDestroyedChars(destroyedCharacters, chars)

		if !targetsRemaining(chars) {
			break
		}
	}

	hitPointsSum := 0
	for _, c := range chars {
		hitPointsSum += c.hitPoints
	}
	fmt.Printf("Part 1: Rounds: %d HP: %d Product: %d\n", rounds, hitPointsSum, hitPointsSum*rounds)
}

func part2(contents *[]string) {
	elfAttackPower := 3
	_, charss := buildMap(contents, elfAttackPower)

	elfCount := charactersRemaining(charss, "elf")
	rounds := 0
	hitPointsSum := 0

	for {
		m, chars := buildMap(contents, elfAttackPower)
		destroyedCharacters := map[int]struct{}{}
		for {
			rounds++
			// fmt.Printf("\nRound: %d\n", rounds)
			sort.Sort(chars)
			for i, src := range chars {
				// If no targets remain, combat ends.
				if !targetsRemaining(lessDestroyedChars(destroyedCharacters, chars)) {
					rounds--
					break
				}
				if _, ok := destroyedCharacters[src.id]; ok {
					continue
				}
				// fmt.Printf("Character: (%d,%d)\n", src.coord.x, src.coord.y)

				// After moving (or if the unit began its turn in range of a target), the unit attacks.
				ldc := lessDestroyedChars(destroyedCharacters, chars)
				we := weakestEnemy(src, ldc)
				if (we == Character{}) {
					// else move closer, and attack
					a := inRangePositions(overlayedMap(m, ldc, nil), ldc, src)
					cc := connectedComponents(overlayedMap(m, ldc, nil), src.coord)
					r := reachablePositions(a, cc)
					// printMap(m, lessDestroyedChars(destroyedCharacters, chars), r)
					g := NewGraph(overlayedMap(m, ldc, nil), src.coord)
					chars[i].coord = g.nextMove(src.coord, r)
				}
				we = weakestEnemy(chars[i], ldc)
				if (we != Character{}) {
					// After attacking, the unit's turn ends.
					for j := range chars {
						if we.id == chars[j].id {
							chars[j].hitPoints -= src.attackPower
							if chars[j].hitPoints <= 0 {
								destroyedCharacters[chars[j].id] = struct{}{}
							}
						}
					}
				}
			}
			// printMap(m, lessDestroyedChars(destroyedCharacters, chars), nil)
			chars = lessDestroyedChars(destroyedCharacters, chars)

			if !targetsRemaining(chars) {
				hps := 0
				for _, c := range chars {
					hps += c.hitPoints
				}
				break
			}
		}

		if elfCount == charactersRemaining(chars, "elf") {
			for _, c := range chars {
				hitPointsSum += c.hitPoints
			}
			break
		}
		rounds = 0
		elfAttackPower++
	}
	fmt.Printf("Rounds: %d HP: %d AP: %d\n", rounds, hitPointsSum, elfAttackPower)
	fmt.Printf("Part 2: %d\n", rounds*hitPointsSum)
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

func targetsRemaining(chars Characters) bool {
	goblinCount, elfCount := 0, 0
	for _, c := range chars {
		if c.ctype == "goblin" {
			goblinCount++
		}
		if c.ctype == "elf" {
			elfCount++
		}
	}
	if goblinCount == 0 || elfCount == 0 {
		return false
	}
	return true
}

func charactersRemaining(chars Characters, character string) int {
	elfCount := 0
	for _, c := range chars {
		if c.ctype == character {
			elfCount++
		}
	}
	return elfCount
}

func lessDestroyedChars(d map[int]struct{}, chars Characters) Characters {
	newChars := Characters{}
	for _, v := range chars {
		if _, ok := d[v.id]; !ok {
			newChars = append(newChars, v)
		}
	}
	return newChars
}

func printMap(m [][]rune, chars Characters, other Coordinates) {
	mm := overlayedMap(m, chars, other)

	charsByLine := map[int]string{}
	sort.Sort(chars)
	for _, v := range chars {
		if v.ctype == "elf" {
			charsByLine[v.coord.y] += fmt.Sprintf("E(%d,%d,%d), ", v.hitPoints, v.coord.x, v.coord.y)
		} else {
			charsByLine[v.coord.y] += fmt.Sprintf("G(%d,%d,%d), ", v.hitPoints, v.coord.x, v.coord.y)
		}
	}

	for y := range mm {
		fmt.Printf("%2d ", y)
		for x := range mm[y] {
			fmt.Printf("%s", string(mm[y][x]))
		}
		if _, ok := charsByLine[y]; ok {
			fmt.Printf("  %s\n", charsByLine[y])
		} else {
			fmt.Println()
		}
	}
}

func weakestEnemy(src Character, chars Characters) Character {

	var enemyType string
	if src.ctype == "elf" {
		enemyType = "goblin"
	} else {
		enemyType = "elf"
	}

	enemies := map[Coordinate]Character{}
	for _, v := range chars {
		if v.ctype == enemyType {
			enemies[v.coord] = v
		}
	}

	weakestEnemy := Character{}
	healthOfWeakestEnemy := 201
	for _, n := range neighbors {
		if e, ok := enemies[Coordinate{x: (src.coord.x + n.x), y: (src.coord.y + n.y)}]; ok {
			if e.hitPoints < healthOfWeakestEnemy {
				healthOfWeakestEnemy = e.hitPoints
				weakestEnemy = e
			}
		}
	}
	return weakestEnemy
}

func inRangePositions(m [][]rune, characters []Character, source Character) Coordinates {
	openPositions := Coordinates{}
	for _, c := range characters {
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

func overlayedMap(m [][]rune, chars Characters, other Coordinates) [][]rune {
	cMap := map[Coordinate]rune{}
	mm := [][]rune{}
	for y := range m {
		mm = append(mm, []rune{})
		for x := range m[y] {
			mm[y] = append(mm[y], m[y][x])
		}
	}
	for _, v := range chars {
		if v.ctype == "elf" {
			cMap[v.coord] = 'E'
		}
		if v.ctype == "goblin" {
			cMap[v.coord] = 'G'
		}
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

func buildMap(contents *[]string, elfAttackPower int) ([][]rune, Characters) {
	var m = make([][]rune, len(*contents))
	chars := Characters{}
	for i := range m {
		m[i] = make([]rune, len((*contents)[0]))
	}
	id := 0
	for y, line := range *contents {
		r := []rune(line)
		for x, c := range r {
			switch c {
			case 'E':
				chars = append(chars, Character{
					id:        id,
					hitPoints: 200,
					coord: Coordinate{
						x: x,
						y: y,
					},
					ctype:       "elf",
					attackPower: elfAttackPower,
				})
				c = '.'
				id++
			case 'G':
				chars = append(chars, Character{
					id:        id,
					hitPoints: 200,
					coord: Coordinate{
						x: x,
						y: y,
					},
					ctype:       "goblin",
					attackPower: 3,
				})
				c = '.'
				id++
			default:
			}
			m[y][x] = c
		}
	}
	return m, chars
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
	// the unit chooses the step which is first in reading order
	minDistance := count
	for _, pos := range reachablePositions {
		// fmt.Printf("Possible: %v %v\n", pos, paths[pos])
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
	sort.Sort(targetPositions)
	// fmt.Printf("DST: %v STEP: %v\n", targetPositions[0], paths[targetPositions[0]][len(paths[targetPositions[0]])-1])
	return paths[targetPositions[0]][len(paths[targetPositions[0]])-1]
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

func (c Characters) Len() int      { return len(c) }
func (c Characters) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Characters) Less(i, j int) bool {
	if c[i].coord.y < c[j].coord.y {
		return true
	}
	if c[i].coord.y == c[j].coord.y {
		return c[i].coord.x < c[j].coord.x
	}
	return false
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
