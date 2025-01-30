package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMagnificentSets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 6
		edges := [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}
		expected := 4

		assert.Equal(t, expected, magnificentSets(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		edges := [][]int{{1, 2}, {2, 3}, {3, 1}}
		expected := -1

		assert.Equal(t, expected, magnificentSets(n, edges))
	})

}
