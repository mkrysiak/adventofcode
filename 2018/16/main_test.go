package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddr(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("addr", func(t *testing.T) {
		addr([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 3, 1}, registers)
	})
}
func TestAddi(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("addi", func(t *testing.T) {
		addi([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 2, 1}, registers)
	})
}

func TestMulr(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("mulr", func(t *testing.T) {
		mulr([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 2, 1}, registers)
	})
}

func TestMuli(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("muli", func(t *testing.T) {
		muli([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 1, 1}, registers)
	})
}

func TestSeti(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("seti", func(t *testing.T) {
		seti([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 2, 1}, registers)
	})
}

func TestGtir(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("gtir", func(t *testing.T) {
		gtir([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 0, 1}, registers)
	})
}

func TestGtirTruth(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("gtir truth", func(t *testing.T) {
		gtir([4]int{9, 3, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 1, 1}, registers)
	})
}

func TestGtri(t *testing.T) {
	registers := [4]int{3, 2, 1, 1}
	t.Run("gtri", func(t *testing.T) {
		gtri([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 0, 1}, registers)
	})
}
func TestGtriTruthy(t *testing.T) {
	registers := [4]int{3, 2, 2, 1}
	t.Run("gtri truthy", func(t *testing.T) {
		gtri([4]int{9, 2, 1, 2}, &registers)
		assert.Equal(t, [4]int{3, 2, 1, 1}, registers)
	})
}
