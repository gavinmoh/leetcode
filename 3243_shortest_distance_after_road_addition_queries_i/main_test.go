package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestDistanceAfterQueries(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
		expected := []int{3, 2, 1}

		assert.Equal(t, expected, shortestDistanceAfterQueries(n, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		queries := [][]int{{0, 3}, {0, 2}}
		expected := []int{1, 1}

		assert.Equal(t, expected, shortestDistanceAfterQueries(n, queries))
	})
}
