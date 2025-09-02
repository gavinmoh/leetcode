package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		points := [][]int{{1, 1}, {2, 2}, {3, 3}}
		expected := 0

		assert.Equal(t, expected, numberOfPairs(points))
	})

	t.Run("test case 2", func(t *testing.T) {
		points := [][]int{{6, 2}, {4, 4}, {2, 6}}
		expected := 2

		assert.Equal(t, expected, numberOfPairs(points))
	})

	t.Run("test case 3", func(t *testing.T) {
		points := [][]int{{3, 1}, {1, 3}, {1, 1}}
		expected := 2

		assert.Equal(t, expected, numberOfPairs(points))
	})

	t.Run("test case 4", func(t *testing.T) {
		points := [][]int{{0, 0}, {0, 3}}
		expected := 1

		assert.Equal(t, expected, numberOfPairs(points))
	})
}
