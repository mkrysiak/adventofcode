package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Cart struct {
	coordinate Coordinate
	direction  Direction
	lastTurn   Turn
	id         int
}

type Carts []Cart

type Coordinate struct {
	x int
	y int
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Turn int

const (
	Left Turn = iota
	Straight
	Right
)

func (c Carts) Len() int      { return len(c) }
func (c Carts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Carts) Less(i, j int) bool {
	if c[i].coordinate.x < c[j].coordinate.x {
		return true
	}
	if c[i].coordinate.x == c[j].coordinate.x {
		return c[i].coordinate.y < c[j].coordinate.y
	}
	return false
}

func main() {
	contents := readInputFile("input")
	fmt.Printf("Part 1: %s\n", part1(contents))
	fmt.Printf("Part 2: %s\n", part2(contents))
}

func part1(contents *[]string) string {
	trackMap, carts := buildMap(contents)
	cartsMap := map[Coordinate]Cart{}
	collision := ""

	for _, c := range carts {
		cartsMap[c.coordinate] = c
	}
	// printMap(trackMap, cartsMap)
	// fmt.Println(carts)

	for collision == "" {
		sort.Sort(carts)
		var newCarts Carts
		for _, cart := range carts {
			newCart := nextMove(trackMap, cart)
			if _, ok := cartsMap[newCart.coordinate]; ok {
				// Collision
				collision = strconv.Itoa(newCart.coordinate.y) + "," + strconv.Itoa(newCart.coordinate.x)
				break
			} else {
				delete(cartsMap, cart.coordinate)
				cartsMap[newCart.coordinate] = newCart
			}
			newCarts = append(newCarts, newCart)
		}
		carts = newCarts
		// printMap(trackMap, cartsMap)
		// fmt.Println()
	}

	return collision
}

func part2(contents *[]string) string {
	trackMap, carts := buildMap(contents)
	cartsMap := map[Coordinate]Cart{}

	for _, c := range carts {
		cartsMap[c.coordinate] = c
	}
	// printMap(trackMap, cartsMap)
	// fmt.Println(carts)

	removedCarts := map[int]struct{}{}
	for len(carts) > 1 {
		sort.Sort(carts)
		for _, cart := range carts {
			if _, ok := removedCarts[cart.id]; ok {
				continue
			}
			newCart := nextMove(trackMap, cart)
			if _, ok := cartsMap[newCart.coordinate]; ok {
				// Collision
				removedCarts[cart.id] = struct{}{}
				removedCarts[cartsMap[newCart.coordinate].id] = struct{}{}
				delete(cartsMap, cart.coordinate)
				delete(cartsMap, newCart.coordinate)
			} else {
				// Update position
				delete(cartsMap, cart.coordinate)
				cartsMap[newCart.coordinate] = newCart
			}
		}
		carts = Carts{}
		for _, v := range cartsMap {
			carts = append(carts, v)
		}
		// printMap(trackMap, cartsMap)
		// fmt.Println()
	}

	return strconv.Itoa(carts[0].coordinate.y) + "," + strconv.Itoa(carts[0].coordinate.x)
}

func printMap(trackMap [][]rune, cartsMap map[Coordinate]Cart) {
	for i := range trackMap {
		for j := range trackMap[i] {
			s := fmt.Sprintf("%s", string(trackMap[i][j]))
			if d, ok := cartsMap[Coordinate{x: i, y: j}]; ok {
				if d.direction == North {
					s = "^"
				}
				if d.direction == South {
					s = "v"
				}
				if d.direction == East {
					s = ">"
				}
				if d.direction == West {
					s = "<"
				}
			}
			fmt.Printf("%s", s)
		}
		fmt.Println()
	}
}

func nextMove(trackMap [][]rune, cart Cart) Cart {
	// Cart is not at an intersection
	// Find next coordinate.  If it's a turn, change direction.
	switch cart.direction {
	case North:
		cart.coordinate.x--
		r := trackMap[cart.coordinate.x][cart.coordinate.y]
		if r == '\\' {
			cart.direction = West
		}
		if r == '/' {
			cart.direction = East
		}
		if r == '+' {
			t := nextTurn(cart.lastTurn)
			cart.lastTurn = t
			switch t {
			case Left:
				cart.direction = West
			case Right:
				cart.direction = East
			}
		}
	case South:
		cart.coordinate.x++
		r := trackMap[cart.coordinate.x][cart.coordinate.y]
		if r == '\\' {
			cart.direction = East
		}
		if r == '/' {
			cart.direction = West
		}
		if r == '+' {
			t := nextTurn(cart.lastTurn)
			cart.lastTurn = t
			switch t {
			case Left:
				cart.direction = East
			case Right:
				cart.direction = West
			}
		}
	case East:
		cart.coordinate.y++
		r := trackMap[cart.coordinate.x][cart.coordinate.y]
		if r == '\\' {
			cart.direction = South
		}
		if r == '/' {
			cart.direction = North
		}
		if r == '+' {
			t := nextTurn(cart.lastTurn)
			cart.lastTurn = t
			switch t {
			case Left:
				cart.direction = North
			case Right:
				cart.direction = South
			}
		}
	case West:
		cart.coordinate.y--
		r := trackMap[cart.coordinate.x][cart.coordinate.y]
		if r == '\\' {
			cart.direction = North
		}
		if r == '/' {
			cart.direction = South
		}
		if r == '+' {
			t := nextTurn(cart.lastTurn)
			cart.lastTurn = t
			switch t {
			case Left:
				cart.direction = South
			case Right:
				cart.direction = North
			}
		}
	}
	return cart
}

func isTurn(trackMap [][]rune, coordinate Coordinate) (rune, bool) {
	c := trackMap[coordinate.x][coordinate.y]
	if c == '\\' || c == '/' {
		return c, true
	}
	return c, false
}

func nextTurn(lastTurn Turn) Turn {
	switch lastTurn {
	case Left:
		return Straight
	case Straight:
		return Right
	}
	// case Right:
	return Left
}

func buildMap(contents *[]string) ([][]rune, Carts) {
	var trackMap = make([][]rune, len(*contents))
	for i := range trackMap {
		trackMap[i] = make([]rune, len((*contents)[0]))
	}
	carts := Carts{}
	idn := 0
	for i, line := range *contents {
		nodes := []rune(line)
		for j, r := range nodes {
			c := Coordinate{
				x: i,
				y: j,
			}
			switch string(r) {
			case "^":
				cart := Cart{
					direction:  North,
					lastTurn:   Right,
					coordinate: c,
					id:         idn,
				}
				carts = append(carts, cart)
			case "v":
				cart := Cart{
					direction:  South,
					lastTurn:   Right,
					coordinate: c,
					id:         idn,
				}
				carts = append(carts, cart)
			case ">":
				cart := Cart{
					direction:  East,
					lastTurn:   Right,
					coordinate: c,
					id:         idn,
				}
				carts = append(carts, cart)
			case "<":
				cart := Cart{
					direction:  West,
					lastTurn:   Right,
					coordinate: c,
					id:         idn,
				}
				carts = append(carts, cart)
			default:
			}
			if r == 'v' || r == '^' {
				r = '|'
			}
			if r == '<' || r == '>' {
				r = '-'
			}
			trackMap[i][j] = r
			idn++
		}
	}
	return trackMap, carts
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
