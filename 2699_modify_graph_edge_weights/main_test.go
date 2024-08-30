package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifiedGraphEdges(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		edges := [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}
		source := 0
		destination := 1
		target := 5
		expected := [][]int{{4, 1, 1}, {2, 0, 1}, {0, 3, 1}, {4, 3, 3}}

		assert.Equal(t, expected, modifiedGraphEdges(n, edges, source, destination, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1, -1}, {0, 2, 5}}
		source := 0
		destination := 2
		target := 6
		expected := [][]int{}

		assert.Equal(t, expected, modifiedGraphEdges(n, edges, source, destination, target))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 4
		edges := [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}
		source := 0
		destination := 2
		target := 6
		expected := [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, 1}}

		assert.Equal(t, expected, modifiedGraphEdges(n, edges, source, destination, target))
	})
}
