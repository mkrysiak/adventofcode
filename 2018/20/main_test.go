package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := "^WNE$"
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, 3, part1(adjMap(input)))
	})
}

func Test2(t *testing.T) {
	input := "^ENWWW(NEEE|SSE(EE|N))$"
	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, 10, part1(adjMap(input)))
	})
}

func Test3(t *testing.T) {
	input := "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"
	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, 31, part1(adjMap(input)))
	})
}
