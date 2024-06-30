package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumEdgesToRemove(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 4
		edges := [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}
		expected := 2

		assert.Equal(t, expected, maxNumEdgesToRemove(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		edges := [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}
		expected := 0

		assert.Equal(t, expected, maxNumEdgesToRemove(n, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 4
		edges := [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}
		expected := -1

		assert.Equal(t, expected, maxNumEdgesToRemove(n, edges))
	})

}
