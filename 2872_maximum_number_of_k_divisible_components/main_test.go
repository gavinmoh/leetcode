package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxKDivisibleComponents(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		edges := [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}
		values := []int{1, 8, 1, 4, 4}
		k := 6
		expected := 2

		assert.Equal(t, expected, maxKDivisibleComponents(n, edges, values, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 7
		edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}
		values := []int{3, 0, 6, 1, 5, 2, 1}
		k := 3
		expected := 3

		assert.Equal(t, expected, maxKDivisibleComponents(n, edges, values, k))
	})
}
