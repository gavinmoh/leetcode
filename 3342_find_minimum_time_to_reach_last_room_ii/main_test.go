package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinTimeToReach(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		moveTime := [][]int{{0, 4}, {4, 4}}
		expected := 7

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})

	t.Run("test case 2", func(t *testing.T) {
		moveTime := [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}}
		expected := 6

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})

	t.Run("test case 3", func(t *testing.T) {
		moveTime := [][]int{{0, 1}, {1, 2}}
		expected := 4

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})
}
