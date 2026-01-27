package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}
		expected := 5

		assert.Equal(t, expected, minCost(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}}
		expected := 3

		assert.Equal(t, expected, minCost(n, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 5
		edges := [][]int{{3, 4, 18}, {2, 1, 25}, {4, 3, 21}, {1, 2, 20}, {2, 3, 18}, {3, 1, 13}, {4, 2, 9}}
		expected := -1

		assert.Equal(t, expected, minCost(n, edges))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 5
		edges := [][]int{{1, 2, 8}, {3, 1, 10}, {0, 3, 3}}
		expected := -1

		assert.Equal(t, expected, minCost(n, edges))
	})
}
