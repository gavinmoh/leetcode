package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDiameterAfterMerge(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}}
		edges2 := [][]int{{0, 1}}
		expected := 3

		assert.Equal(t, expected, minimumDiameterAfterMerge(edges1, edges2))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges1 := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}
		edges2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}
		expected := 5

		assert.Equal(t, expected, minimumDiameterAfterMerge(edges1, edges2))
	})
}
