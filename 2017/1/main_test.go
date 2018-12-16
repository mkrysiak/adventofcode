package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	in  [4]int
	sum int
}{
	{[4]int{1, 1, 2, 2}, 3},
	{[4]int{1, 1, 1, 1}, 4},
	{[4]int{1, 2, 3, 4}, 0},
	{[4]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		t.Run("Part 1", func(t *testing.T) {
			sum := part1(tt.in)
			assert.Equal(t, tt.sum, sum)
		})
	}
}

var part2Tests = []struct {
	in  [4]int
	sum int
}{
	{[4]int{1, 2, 1, 2}, 6},
	{[4]int{1, 2, 2, 1}, 0},
	{[4]int{1, 2, 3, 4, 2, 5}, 4},
	{[4]int{1, 2, 3, 1, 2, 3}, 12},
	{[4]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		t.Run("Part 2", func(t *testing.T) {
			sum := part2(tt.in)
			assert.Equal(t, tt.sum, sum)
		})
	}
}
