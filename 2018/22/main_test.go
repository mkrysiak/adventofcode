package main

import (
	"testing"

	"github.com/mkrysiak/adventofcode/2018/22/lib"
	"github.com/stretchr/testify/assert"
)

var depth = 510
var target = lib.Coordinate{10, 10}

func TestGeoIndex1(t *testing.T) {
	c := lib.Coordinate{0, 0}
	fillCavesErosionLevel(depth, target)
	t.Run("Test Geo Index 1", func(t *testing.T) {
		assert.Equal(t, 0, geologicIndex(depth, c, target))
	})
	t.Run("Test Geo Index 1", func(t *testing.T) {
		assert.Equal(t, rocky, regionType(c))
	})
}

func TestGeoIndex2(t *testing.T) {
	c := lib.Coordinate{1, 0}
	fillCavesErosionLevel(depth, target)
	t.Run("Test Geo Index 2", func(t *testing.T) {
		assert.Equal(t, 16807, geologicIndex(depth, c, target))
	})
	t.Run("Test Geo Index 2 - Region Type", func(t *testing.T) {
		assert.Equal(t, wet, regionType(c))
	})
}
func TestGeoIndex3(t *testing.T) {
	c := lib.Coordinate{0, 1}
	fillCavesErosionLevel(depth, target)
	t.Run("Test Geo Index 3", func(t *testing.T) {
		assert.Equal(t, 48271, geologicIndex(depth, c, target))
	})
	t.Run("Test Geo Index 3 - Region Type", func(t *testing.T) {
		assert.Equal(t, rocky, regionType(c))
	})
}
func TestGeoIndex4(t *testing.T) {
	c := lib.Coordinate{1, 1}
	fillCavesErosionLevel(depth, target)
	t.Run("Test Geo Index 4", func(t *testing.T) {
		assert.Equal(t, 145722555, geologicIndex(depth, c, target))
	})
	t.Run("Test Geo Index 5 - Region Type", func(t *testing.T) {
		assert.Equal(t, narrow, regionType(c))
	})
}
func TestGeoIndex5(t *testing.T) {
	c := lib.Coordinate{10, 10}
	fillCavesErosionLevel(depth, target)
	t.Run("Test Geo Index 5", func(t *testing.T) {
		assert.Equal(t, 0, geologicIndex(depth, c, target))
	})
	t.Run("Test Geo Index 5 - Region Type", func(t *testing.T) {
		assert.Equal(t, rocky, regionType(c))
	})
}
