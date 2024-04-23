package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinHeightTrees(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
		expected := []int{1}

		assert.Equal(t, expected, findMinHeightTrees(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 6
		edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
		expected := []int{3, 4}

		assert.Equal(t, expected, findMinHeightTrees(n, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 1
		edges := [][]int{}
		expected := []int{0}

		assert.Equal(t, expected, findMinHeightTrees(n, edges))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 8
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 4}, {4, 5}, {4, 6}, {6, 7}}
		expected := []int{0}

		assert.Equal(t, expected, findMinHeightTrees(n, edges))
	})
}
