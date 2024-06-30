package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
		expected := []int{2, 3}

		assert.Equal(t, expected, findRedundantConnection(edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
		expected := []int{1, 4}

		assert.Equal(t, expected, findRedundantConnection(edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		edges := [][]int{{7, 8}, {2, 6}, {2, 8}, {1, 4}, {9, 10}, {1, 7}, {3, 9}, {6, 9}, {3, 5}, {3, 10}}
		expected := []int{3, 10}

		assert.Equal(t, expected, findRedundantConnection(edges))
	})
}
