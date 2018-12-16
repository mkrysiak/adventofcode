package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinateSort(t *testing.T) {
	// Character: (13,26)
	// Possible: [{13 25} {13 25} {12 26}]
	c := Coordinates{
		{13, 25},
		{12, 26},
	}
	t.Run("Testing Coordinate Sort", func(t *testing.T) {
		sort.Sort(c)
		assert.Equal(t, Coordinate{13, 25}, c[0])
	})
}
