package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTilePossibilities(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tiles := "AAB"
		expected := 8

		assert.Equal(t, expected, numTilePossibilities(tiles))
	})

	t.Run("test case 2", func(t *testing.T) {
		tiles := "AAABBC"
		expected := 188

		assert.Equal(t, expected, numTilePossibilities(tiles))
	})

	t.Run("test case 3", func(t *testing.T) {
		tiles := "V"
		expected := 1

		assert.Equal(t, expected, numTilePossibilities(tiles))
	})
}
