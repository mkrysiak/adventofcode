package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := `0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0`
	t.Run("Test 1", func(t *testing.T) {
		assert.Equal(t, 2, part1(input))
	})
}

func Test2(t *testing.T) {
	input := `-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`
	t.Run("Test 2", func(t *testing.T) {
		assert.Equal(t, 4, part1(input))
	})
}

func Test3(t *testing.T) {
	input := `1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2`
	t.Run("Test 3", func(t *testing.T) {
		assert.Equal(t, 8, part1(input))
	})
}

func Test4(t *testing.T) {
	input := `1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2`
	t.Run("Test 4", func(t *testing.T) {
		assert.Equal(t, 3, part1(input))
	})
}
