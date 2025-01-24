package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventualSafeNodes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
		expected := []int{2, 4, 5, 6}

		assert.Equal(t, expected, eventualSafeNodes(graph))
	})

	t.Run("test case 2", func(t *testing.T) {
		graph := [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}
		expected := []int{4}

		assert.Equal(t, expected, eventualSafeNodes(graph))
	})

	t.Run("test case 3", func(t *testing.T) {
		graph := [][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}
		expected := []int{0, 1, 2, 3, 4}

		assert.Equal(t, expected, eventualSafeNodes(graph))
	})
}
