package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	t.Run("should return shortest path cost", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}}
		obj := Constructor(n, edges)
		assert.Equal(t, 6, obj.ShortestPath(3, 2))
		assert.Equal(t, -1, obj.ShortestPath(0, 3))
		obj.AddEdge([]int{1, 3, 4})
		assert.Equal(t, 6, obj.ShortestPath(0, 3))
	})
}
