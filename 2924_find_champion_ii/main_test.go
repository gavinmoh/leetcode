package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindChampion(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1}, {1, 2}}
		expected := 0

		assert.Equal(t, expected, findChampion(n, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 2}, {1, 3}, {1, 2}}
		expected := -1

		assert.Equal(t, expected, findChampion(n, edges))
	})
}
