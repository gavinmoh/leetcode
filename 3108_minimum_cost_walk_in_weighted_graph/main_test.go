package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumCost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		edges := [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}
		query := [][]int{{0, 3}, {3, 4}}
		expected := []int{1, -1}

		assert.Equal(t, expected, minimumCost(n, edges, query))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}
		query := [][]int{{1, 2}}
		expected := []int{0}

		assert.Equal(t, expected, minimumCost(n, edges, query))
	})
}
