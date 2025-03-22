package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCompleteComponents(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}
		expected := 3

		assert.Equal(t, expected, countCompleteComponents(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}
		expected := 1

		assert.Equal(t, expected, countCompleteComponents(n, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 3
		edges := [][]int{{1, 0}, {2, 0}, {2, 1}}
		expected := 1

		assert.Equal(t, expected, countCompleteComponents(n, edges))
	})
}
