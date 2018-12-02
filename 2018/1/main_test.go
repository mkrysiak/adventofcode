package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	in  []string
	out int
}{
	{[]string{"+1", "+1", "+1"}, 3},
	{[]string{"+1", "+1", "-2"}, 0},
	{[]string{"-1", "-2", "-3"}, -6},
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		t.Run("Part 1 - Frequency Test", func(t *testing.T) {
			freq := part1(&tt.in)
			assert.Equal(t, tt.out, freq)
		})
	}
}

var part2Tests = []struct {
	in  []string
	out int
}{
	{[]string{"+1", "-1"}, 0},
	{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
	{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
	{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		t.Run("Part 2 - Duplicate Test", func(t *testing.T) {
			freq := part2(&tt.in)
			assert.Equal(t, tt.out, freq)
		})
	}
}
