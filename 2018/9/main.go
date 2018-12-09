package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func main() {
	playerCount, _ := strconv.Atoi(os.Args[1])
	lastMarble, _ := strconv.Atoi(os.Args[2])
	fmt.Printf("Part 1: %d\n", part1(playerCount, lastMarble))
	fmt.Printf("Part 2: %d\n", part1(playerCount, 100*lastMarble))
}

func part1(playerCount int, lastMarble int) int {
	var players = make([]int, playerCount)
	r := ring.New(1)
	r.Value = 0
	// head := r

	currPlayer := 0
	for turn := 1; turn <= lastMarble; turn++ {
		// head.Do(func(p interface{}) {
		// 	fmt.Printf("%d ", p.(int))
		// })
		// fmt.Println()

		if (turn % 23) == 0 {
			r = r.Move(-8)
			players[currPlayer] += turn
			players[currPlayer] += r.Value.(int)

			// ring.Unlink() removes r.Next(), not r.  So move back one more.
			r = r.Prev()
			r.Unlink(1)
			r = r.Move(2)
		} else {
			newMarble := ring.New(1)
			newMarble.Value = turn

			r.Link(newMarble)
			r = r.Move(2)
		}
		currPlayer++
		if currPlayer > (playerCount - 1) {
			currPlayer = 0
		}
	}

	highestScore := 0
	for _, v := range players {
		if v > highestScore {
			highestScore = v
		}
	}
	return highestScore
}
