package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	t.Run("Test Checksum 1", func(t *testing.T) {
		assert.Equal(t, "abxyz", checksum("aaaaa-bbb-z-y-x"))
	})
}

func TestChecksum2(t *testing.T) {
	t.Run("Test Checksum 2", func(t *testing.T) {
		assert.Equal(t, "abcde", checksum("a-b-c-d-e-f-g-h"))
	})
}

func TestChecksum3(t *testing.T) {
	t.Run("Test Checksum 3", func(t *testing.T) {
		assert.Equal(t, "oarel", checksum("not-a-real-room"))
	})
}

func TestShift(t *testing.T) {
	t.Run("Test Shift", func(t *testing.T) {
		assert.Equal(t, "very encrypted name", shift("qzmt-zixmtkozy-ivhz", 343))
	})
}
