package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinTimeToReach(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		moveTime := [][]int{{0, 4}, {4, 4}}
		expected := 6

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})

	t.Run("test case 2", func(t *testing.T) {
		moveTime := [][]int{{0, 0, 0}, {0, 0, 0}}
		expected := 3

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})

	t.Run("test case 3", func(t *testing.T) {
		moveTime := [][]int{{0, 1}, {1, 2}}
		expected := 3

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})

	t.Run("test case 4", func(t *testing.T) {
		moveTime := [][]int{{94, 79, 62, 27, 69, 84}, {6, 32, 11, 82, 42, 30}}
		expected := 72

		assert.Equal(t, expected, minTimeToReach(moveTime))
	})
}
