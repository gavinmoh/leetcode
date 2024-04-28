package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfDistancesInTree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
		expected := []int{8, 12, 6, 10, 10, 10}

		assert.Equal(t, expected, sumOfDistancesInTree(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		edges := [][]int{}
		expected := []int{0}

		assert.Equal(t, expected, sumOfDistancesInTree(n, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 2
		edges := [][]int{{1, 0}}
		expected := []int{1, 1}

		assert.Equal(t, expected, sumOfDistancesInTree(n, edges))
	})
}
