package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var part1Test = struct {
	in  []string
	out int
}{
	[]string{"5 1 9 5", "7 5 3", "2 4 6 8"}, 18,
}

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		sum := part1(&part1Test.in)
		assert.Equal(t, part1Test.out, sum)
	})
}
