package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStones(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
		expected := 5

		assert.Equal(t, expected, removeStones(stones))
	})

	t.Run("test case 2", func(t *testing.T) {
		stones := [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}
		expected := 3

		assert.Equal(t, expected, removeStones(stones))
	})

	t.Run("test case 3", func(t *testing.T) {
		stones := [][]int{{0, 0}}
		expected := 0

		assert.Equal(t, expected, removeStones(stones))
	})
}
