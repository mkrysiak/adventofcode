package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var part1Tests = []string{
	"#1 @ 1,3: 4x4",
	"#2 @ 3,1: 4x4",
	"#3 @ 5,5: 2x2",
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		sum := part1(&part1Tests)
		assert.Equal(t, 4, sum)
	})
}

func TestPart1Alt(t *testing.T) {
	t.Run("Part 1 Alt", func(t *testing.T) {
		sum := part1Alt(&part1Tests)
		assert.Equal(t, 4, sum)
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		id := part2(&part1Tests)
		assert.Equal(t, 3, id)
	})
}

func BenchmarkPart1(b *testing.B) {
	// run the Fib function b.N times
	contents := readInputFile("input")
	for n := 0; n < b.N; n++ {
		part1(contents)
	}
}

func BenchmarkPart1Alt(b *testing.B) {
	// run the Fib function b.N times
	contents := readInputFile("input")
	for n := 0; n < b.N; n++ {
		part1Alt(contents)
	}
}
