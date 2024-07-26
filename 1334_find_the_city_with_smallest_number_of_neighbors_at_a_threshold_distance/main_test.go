package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheCity(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
		distanceThreshold := 4
		expected := 3

		assert.Equal(t, expected, findTheCity(n, edges, distanceThreshold))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
		distanceThreshold := 2
		expected := 0

		assert.Equal(t, expected, findTheCity(n, edges, distanceThreshold))
	})
}
