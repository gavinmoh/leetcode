package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondMinimum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		edges := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
		time := 3
		change := 5
		expected := 13

		assert.Equal(t, expected, secondMinimum(n, edges, time, change))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 2
		edges := [][]int{{1, 2}}
		time := 3
		change := 2
		expected := 11

		assert.Equal(t, expected, secondMinimum(n, edges, time, change))
	})
}
