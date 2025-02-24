package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostProfitablePath(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
		bob := 3
		amount := []int{-2, 4, 2, -4, 6}
		expected := 6

		assert.Equal(t, expected, mostProfitablePath(edges, bob, amount))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges := [][]int{{0, 1}}
		bob := 1
		amount := []int{-7280, 2350}
		expected := -7280

		assert.Equal(t, expected, mostProfitablePath(edges, bob, amount))
	})

	t.Run("test case 3", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
		bob := 3
		amount := []int{-5644, -6018, 1188, -8502}
		expected := -11662

		assert.Equal(t, expected, mostProfitablePath(edges, bob, amount))
	})
}
