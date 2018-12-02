package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	in     string
	twos   bool
	threes bool
}{
	{"abcdef", false, false},
	{"bababc", true, true},
	{"abbcde", true, false},
	{"abcccd", false, true},
	{"aabcdd", true, false},
	{"abcdee", true, false},
	{"ababab", false, true},
}

func TestPart1HasTwoOrThree(t *testing.T) {
	for _, tt := range part1Tests {
		t.Run("Part 1 - hasTwosOrThrees", func(t *testing.T) {
			twos, threes := hasTwosOrThrees(tt.in)
			assert.Equal(t, tt.twos, twos)
			assert.Equal(t, tt.threes, threes)
		})
	}
}

func TestPart1(t *testing.T) {
	part1Test := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	t.Run("Part 1", func(t *testing.T) {
		checksum := part1(&part1Test)
		assert.Equal(t, 12, checksum)
	})
}

func TestPart2(t *testing.T) {
	part2Test := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	t.Run("Part 2", func(t *testing.T) {
		match := part2(&part2Test)
		assert.Equal(t, "fgij", match)
	})
}

func TestPart2Altnerative(t *testing.T) {
	part2Test := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	t.Run("Part 2 - Alternative", func(t *testing.T) {
		match := part2Alternative(&part2Test)
		assert.Equal(t, "fgij", match)
	})
}
