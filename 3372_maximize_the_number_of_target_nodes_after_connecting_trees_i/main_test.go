package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTargetNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges1 := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
		edges2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}
		k := 2
		expected := []int{9, 7, 9, 8, 8}

		assert.Equal(t, expected, maxTargetNodes(edges1, edges2, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}
		edges2 := [][]int{{0, 1}, {1, 2}, {2, 3}}
		k := 1
		expected := []int{6, 3, 3, 3, 3}

		assert.Equal(t, expected, maxTargetNodes(edges1, edges2, k))
	})
}
